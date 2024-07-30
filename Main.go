package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.GET("hasan", Hello)

	dsn := "sqlserver://localhost?database=HsnBlog&trusted_connection=yes"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&employee{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router.Run("localhost:8080")
}

func Hello(c *gin.Context) {
	deneme := "hasan"
	c.IndentedJSON(http.StatusOK, deneme)
}

type employee struct {
	gorm.Model
	Name string
}
