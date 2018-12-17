package models

import (
  // "fmt"
)

// 密码验证
func VerifyPassword(email, password string) (bool, int) {
  var realPassword string
  var id int

  db, _ := Open()
  defer db.Close()

  db.QueryRow("SELECT id, password FROM users WHERE email = ?", email).Scan(&id, &realPassword)

  if realPassword == password {
    return true, id
  } else {
    return false, id
  }
}

// 验证TOKEN是否有效
func VerifyToken(id int, token string) bool {
  var realToken string

  db, _ := Open()
  defer db.Close()

  db.QueryRow("SELECT token FROM users WHERE id = ?", id).Scan(&realToken)

  if realToken == token {
    return true
  } else {
    return false
  }
}

// 更新TOKEN
func UpdateToken(token string, id int) {
  db, _ := Open()
  defer db.Close()

  stmt, _ := db.Prepare("UPDATE users SET token = ? WHERE id = ?")
  stmt.Exec(token, id)
}