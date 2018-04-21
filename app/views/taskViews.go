package views

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/PGonLib/PGo-Auth/pkg/security"
	"github.com/pchan37/tasky/app/lib/taskDatabase"
)

func RegisterTaskViews() {
	http.HandleFunc("/load_tasks", GetTasks)
	http.HandleFunc("/new_task", NewTask)
	http.HandleFunc("/update_task", UpdateTask)
	http.HandleFunc("/update_task_pos", UpdateTaskPosition)
	http.HandleFunc("/delete_task", DeleteTask)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	username, _ := security.GetUsername(w, r)
	js := taskDatabase.GetAll(username)
	if js == nil {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			username, _ := security.GetUsername(w, r)
			task := taskDatabase.Task{}
			json.Unmarshal(body, &task)
			task.Username = username
			if !taskDatabase.Insert(task) {
				log.Println("Error occurred while inserting task!")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			username, _ := security.GetUsername(w, r)
			task := taskDatabase.Task{}
			json.Unmarshal(body, &task)
			task.Username = username
			if !taskDatabase.Update(task) {
				log.Println("Error occurred while updating task!")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func UpdateTaskPosition(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			username, _ := security.GetUsername(w, r)
			taskPosition := taskDatabase.TaskPosition{}
			json.Unmarshal(body, &taskPosition)
			taskPosition.Username = username
			if !taskDatabase.UpdatePosition(taskPosition) {
				log.Println("Error occurred while updating task position!")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			username, _ := security.GetUsername(w, r)
			taskPosition := taskDatabase.TaskPosition{}
			json.Unmarshal(body, &taskPosition)
			taskPosition.Username = username
			if !taskDatabase.Remove(taskPosition) {
				log.Println("Error occurred while removing task!")
				w.WriteHeader(http.StatusBadRequest)
			} else {
				w.WriteHeader(http.StatusOK)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}
