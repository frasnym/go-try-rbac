package delivery

import "github.com/mikespook/gorbac"

var rbac *gorbac.RBAC
var adminPerm gorbac.Permission
var userPerm gorbac.Permission

func init() {
	rbac = gorbac.New()

	// Create roles and permissions
	adminRole := gorbac.NewStdRole("admin")
	userRole := gorbac.NewStdRole("user")

	adminPerm = gorbac.NewStdPermission("adminPermission")
	userPerm = gorbac.NewStdPermission("userPermission")

	// Add permissions to roles
	adminRole.Assign(adminPerm)
	userRole.Assign(userPerm)

	// Add roles to RBAC
	rbac.Add(adminRole)
	rbac.Add(userRole)
}
