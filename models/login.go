package models


// 密码验证
func VerifyPassword(email string, password string) bool {
  var realPassword string
  db, _ := Open()
  defer db.Close()

  db.QueryRow("SELECT password FROM users WHERE email = ?", email).Scan(&realPassword)

  if realPassword == password {
    return true
  } else {
    return false
  }
}