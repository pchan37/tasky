package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"tasky/src/lib/tasks"
)

func RegisterTaskViews() {
	http.HandleFunc("/new_task", NewTask)
	http.HandleFunc("/update_task", UpdateTask)
	http.HandleFunc("/update_task_pos", UpdateTaskPosition)
	http.HandleFunc("/delete_task", DeleteTask)
}

func NewTask(w http.ResponseWriter, r *http.Request) {

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			response := tasks.Task{}
			json.Unmarshal(body, &response)
			fmt.Println(response)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}

func UpdateTaskPosition(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			response := tasks.TaskPosition{}
			json.Unmarshal(body, &response)
			fmt.Println(response)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			response := tasks.TaskPosition{}
			json.Unmarshal(body, &response)
			fmt.Println(response)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}
