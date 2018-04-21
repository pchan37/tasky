package views

import (
	"net/http"

	"github.com/PGonLib/PGo-Auth/pkg/security"
	"github.com/pchan37/tasky/app/lib/templateManager"
)

func RegisterSecurityViews() {
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/logout", LogoutHandler)
	http.HandleFunc("/permission_denied", NotAuthorizedHandler)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templateManager.RenderTemplate(w, "register.tmpl", nil)
	case http.MethodPost:
		userCredential := &security.Credential{
			Username:             r.FormValue("username"),
			Password:             r.FormValue("password"),
			ConfirmationPassword: r.FormValue("confirmationPassword"),
		}
		validateRegisterCredentials(w, r, userCredential)
		if success := security.Register(userCredential); success {
			fullCredential, _ := security.GetUserCredential(userCredential.Username)
			session, err := security.GetSession(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			session.Values["authenticated"] = true
			session.Values["username"] = fullCredential.Username
			session.Values["role"] = fullCredential.Role
			session.Save(r, w)
			redirectBack(w, r)
		}
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		templateManager.RenderTemplate(w, "login.tmpl", nil)
	case http.MethodPost:
		userCredential := &security.Credential{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		if success := security.Login(userCredential); success {
			fullCredential, _ := security.GetUserCredential(userCredential.Username)
			session, err := security.GetSession(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			session.Values["authenticated"] = true
			session.Values["username"] = fullCredential.Username
			session.Values["role"] = fullCredential.Role
			redirectBack(w, r)
		}
		data := map[string]string{"Messages": "Incorrect username or password!"}
		templateManager.RenderTemplate(w, "login.tmpl", data)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := security.GetSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func NotAuthorizedHandler(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderTemplate(w, "403.tmpl", nil)
}

func redirectBack(w http.ResponseWriter, r *http.Request) {
	session, _ := security.GetSession(w, r)
	if redirectURL, ok := session.Values["redirect-url"].(string); ok && redirectURL != "" {
		session.Values["redirect-url"] = ""
		session.Save(r, w)
		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
	}
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func validateRegisterCredentials(w http.ResponseWriter, r *http.Request, c *security.Credential) {
	if security.IsRegistered(c.Username) {
		data := map[string]string{"Messages": "Username already taken!"}
		templateManager.RenderTemplate(w, "register.tmpl", data)
	} else if c.Password != c.ConfirmationPassword {
		data := map[string]string{"Messages": "Password does not match confirmation password!"}
		templateManager.RenderTemplate(w, "register.tmpl", data)
	}
}
