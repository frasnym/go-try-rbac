package model

import "github.com/jinzhu/gorm"

// TODO: Unused for now, use it later
type User struct {
	gorm.Model
	Username string
	Password string
	Roles    []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
	ID          int `gorm:"primary_key"`
	Name        string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	ID   int `gorm:"primary_key"`
	Name string
}
