package models

func UpdatePvs(id int) error {
  db, _ := Open()
  defer db.Close()
  stmt, preErr := db.Prepare("UPDATE articles SET pv = pv + 1 WHERE id = ?")
  
  if preErr != nil {
    return preErr
  }

  _, err := stmt.Exec(id)

  return err
}


func UpdateLikes(id int) error {
  db, _ := Open()
  defer db.Close()
  stmt, preErr := db.Prepare("UPDATE articles SET like = like + 1 WHERE id = ?")

  if preErr != nil {
    return preErr
  }

  _, err := stmt.Exec(id)

  return err
}