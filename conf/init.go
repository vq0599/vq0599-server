package conf

import (
  "time"
)

const (
  DB_HOST         = "127.0.0.1:3306"
  DB_USER         = "root"
  DB_TYPE         = "mysql"
  DB_PASSWORD     = "ceoms999"
  DB_NAME         = "blog"

  SERVER_PORT     = "8180"

  JWT_SECRET      = "golang"
  // 24 小时有效期
  JWT_MAXAGE      = 24 * time.Hour
  // 1小时刷新TOKEN
  JWT_REFRESH     = 1 * time.Hour

  COOKIES_DOMAIN  = "vq0599.xyz"
  COOKIES_MAXAGE  = 7 * 24 * 3600

  STATIC_DOMAIN   = "http://static.vq0599.xyz"
)