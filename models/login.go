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