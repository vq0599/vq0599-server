package common

import (
  "net/http"
	"github.com/gin-gonic/gin"
)

func (g *Gin) Request(params interface{}) error {
  paramsErr := g.C.ShouldBindJSON(params)

  if (paramsErr != nil) {
    // StatusBadRequest 浏览器无法获取到 400 的response
    g.C.JSON(http.StatusOK, gin.H{
      "code": INVALID_PARAMS,
      "msg": GetMsg(INVALID_PARAMS),
      "data": nil,
    })
  }
  return paramsErr
}