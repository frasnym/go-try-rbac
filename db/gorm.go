package db

import (
	"database/sql"

	"github.com/frasnym/go-try-rbac/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, *sql.DB) {
	// Initialize the database
	db, err := gorm.Open(sqlite.Open("sqlite.db"))
	if err != nil {
		panic("Failed to connect to database")
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{})

	return db, sqlDB
}
