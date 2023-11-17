package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mikespook/gorbac"
)

func main() {
	e := echo.New()

	// Initialize gorbac
	rbac := gorbac.New()

	// Create roles and permissions
	adminRole := gorbac.NewStdRole("admin")
	userRole := gorbac.NewStdRole("user")

	adminPerm := gorbac.NewStdPermission("adminPermission")
	userPerm := gorbac.NewStdPermission("userPermission")

	// Add permissions to roles
	adminRole.Assign(adminPerm)
	userRole.Assign(userPerm)

	// Add roles to RBAC
	rbac.Add(adminRole)
	rbac.Add(userRole)

	e.Use(checkAuth)

	// Define routes with RBAC enforcement middleware
	e.GET("/admin", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin route")
	}, enforceRBAC(rbac, adminPerm))

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "User route")
	}, enforceRBAC(rbac, userPerm))

	e.POST("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin can access this route")
	}, enforceRBAC(rbac, adminPerm))

	e.GET("/admin-or-user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Admin or User route")
	}, enforceRBAC(rbac, adminPerm, userPerm))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! (Accessible by all users)")
	})

	// Start the server
	err := e.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

// Echo middleware to simulate authentication and retrieve user role
func checkAuth(next echo.HandlerFunc) echo.HandlerFunc {
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

// Middleware to enforce RBAC for a specific route
func enforceRBAC(rbac *gorbac.RBAC, perms ...gorbac.Permission) echo.MiddlewareFunc {
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
