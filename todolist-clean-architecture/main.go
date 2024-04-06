package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/montheankul-k/todolist-clean-architecture/auth"
	"github.com/montheankul-k/todolist-clean-architecture/router"
	"github.com/montheankul-k/todolist-clean-architecture/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/time/rate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/montheankul-k/todolist-clean-architecture/todo"
)

var (
	buildcommit = "dev"
	buildtime   = time.Now().String()
)

func main() {
	_, err := os.Create("/tmp/live")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove("/tmp/live")

	err = godotenv.Load("local.env")
	if err != nil {
		log.Printf("please consider environment variable: %s", err)
	}

	db, err := gorm.Open(mysql.Open(os.Getenv("DB_CONN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic("failed to connect database")
	}
	collection := client.Database("myapp").Collection("todos")

	r := router.NewMyRouter()
	// r := router.NewFiberRouter()

	r.GET("/health", func(c *gin.Context) {
		c.Status(200)
	})

	r.GET("/limit", limitedHandler)

	r.GET("/x", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"buildcommit": buildcommit,
			"buildtime":   buildtime,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/token", auth.AccessToken(os.Getenv("SIGN")))
	protected := r.Group("", auth.Protect([]byte(os.Getenv("SIGN"))))

	gormStore := store.NewGormStore(db)
	_ = store.NewMongoDBStore(collection)

	handler := todo.NewTodoHandler(gormStore)
	r.POST("/todos", handler.NewTask)
	protected.GET("/todos", router.NewGinHandler(handler.List))
	protected.DELETE("/todos/:id", router.NewGinHandler(handler.Remove))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	s := &http.Server{
		Addr:              ":" + os.Getenv("PORT"),
		Handler:           r, // dont need r.Run()
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// if use fiber use this go routine and don't need &http.Server{}
	/*
		go func() {
			if err := r.Listen(":" + os.Getenv("PORT")); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()
	*/

	<-ctx.Done()
	stop()
	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(timeoutCtx); err != nil {
		fmt.Println(err)
	}
}

var limiter = rate.NewLimiter(5, 5)

func limitedHandler(c *gin.Context) {
	if !limiter.Allow() {
		c.AbortWithStatus(http.StatusTooManyRequests)
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
