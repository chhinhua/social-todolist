package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
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

func (TodoItem) TableName() string {
	return "todo_items"
}

type TodoItemCreation struct {
	Id          int    `json:"-" gorm:"column:id"`
	Title       string `json:"title" gorm:"column:title"`
	Description string `json:"description" gorm:"column:description"`
	//Status      string `json:"status" gorm:"column:status"`
}

func (TodoItemCreation) TableName() string {
	return TodoItem{}.TableName()
}

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)

	r := gin.Default()
	//	CRUD: Create, Read, Update, Delete
	//	POST /v1/items (create a new item)
	//	GET /v1/items (list items) /v1/items?page=1
	//	GET /v1/items/:id (get item detail by id)
	//	(PUT || PATCH) /v1/items/:id (update an item by id)
	//	DELETE /v1/items/:id (delete an item by id)

	v1 := r.Group("/v1")
	{
		items := v1.Group("/items")
		{
			items.POST("", CreateItem(db))
			items.GET("")
			items.GET("/:id", GetItem(db))
			items.PATCH("/:id")
			items.DELETE("/:id")
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	_ = r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// CreateItem returns a gin.HandlerFunc that creates a new todo item in the database.
//
// It takes a *gorm.DB as a parameter to perform database operations.
//
// The returned function handles the HTTP request to create a new todo item:
//   - It binds the request body to a TodoItemCreation struct.
//   - If binding fails, it returns a 400 Bad Request error.
//   - It then attempts to create the item in the database.
//   - If creation fails, it returns a 400 Bad Request error.
//   - On success, it returns a 201 Created status with the new item's ID.
//
// Parameters:
//   - db: A pointer to a gorm.DB instance for database operations.
//
// Returns:
//   - A gin.HandlerFunc that can be used as a route handler in a Gin application.
func CreateItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreation

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Create(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": data.Id,
		})
	}
}

func GetItem(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": data,
		})
	}
}
