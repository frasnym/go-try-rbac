package delivery

import (
	"net/http"

	"github.com/frasnym/go-try-rbac/middleware"
	"github.com/labstack/echo/v4"
)

func registerHandler(e *echo.Echo) {
	// Define routes with RBAC enforcement middleware
	e.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET admin")
	}, middleware.EnforceRBAC(rbac))

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "GET user")
	}, middleware.EnforceRBAC(rbac))

	e.POST("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin can access this route")
	}, middleware.EnforceRBAC(rbac))

	e.GET("/admin-or-user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin or User route")
	}, middleware.EnforceRBAC(rbac))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! (Accessible by all users)")
	})
}
