// models/todo.go
package models

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
