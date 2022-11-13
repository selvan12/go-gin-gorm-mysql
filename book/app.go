package book

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/selvan12/go-gin-gorm-mysql/database"
	"github.com/selvan12/go-gin-gorm-mysql/model"
	"gorm.io/gorm"
)

// App struct contains the app members
type App struct {
	Engine *gin.Engine
	Db     *gorm.DB
}

// NewApp creates the App instance with necessary init values
func NewApp() *App {
	log.Println("NewApp In")
	r := gin.Default()
	db := database.InitDatabase()
	db.AutoMigrate(&model.Book{})
	return &App{
		Engine: r,
		Db:     db,
	}
}
