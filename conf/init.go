package conf

import (
  "time"
)

const (
  DB_HOST         = "127.0.0.1:3306"
  DB_USER         = "root"
  DB_TYPE         = "mysql"
  DB_PASSWORD     = "******"
  DB_NAME         = "blog"

  SERVER_PORT     = "8180"

  JWT_ISSUER      = "vq0599"
  JWT_SECRET      = "******"
  // 24 小时有效期
  JWT_MAXAGE      = 24 * time.Hour
  // 1小时刷新TOKEN
  JWT_REFRESH     = 1 * time.Hour

  COOKIES_DOMAIN  = "vq0599.com"
  COOKIES_MAXAGE  = 7 * 24 * 3600

  STATIC_DOMAIN   = "https://static.vq0599.com"

  DEFAULT_PER_SIZE = 10
)