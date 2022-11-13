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

func GetBooks(db *gorm.DB, book *[]Book) (err error) {
	err = db.Find(book).Error
	if err != nil {
		return err
	}
	return nil
}

func GetBook(db *gorm.DB, book *Book, id string) (err error) {
	err = db.Where("id = ?", id).Find(book).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateBook(db *gorm.DB, book *Book) (err error) {
	err = db.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateBook(db *gorm.DB, book *Book, id string) (err error) {
	db.Save(book)
	return nil
}

func DeleteBook(db *gorm.DB, book *Book, id string) (err error) {
	err = db.Where("id = ?", id).Delete(book).Error
	if err != nil {
		return err
	}
	return nil
}
