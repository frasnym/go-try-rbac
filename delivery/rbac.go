package delivery

import (
	"github.com/frasnym/go-try-rbac/model"

	"github.com/mikespook/gorbac"
	"gorm.io/gorm"
)

var rbac *gorbac.RBAC

func InitRBAC(db *gorm.DB) {
	rbac = gorbac.New()

	roles, err := GetUserRolesAndPermissions(db)
	if err != nil {
		panic("Error retrieving roles and permissions")
	}

	// Add roles and permissions to the RBAC instance
	for _, role := range *roles {
		r := gorbac.NewStdRole(role.Name)
		for _, permission := range role.Permissions {
			p := gorbac.NewStdPermission(permission.Name)
			r.Assign(p)
		}
		rbac.Add(r)
	}
}

// GetUserRolesAndPermissions retrieves roles and permissions for a user from the database
func GetUserRolesAndPermissions(db *gorm.DB) (*[]model.Role, error) {
	var roles []model.Role
	if err := db.Preload("Permissions").Find(&roles).Error; err != nil {
		return nil, err
	}
	return &roles, nil
}
