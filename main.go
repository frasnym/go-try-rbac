package main

import (
	"time"

	"github.com/frasnym/go-try-rbac/db"
	"github.com/frasnym/go-try-rbac/delivery"
	"github.com/frasnym/go-try-rbac/pgk/rbac"
)

func main() {
	gormDB, sqlDB := db.InitDB()
	defer sqlDB.Close()

	rbac.RefreshRBAC(gormDB)

	go func() {
		for {
			time.Sleep(5 * time.Second)
			rbac.RefreshRBAC(gormDB)
		}
	}()

	router := delivery.NewRouter("8080")
	router.Start()
}
