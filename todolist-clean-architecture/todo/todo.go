package todo

import (
	"log"
	"net/http"
	"time"
)

type Todo struct {
	Title     string `json:"text"`
	ID        uint   `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Todo) TableName() string {
	return "todos"
}

type storer interface {
	New(*Todo) error
	List(todo *[]Todo) error
	Remove(todo *Todo, id int) error
}

type TodoHandler struct {
	store storer
}

func NewTodoHandler(store storer) *TodoHandler {
	return &TodoHandler{store: store}
}

type Context interface {
	Bind(interface{}) error
	JSON(int, interface{})
	TransactionID() string
	Audience() string
	IDParam() (int, error)
}

func (t *TodoHandler) NewTask(c Context) {
	var todo Todo
	if err := c.Bind(&todo); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	if todo.Title == "sleep" {
		transactionID := c.TransactionID()
		aud := c.Audience()
		log.Println(transactionID, aud, "not allowed")
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "not allowed",
		})
		return
	}

	err := t.store.New(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, map[string]interface{}{
		"ID": todo.ID,
	})
}

func (t *TodoHandler) List(c Context) {
	var todos []Todo
	err := t.store.List(&todos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) Remove(c Context) {
	id, err := c.IDParam()
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	err = t.store.Remove(&Todo{}, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
	})
}
