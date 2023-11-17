package delivery

import (
	"fmt"
	"log"

	"github.com/frasnym/go-try-rbac/middleware"
	"github.com/labstack/echo/v4"
)

type Router interface {
	Start()
}

type router struct {
	app  *echo.Echo
	port string
}

// Start implements Router.
func (r *router) Start() {
	// Start the server
	err := r.app.Start(fmt.Sprint(":", r.port))
	if err != nil {
		log.Fatal(err)
	}
}

func NewRouter(port string) Router {
	e := echo.New()

	e.Use(middleware.CheckAuth)
	registerHandler(e)

	return &router{port: port, app: e}
}
