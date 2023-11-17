package delivery

import (
	"net/http"

	"github.com/frasnym/go-try-rbac/middleware"
	"github.com/labstack/echo/v4"
)

func registerHandler(e *echo.Echo) {
	// Define routes with RBAC enforcement middleware
	e.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin route")
	}, middleware.EnforceRBAC(rbac, adminPerm))

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "User route")
	}, middleware.EnforceRBAC(rbac, userPerm))

	e.POST("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin can access this route")
	}, middleware.EnforceRBAC(rbac, adminPerm))

	e.GET("/admin-or-user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin or User route")
	}, middleware.EnforceRBAC(rbac, adminPerm, userPerm))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! (Accessible by all users)")
	})
}
