package main

import (
	"github.com/frasnym/go-try-rbac/db"
	"github.com/frasnym/go-try-rbac/delivery"
)

func main() {
	gormDB, sqlDB := db.InitDB()
	defer sqlDB.Close()

	delivery.InitRBAC(gormDB)

	router := delivery.NewRouter("8080")
	router.Start()
}
