package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// if schema change it will automatically alter table
	db.AutoMigrate(&Book{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/books", NewBook)
	r.GET("/books", ListBook)
	r.GET("/books/:id", GetBook)

	r.Run()
}

type Book struct {
	// define db table structure
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
}

func NewBook(c *gin.Context) {
	var book Book
	// bind json to struct Book
	if err := c.Bind(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := db.Create(&book)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

func ListBook(c *gin.Context) {
	// instance for receive data from db
	var books []Book
	result := db.Find(&books)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// return for break function when error
		return
	}

	// gin handler will automatically convert data structure (books) to JSON message
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	// get param "id"
	id := c.Param("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var book Book
	result := db.First(&book, n)
	if err := result.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, book)
}
