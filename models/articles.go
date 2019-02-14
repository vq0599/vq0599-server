package models

import (
  "time"
  "vq0599/common"
  "sort"
  // "fmt"
)

type Article struct {
  Id            int      `json:"id"`
  Title         string   `json:"title"`
  Create_time   int64    `json:"create_time"`
  Tags          []string `json:"tags"`
  Summary       string   `json:"summary,omitempty"`
  Source        string   `json:"source,omitempty"`
  Html          string   `json:"html,omitempty"`
  Read          int      `json:"read_number"`
  Like          int      `json:"like_number"`
}

type Articles []Article

// 排序三部曲，略麻烦
func (a Articles) Len() int { return len(a) }

func (a Articles) Less(i, j int) bool {
  return a[i].Create_time > a[j].Create_time
}

func (a Articles) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 获取文章列表
func GetArticles(admin bool) (Articles, error) {
  var results Articles
  var html string
  var tags string
  var create_time time.Time

  db, _ := Open()
  defer db.Close()

  rows, err := db.Query("SELECT id, title, create_time, html, tags, read_number, like_number FROM articles")
  defer rows.Close()

  if err == nil {
    for rows.Next() {
      var article Article
      rows.Scan(&article.Id, &article.Title, &create_time, &html, &tags, &article.Read, &article.Like)

      if admin == false {
        summaryHtml := common.Split(html, "<!-- more -->")[0]
        article.Summary = common.HtmlToPureText(summaryHtml)
      }
      article.Tags = common.Split(tags, ",")
      article.Create_time = create_time.UnixNano() / 1e6
  
      results = append(results, article)
    }
  }

  sort.Sort(results)
  return results, err
}

// 获取单篇文章
func GetArticle(id int, admin bool) (Article, error) {
  var article Article
  var create_time time.Time
  var tags string
  var source string
  var err error

  db, _ := Open()
  defer db.Close()

  err = db.QueryRow("SELECT id, title, create_time, tags, source, html, read_number, like_number FROM articles WHERE id = ?", id).Scan(
    &article.Id,
    &article.Title,
    &create_time,
    &tags,
    &source,
    &article.Html,
    &article.Read,
    &article.Like,
  )

  if err == nil {
    article.Tags = common.Split(tags, ",")
    article.Create_time = create_time.UnixNano() / 1e6

    if admin == true {
      article.Source = source
    }
  }

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