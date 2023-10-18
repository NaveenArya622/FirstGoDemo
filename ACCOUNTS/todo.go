package ACCOUNTS

import (
	"strconv"
	"time"
)

var responceList []Tasks

type Tasks struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	ArchivedAt  time.Time `json:"archived_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AssiniedTo  string    `json:"assinied_to"`
}

type Responce struct {
	Message string   `json:"message"`
	Id      string   `json:"id"`
	Details *Tasks   `json:"details"`
	Tasks   *[]Tasks `json:"tasks"`
}

var todoList = []Tasks{}

func (task Tasks) isAvlable() int {
	for i := range todoList {
		if todoList[i].Id == task.Id && todoList[i].Status == "Active" {
			return i
		}
	}
	return -1
}

func (task Tasks) CreateTask() (int, Responce) {
	task.Id = "T" + strconv.Itoa((100000 + len(todoList)))
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	task.Status = "Active"
	if task.isAvlable() == -1 {
		todoList = append(todoList[0:], task)
		return 200, Responce{Message: "Task Added Successfully", Details: &task}
	}
	return 400, Responce{Message: "task exist"}

}

func (task Tasks) UpdateTask() (int, Responce) {
	task.UpdatedAt = time.Now()
	id := task.isAvlable()
	if id != -1 {
		if task.AssiniedTo != todoList[id].AssiniedTo {
			return 401, Responce{Message: "Unauthorized to get info"}
		}
		if task.Name != "" {
			todoList[id].Name = task.Name
		}
		if task.Description != "" {
			todoList[id].Description = task.Description
		}
		return 200, Responce{Message: "Task Update Successfully", Details: &todoList[id]}
	}
	return 400, Responce{Message: "Task not Exist"}

}

func (task Tasks) RemoveTask() (int, Responce) {
	id := task.isAvlable()
	if id != -1 {
		if task.AssiniedTo != todoList[id].AssiniedTo {
			return 401, Responce{Message: "Unauthorized to get info"}
		}
		task = todoList[id]
		task.Status = "Inactive"
		task.ArchivedAt = time.Now()
		task.UpdatedAt = time.Now()
		todoList[id] = task
		return 200, Responce{Message: "Task Update Successfully", Details: &task}
	}
	return 400, Responce{Message: "Task not Exist"}
}

func (task Tasks) CompletedTask() (int, Responce) {
	id := task.isAvlable()
	if id != -1 {
		if task.AssiniedTo != todoList[id].AssiniedTo {
			return 401, Responce{Message: "Unauthorized to get info"}
		}
		task = todoList[id]
		task.Status = "Completed"
		task.UpdatedAt = time.Now()
		todoList[id] = task
		return 200, Responce{Message: "Task Update Successfully", Details: &task}
	}
	return 400, Responce{Message: "Task not Exist"}
}

func (task Tasks) GetTasks() (int, Responce) {
	responceList = todoList[0:0]
	for i := range todoList {
		if todoList[i].AssiniedTo == task.AssiniedTo && (task.Status == "" || todoList[i].Status == task.Status) {
			responceList = append(responceList[0:], todoList[i])
		}
	}
	if len(responceList) == 0 {
		return 200, Responce{Message: "No Todo List for ID: " + task.AssiniedTo, Tasks: &responceList}
	}
	return 200, Responce{Message: "Todo List for ID: " + task.AssiniedTo, Tasks: &responceList}
}
func (task Tasks) GetTask() (int, Responce) {
	id := task.isAvlable()
	if id != -1 {
		if task.AssiniedTo != todoList[id].AssiniedTo {
			return 401, Responce{Message: "Unauthorized to get info"}
		}
		return 200, Responce{Message: "Task Update Successfully", Details: &todoList[id]}
	}
	return 400, Responce{Message: "Todo List for ID: " + task.AssiniedTo}
}
