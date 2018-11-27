package models

import (
  "time"
)

type Article struct {
  Id            int `json:"id"`
  Title          string `json:"title"`
  Content        string `json:"content"`
  Create_time    int64 `json:"create_time"`
  Pv            int `json:"pv"`
}

// 获取文章列表
func GetArticles() []Article {
  var results []Article

  db, _ := Open()
  defer db.Close()
  
  rows, _ := db.Query("select * from articles")
  defer rows.Close()

  for rows.Next() {
    var article Article
    var create_time time.Time
    rows.Scan(&article.Id, &article.Title, &article.Content, &create_time, &article.Pv)
    
    article.Create_time = create_time.UnixNano()/1e6
    results = append(results, article)
  }

  return results
}

// 检查文章id存不存在
func CheckArticleExist(id int) bool {
  var count int
  db, _ := Open()
  defer db.Close()

  db.QueryRow("select count(*) from articles where id = ?", id).Scan(&count)

  if count == 0 {
    return false
  } else {
    return true
  }
}


// 添加文章的pv
func Count(id int) error {
  db, _ := Open()
  defer db.Close()

  _, err := db.Exec("update articles set pv = pv + 1 where id = ?", id)
  return err
}