package book

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/selvan12/go-gin-gorm-mysql/model"
)

/*
// Book struct to represent the book member fields
type Req struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	Price         float32 `json:"price"`
	Pages         int     `json:"pages"`
	DatePublished string  `json:"date_published"`
}
*/

// listBooks handler to list books
func (a *App) listBooks(c *gin.Context) {
	var books []model.Book
	err := model.GetBooks(a.Db, &books)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	fmt.Println("List API Success\n", books)
	c.JSON(http.StatusOK, books)
}

// getBook handler to get book by id
func (a *App) getBook(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	fmt.Println("GET api called and ID = ", id)

	err := model.GetBook(a.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	fmt.Println("GET by ID API Success\n", book)
	c.JSON(http.StatusOK, book)
}

// createBook handler to create book entry
func (a *App) createBook(c *gin.Context) {
	var book model.Book
	//var req Book

	c.BindJSON(&book)
	id := uuid.New()
	book.ID = id.String()
	fmt.Println("create book json req : ", book)

	// book.ID = req.ID
	// book.Name = req.Name
	// book.Author = req.Author
	// book.Price = req.Price
	// book.Pages = req.Pages
	// book.DatePublished = req.DatePublished

	err := model.CreateBook(a.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	fmt.Println("Create API Success\n", book)
	c.JSON(http.StatusCreated, book)
}

// patchBook handler to update book by id
func (a *App) patchBook(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	fmt.Println("PATCH api called and ID = ", id)

	err := model.GetBook(a.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.BindJSON(&book)
	fmt.Println("Patch book json req : ", book)
	err = model.UpdateBook(a.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	fmt.Println("Patch API Success\n", book)
	c.JSON(http.StatusOK, book)
}

// deleteBook handler to delete the book entry by id
func (a *App) deleteBook(c *gin.Context) {
	var book model.Book
	id := c.Param("id")
	fmt.Println("Delete api called and ID = ", id)

	err := model.DeleteBook(a.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	fmt.Println("Delete API Success")
	c.JSON(http.StatusNoContent, nil)
}
