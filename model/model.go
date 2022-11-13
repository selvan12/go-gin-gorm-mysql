package model

import (
	"gorm.io/gorm"
)

// Book struct to represent the book member fields
type Book struct {
	gorm.Model
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Author        string  `json:"author"`
	Price         float32 `json:"price"`
	Pages         int     `json:"pages"`
	DatePublished string  `json:"date_published"`
}

// GetBooks - fetch the list of books
func GetBooks(db *gorm.DB, book *[]Book) (err error) {
	err = db.Find(book).Error
	if err != nil {
		return err
	}
	return nil
}

// GetBook by ID
func GetBook(db *gorm.DB, book *Book, id string) (err error) {
	err = db.Where("id = ?", id).Find(book).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateBook create a book entry
func CreateBook(db *gorm.DB, book *Book) (err error) {
	err = db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateBook update book by id
func UpdateBook(db *gorm.DB, book *Book, id string) (err error) {
	db.Save(book)
	return nil
}

// DeleteBook delete the book entry by id
func DeleteBook(db *gorm.DB, book *Book, id string) (err error) {
	err = db.Where("id = ?", id).Delete(book).Error
	if err != nil {
		return err
	}
	return nil
}
