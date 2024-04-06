package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/montheankul-k/todolist-clean-architecture/todo"
	"strconv"
)

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()

	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

func (r *FiberRouter) POST(path string, handler func(ctx todo.Context)) {
	r.App.Post(path, func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	})
}

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) JSON(statusCode int, v interface{}) {
	c.Ctx.Status(statusCode).JSON(v)
}

func (c *FiberContext) TransactionID() string {
	return string(c.Ctx.Request().Header.Peek("TransactionID"))
}

func (c *FiberContext) Audience() string {
	return c.Ctx.Get("aud")
}

func (c *FiberContext) IDParam() (int, error) {
	idParam := c.Ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0, err
	}
	return id, nil
}
