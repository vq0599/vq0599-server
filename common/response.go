package common

import (
  "net/http"
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

func (g *Gin) ResponseParamError() {
  g.C.JSON(http.StatusBadRequest, gin.H{
    "code": ERROR_INVALID_PARAMS,
    "msg": GetMsg(ERROR_INVALID_PARAMS),
    "data": nil,
  })
}
