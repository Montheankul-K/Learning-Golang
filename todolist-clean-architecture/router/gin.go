package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/montheankul-k/todolist-clean-architecture/todo"
	"strconv"
)

type MyContext struct {
	*gin.Context
}

func NewGinContext(c *gin.Context) *MyContext {
	return &MyContext{Context: c} // composition
}

func (c *MyContext) Bind(v interface{}) error {
	return c.Context.ShouldBindJSON(v)
}

func (c *MyContext) JSON(statusCode int, v interface{}) {
	c.Context.JSON(statusCode, v)
}

func (c *MyContext) TransactionID() string {
	return c.Request.Header.Get("TransactionID")
}

func (c *MyContext) Audience() string {
	if aud, ok := c.Get("aud"); ok {
		if s, ok := aud.(string); ok {
			return s
		}
	}
	return ""
}

func (c *MyContext) IDParam() (int, error) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewGinHandler(handler func(todo.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(NewGinContext(c))
	}
}

type MyRouter struct {
	*gin.Engine
}

func NewMyRouter() *MyRouter {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:8080",
	}

	config.AllowHeaders = []string{
		"Origin",
		"Authorization",
		"TransactionID",
	}
	r.Use(cors.New(config))

	return &MyRouter{r}
}

func (r *MyRouter) POST(path string, handler func(todo.Context)) {
	r.Engine.POST(path, NewGinHandler(handler))
}
