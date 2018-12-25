package common

import (
  "github.com/gin-gonic/gin"
)

func (g *Gin) Response(httpCode int, errCode string, data interface{}) {
  g.C.JSON(httpCode, gin.H{
    "code": errCode,
    "msg": GetMsg(errCode),
    "data": data,
  })

  return
}