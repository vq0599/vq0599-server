package models

import (
  "time"
  "vq0599/common"
  // "fmt"
)

type Article struct {
  Id            int      `json:"id"`
  Title         string   `json:"title"`
  Source        string   ` json:"source"`
  Create_time   int64    `json:"create_time"`
  Pv            int      `json:"pv"`
  Tags          []string `json:"tags"`
  Summary       string   `json:"summary"`
  Html          string   `json:"html"`
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
    rows.Scan(&article.Id, &article.Title, &article.Source, &create_time, &article.Pv, &tags, &article.Html)

    summaryHtml := common.Split(article.Html, "<!-- more -->")[0]
    article.Summary = common.HtmlToPureText(summaryHtml)
    article.Tags = common.Split(tags, ",")
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
    &article.Source,
    &create_time,
    &article.Pv,
    &tags,
    &article.Html,
  )

  article.Tags = common.Split(tags, ",")
  article.Create_time = create_time.UnixNano() / 1e6

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
func UpdateArticle(id int, title, source, tags, html string) error {
  db, _ := Open()
  defer db.Close()

  stmt, preErr := db.Prepare("UPDATE articles SET title = ?, source = ?, tags = ?, html = ? WHERE id = ?")

  if preErr != nil {
    return preErr
  }

  _, err := stmt.Exec(title, source, tags, html, id)

  return err
}

// 添加文章
func AddArticle(title, source, html, tags string) (int64, error) {
  db, _ := Open()
  defer db.Close()

  stmt, preErr := db.Prepare("INSERT articles SET title=?, source=?, html=?, tags=?")

  if preErr != nil {
    return 0, preErr
  }

  res, execErr := stmt.Exec(title, source, html, tags)

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