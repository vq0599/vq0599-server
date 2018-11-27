package models

import (
  "database/sql"
  _"github.com/go-sql-driver/mysql"
  "fmt"
  "vq0599/conf"
)


func Open() (*sql.DB, error) {
  db, err := sql.Open(conf.DB_TYPE, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
    conf.DB_USER,
    conf.DB_PASSWORD,
    conf.DB_HOST,
    conf.DB_NAME,
  ))

  return db, err
}
