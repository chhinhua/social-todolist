package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	//init the loc
	loc, _ := time.LoadLocation("Asia/Bangkok")

	//set timezone,
	now := time.Now().In(loc)
	item := TodoItem{
		Id:          1,
		Title:       "This is task 1",
		Description: "Task 1 description",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	_ = r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
