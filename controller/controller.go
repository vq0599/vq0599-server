package controller

import (
		"net/http"
		"github.com/gin-gonic/gin"
		"database/sql"
		// "fmt"
    _"github.com/go-sql-driver/mysql"
		"log"
		"time"

)

type Article struct {
	Id						int `json:"id"`
	Title					string `json:"title"`
	Content				string `json:"content"`
	Create_time		int64 `json:"create_time"`
}


var (
	id int
	title string
	content string
	create_time time.Time
)

func GetArticle (c *gin.Context) {
	var results []Article

	db, _ := sql.Open("mysql", "root:ceoms999@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=true")
	defer db.Close()
	
	db.Query("use test")

	rows, _ := db.Query("select * from test.articles")
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &title, &content, &create_time)
    if err != nil {
        log.Fatal(err)
		}
		
		unit64_create_time := create_time.UnixNano()/1e6
		results = append(results, Article{id, title, content, unit64_create_time})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "SUCCESS",
		"code": 200,
		"data": results,
	})
}

func AddArticle (c *gin.Context) {
	// db, err := sql.Open("mysql", "root:ceoms999@tcp(127.0.0.1:3306)/?charset=utf8")
	
	// db.Query("use test")

	// row, err := db.Query("INSERT INTO articles (title, content) VALUES ('文章标题3', 'Champs-Elysees')")

	// log.Printf("insert result %v\n", row)

	// c.JSON(http.StatusOK, gin.H{
	// 	"code": 200,
	// 	"data": nil,
	// 	"msg": "添加成功",
	// })

	// defer db.Close()
}
