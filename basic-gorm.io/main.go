package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "Kittamet"})

	r := gin.Default()
	r.GET("/users/first", func(c *gin.Context) {
		var u User
		// find first user
		db.First(&u)
		c.JSON(200, u)
	})

	userHandler := UserHandler{db: db}
	r.GET("/users", userHandler.User)
	r.Run()
}

// gorm provide pool connection can define like this but it not best practice
// var db *gorm.DB

// best practice for create db instance outside main
type UserHandler struct {
	db *gorm.DB
}

func (h *UserHandler) User(c *gin.Context) {
	var u User
	h.db.Find(&u)
	c.JSON(200, u)
}
