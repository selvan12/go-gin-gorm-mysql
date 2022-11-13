package main

import (
	"log"

	"github.com/selvan12/go-gin-gorm-mysql/book"
)

func main() {
	log.Println("Starting the go-gin-gorm-mysql")
	app := book.NewApp()
	app.AddRoutes(app.Engine)
	_ = app.Engine.Run(":8080")
}
