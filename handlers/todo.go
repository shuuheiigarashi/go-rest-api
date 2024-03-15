// handlers/todo.go
package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shuuheiigarashi/go-rest-api/models"
)

var todoList []models.ToDo
var currentID int

// GetToDoList はToDoリストを取得するハンドラーです
func GetToDoList(c *gin.Context) {
	c.JSON(http.StatusOK, todoList)
}

// CreateToDoItem は新しいToDoアイテムを作成するハンドラーです
func CreateToDoItem(c *gin.Context) {
	var newItem models.ToDo
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newItem.ID = currentID
	currentID++
	todoList = append(todoList, newItem)
	c.JSON(http.StatusCreated, newItem)
}

// UpdateToDoItem は既存のToDoアイテムを更新するハンドラーです
func UpdateToDoItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedItem models.ToDo
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	found := false
	for i, item := range todoList {
		if item.ID == id {
			todoList[i] = updatedItem
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, updatedItem)
}

// DeleteToDoItem はToDoアイテムを削除するハンドラーです
func DeleteToDoItem(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	found := false
	for i, item := range todoList {
		if item.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}
