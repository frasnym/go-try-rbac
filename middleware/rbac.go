package middleware

import (
	"fmt"
	"net/http"

	"github.com/frasnym/go-try-rbac/pgk/rbac"

	"github.com/labstack/echo/v4"
	"github.com/mikespook/gorbac"
)

// Middleware to enforce RBAC for a specific route
func EnforceRBAC() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rbac := rbac.GetRBAC()

			// Retrieve the user role from the context
			userRole, ok := c.Get("userRole").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized userRole"})
			}

			userPermission, ok := c.Get("userPermission").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized userPermission"})
			}

			// Check if the user role has any of the required permissions to access the route
			if rbac.IsGranted(userRole, gorbac.NewStdPermission(userPermission), nil) {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{"error": fmt.Sprintf("Unauthorized EnforceRBAC (%s %s)", userRole, userPermission)})
		}
	}
}
