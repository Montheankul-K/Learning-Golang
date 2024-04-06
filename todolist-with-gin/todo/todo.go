package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	Title string `json:"text"`
	gorm.Model
}

// manual naming table
func (Todo) TableName() string {
	return "todos"
}

type TodoHandler struct {
	db *gorm.DB
}

// best practice
func NewTodoHandler(db *gorm.DB) *TodoHandler {
	return &TodoHandler{db: db}
}

func (t *TodoHandler) NewTask(c *gin.Context) {
	/*
		// get token from request header
		s := c.Request.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(s, "Bearer ")

		if err := auth.Protect(tokenString); err != nil {
			// abort for stop chain of this handler (middlewares)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	*/
	var todo Todo
	// bind json
	// c.BindJSON: must bind if err return 400
	// c.ShouldBindJSON: can handle if err (more flexible)
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if todo.Title == "sleep" {
		transactionID := c.Request.Header.Get("TransactionID")
		aud, _ := c.Get("aud")
		log.Println(transactionID, aud, "not allowed")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "not allowed",
		})
		return
	}

	r := t.db.Create(&todo)
	if err := r.Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"ID": todo.Model.ID,
	})
}

func (t *TodoHandler) List(c *gin.Context) {
	var todos []Todo
	r := t.db.Find(&todos)
	if err := r.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (t *TodoHandler) Remove(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	r := t.db.Delete(&Todo{}, id)
	if err := r.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
