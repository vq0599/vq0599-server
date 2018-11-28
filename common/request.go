package common

import (
  "net/http"
	"github.com/gin-gonic/gin"
)

func (g *Gin) Request(params interface{}) error {
  paramsErr := g.C.ShouldBindJSON(params)

  if (paramsErr != nil) {
    g.C.JSON(http.StatusBadRequest, gin.H{
      "code": INVALID_PARAMS,
      "msg": GetMsg(INVALID_PARAMS),
      "data": nil,
    })
  }
  return paramsErr
}