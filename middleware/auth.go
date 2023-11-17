package middleware

import "github.com/labstack/echo/v4"

// Echo middleware to simulate authentication and retrieve user role
func CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Simulate authentication logic by extracting the user role from a custom header
		userRole := c.Request().Header.Get("X-User-Role")

		// Check if the user role is empty
		if userRole == "" {
			userRole = "guest"
		}

		// Set the user role in the context for later use
		c.Set("userRole", userRole)

		return next(c)
	}
}
