package book

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// add all available routes
func (a *App) AddRoutes(r *gin.Engine) {
	log.Println("AddRoutes In")

	// use the ping command to test the API
	// "curl -X GET http://localhost:8080/ping"
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello.. Welcome !!!")
	})

	// REST CRUD operation handlers for books GORM MySQL database
	r.GET("/books", a.listBooks)
	r.GET("/books/:id", a.getBook)
	r.POST("/books", a.createBook)
	r.DELETE("/books/:id", a.deleteBook)
	r.PATCH("/books/:id", a.patchBook)

	log.Println("AddRoutes Out")
}
