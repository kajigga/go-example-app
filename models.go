package main

import "strconv"

type VersionJSON struct {
	Version string `json:"version"`
}
type HomeJSON struct {
	Elements  []int  `json:"elements"`
	PageTitle string `json:"page_title"`
}

type ToDoItem struct {
	//ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID     float64 `json:"_id,omitempty" bson:"_id,omitempty"`
	Task   string  `json:"task,omitempty"`
	Status bool    `json:"status,omitempty"`
}

type ToDoList struct {
	ToDos []ToDoItem `json:"todos"`
	Owner string     `json:"owner"`
}

var AllTodos = map[string]ToDoList{}

func (list *ToDoList) addTodo(item ToDoItem) {
	list.ToDos = append(list.ToDos, item)

}

func init() {
	list := ToDoList{Owner: "Kevin"}
	// AllTodos["kevin"] = ToDoList{Owner: "Kevin"}
	for x := 0; x < 20; x++ {
		list.addTodo(ToDoItem{Task: strconv.Itoa(x)})
	}
	AllTodos["kevin"] = list
}
