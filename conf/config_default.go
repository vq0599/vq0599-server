package conf

import (
  "time"
)

const (
  // database
  DB_HOST     = "127.0.0.1:3306"
  DB_USER     = "root"
  DB_TYPE     = "mysql"
  DB_NAME     = ""
  DB_PASSWORD = ""

  // server
  SERVER_PORT = "8180"

  // json web token
  JWT_ISSUER  = ""
  JWT_SECRET  = ""
  JWT_MAXAGE  = 24 * time.Hour
  JWT_REFRESH = 1 * time.Hour

  // cookies
  COOKIES_DOMAIN = "vq0599.com"
  COOKIES_MAXAGE = 7 * 24 * 3600

  // pagination
  DEFAULT_PER_SIZE = 10

  // tencent oss
  COS_BUCKET_NAME   = ""
  COS_APP_ID        = ""
  COS_REGION        = ""
  COS_SECRET_ID     = ""
  COS_SECRET_KEY    = ""
  COS_DOMAIN        = ""
  COS_API_DOMAIN    = ""
  COS_STATIC_DOMAIN = "https://static.vq0599.com"

  // shortid https://github.com/teris-io/shortid
  SHORT_SEED   = 0
  SHORT_WORKER = 0
)