package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikespook/gorbac"
)

// Middleware to enforce RBAC for a specific route
func EnforceRBAC(rbac *gorbac.RBAC, perms ...gorbac.Permission) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Retrieve the user role from the context
			userRole, ok := c.Get("userRole").(string)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
			}

			// Check if the user role has any of the required permissions to access the route
			for _, perm := range perms {
				if rbac.IsGranted(userRole, perm, nil) {
					return next(c)
				}
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
		}
	}
}
