package main

import (
	"github.com/frasnym/go-try-rbac/delivery"
)

func main() {
	router := delivery.NewRouter("8080")
	router.Start()
}
