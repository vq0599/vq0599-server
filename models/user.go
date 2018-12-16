package models


// 检查用户email存不存在
func CheckEmailExist(email string) bool {
  var flag string
  db, _ := Open()
  defer db.Close()

  err := db.QueryRow("SELECT email FROM users WHERE email = ?", email).Scan(&flag)

  if err != nil {
    return false
  } else {
    return true
  }
}

