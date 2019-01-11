package models

import (
  "time"
  "html"
  "vq0599/common"
  // "fmt"
)

type Article struct {
  Id            int      `json:"id"`
  Title         string   `json:"title"`
  Content       string   `json:"content"`
  Create_time   int64    `json:"create_time"`
  Pv            int      `json:"pv"`
  Tags          []string `json:"tags"`
}

// 获取文章列表
func GetArticles() ([]Article, error) {
  var results []Article
  var tags string

  db, _ := Open()
  defer db.Close()
  
  rows, err := db.Query("SELECT * FROM articles")
  defer rows.Close()

  for rows.Next() {
    var article Article
    var create_time time.Time
    rows.Scan(&article.Id, &article.Title, &article.Content, &create_time, &article.Pv, &tags)
    
    article.Tags = common.Split(tags, ",")
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
  var tags string

  db, _ := Open()
  defer db.Close()

  err := db.QueryRow("SELECT * FROM articles WHERE id = ?", id).Scan(
    &article.Id,
    &article.Title,
    &article.Content,
    &create_time,
    &article.Pv,
    &tags,
  )

  article.Tags = common.Split(tags, ",")
  article.Create_time = create_time.UnixNano() / 1e6
  article.Content = html.UnescapeString(article.Content)

  return article, err
}

// 检查文章id存不存在
func CheckArticleExist(id int) bool {
  var flag int
  db, _ := Open()
  defer db.Close()

  err := db.QueryRow("SELECT id FROM articles WHERE id = ?", id).Scan(&flag)

  if err != nil {
    return false
  } else {
    return true
  }
}

// 编辑文章
func UpdateArticle(id int, title, content, tags string) error {
  db, _ := Open()
  defer db.Close()

  stmt, _ := db.Prepare("UPDATE articles SET title = ?, content = ?, tags = ? WHERE id = ?")
  _, err := stmt.Exec(title, content, tags, id)
  return err
}

// 添加文章
func AddArticle(title, content, tags string) (int64, error) {
  db, _ := Open()
  defer db.Close()

  stmt, preErr := db.Prepare("INSERT articles SET title=?, content=?, tags=?")

  if preErr != nil {
    return 0, preErr
  }

  res, execErr := stmt.Exec(title, html.EscapeString(content), tags)

  if execErr != nil {
    return 0, execErr
  }

  id, insetErr := res.LastInsertId()

  if insetErr != nil {
    return 0, insetErr
  } else {
    return id, nil
  }
}

// 删除文章
func DeleteArticle(id int) error {
  db, _ := Open()
  defer db.Close()

  stmt, _ := db.Prepare("DELETE FROM articles WHERE id=?")
  _, err := stmt.Exec(id)

  return err
}