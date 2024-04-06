package todo_test

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/montheankul-k/todolist-with-gin/todo"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateTodoNotAllowSleepTask(t *testing.T) {
	// arrange
	// can use empty &gorm.DB because it not a dependency in this test case
	handler := todo.NewTodoHandler(&gorm.DB{})

	w := httptest.NewRecorder()
	payload := bytes.NewBufferString(`{"text":"sleep"}`)
	req, _ := http.NewRequest("POST", "http://0.0.0.0:8080/todos", payload)
	req.Header.Add("TransactionID", "txnID001")

	c, _ := gin.CreateTestContext(w) // return context, gin engine (not use)
	// add request to context
	c.Request = req

	// act
	handler.NewTask(c) // return json string

	// assert
	want := `{"error":"not allowed"}`

	if want != w.Body.String() {
		t.Errorf("want %s but get %s\n", want, w.Body.String())
	}
}
