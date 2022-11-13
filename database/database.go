package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "bookshop"
)

var Db *gorm.DB

func InitDatabase() *gorm.DB {
	fmt.Println("InitDatabase In")
	return databaseConnection()
}

func databaseConnection() *gorm.DB {
	var err error
	// DataSourceName is "root:root@tcp(127.0.0.1:3306)/bookshop"
	dsn := username + ":" + password + "@tcp" + "(" + hostname + ")/" + dbname + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error at database connection : ", err)
		return nil
	}

	fmt.Println("databaseConnection Out")
	return db
}
