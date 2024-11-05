package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	fmt.Println("hello")

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

	//jsonData, err := json.Marshal(item)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	/*fmt.Println(string(jsonData))

	jsonStr := "{\"id\":1,\"title\":\"This is task 1\",\"description\":\"Task 1 description\",\"status\":\"Doing\",\"created_at\":\"2024-11-03T17:01:05.7242401+07:00\",\"updated_at\":null}"

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(item2)*/

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": item,
		})
	})
	err := r.Run(":8081")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
