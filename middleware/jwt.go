package middleware

import (
  "github.com/gin-gonic/gin"
  "vq0599/common"
  "vq0599/controller"
  "net/http"
  // "fmt"
)

func Jwt() gin.HandlerFunc {
  return func (c *gin.Context) {

    cG := common.Gin{C: c}
    isValid, token := controller.VerifyTokenWithRefresh(c)
    if isValid == false {
      cG.Response(http.StatusUnauthorized, common.ERROR_AUTHENTICATION_FAIL, nil)
      c.Abort()
      return
    }

    common.SetCookie(c, "Token", token)
    c.Next()
  }
}