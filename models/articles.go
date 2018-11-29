package models

import (
  "time"
  "html"
)

type Article struct {
  Id            int `json:"id"`
  Title          string `json:"title"`
  Content        string `json:"content"`
  Create_time    int64 `json:"create_time"`
  Pv            int `json:"pv"`
}

// 获取文章列表
func GetArticles() ([]Article, error) {
  var results []Article

  db, _ := Open()
  defer db.Close()
  
  rows, err := db.Query("select * from articles")
  defer rows.Close()

  for rows.Next() {
    var article Article
    var create_time time.Time
    rows.Scan(&article.Id, &article.Title, &article.Content, &create_time, &article.Pv)
    
    article.Content = html.UnescapeString(article.Content)
    article.Create_time = create_time.UnixNano() / 1e6

    results = append(results, article)
  }

  return results, err
}

// 获取单篇文章
func GetArticle(id int) (Article, error) {
  var article Article
  var create_time time.Time

  db, _ := Open()
  defer db.Close()

  err := db.QueryRow("select * from articles where id = ?", id).Scan(
    &article.Id,
    &article.Title,
    &article.Content,
    &create_time,
    &article.Pv,
  )

  article.Create_time = create_time.UnixNano() / 1e6
  article.Content = html.UnescapeString(article.Content)

  return article, err
}

// 检查文章id存不存在
func CheckArticleExist(id int) bool {
  var flag int
  db, _ := Open()
  defer db.Close()

  err := db.QueryRow("select id from articles where id = ?", id).Scan(&flag)

  if err != nil {
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

// 添加文章
func AddArticle(title string, content string) (int64, error) {
  db, _ := Open()
  defer db.Close()

  stmt, preErr := db.Prepare("insert articles set title=?, content=?")

  if preErr != nil {
    return 0, preErr
  }


  res, _ := stmt.Exec(title, html.EscapeString(content))
  id, insetErr := res.LastInsertId()

  if insetErr != nil {
    return 0, nil
  } else {
    return id, insetErr
  }
}