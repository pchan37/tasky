package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tasky/src/lib/taskDatabase"
)

func RegisterTaskViews() {
	http.HandleFunc("/new_task", NewTask)
	http.HandleFunc("/update_task", UpdateTask)
	http.HandleFunc("/update_task_pos", UpdateTaskPosition)
	http.HandleFunc("/delete_task", DeleteTask)
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			task := taskDatabase.Task{}
			json.Unmarshal(body, &task)
			if !taskDatabase.Insert(task) {
				log.Println("Error occurred while inserting task!")
			}
			fmt.Println(task)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			task := taskDatabase.Task{}
			json.Unmarshal(body, &task)
			if !taskDatabase.Update(task) {
				log.Println("Error occurred while updating task!")
			}
			fmt.Println(task)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}

func UpdateTaskPosition(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			taskPosition := taskDatabase.TaskPosition{}
			json.Unmarshal(body, &taskPosition)
			if !taskDatabase.UpdatePosition(taskPosition) {
				log.Println("Error occurred while updating task position!")
			}
			fmt.Println(taskPosition)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err == nil {
			taskPosition := taskDatabase.TaskPosition{}
			json.Unmarshal(body, &taskPosition)
			if !taskDatabase.Remove(taskPosition) {
				log.Println("Error occurred while removing task!")
			}
			fmt.Println(taskPosition)
		} else {
			fmt.Println("Error occurred: ", err)
		}
	}
}
