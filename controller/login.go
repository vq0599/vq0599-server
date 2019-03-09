package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
  // "fmt"
  "time"
  "vq0599/models"
  "vq0599/common"
  "vq0599/util"
  "vq0599/conf"
)

// 任何更新资源的操作之前，完整的走一遍Token流程
func VerifyTokenWithRefresh(c *gin.Context) (bool, string) {
  cookie, cookieErr := c.Request.Cookie("Token")

  if cookieErr != nil {
    return false, ""
  }

  token := cookie.Value

  claims, parseResult := util.ParseToken(token)

  if parseResult == false {
    return false, ""
  }

  // 是否需要刷新
  isRefreshToken := (time.Now().Unix() - claims.IssuedAt) > int64(conf.JWT_REFRESH / time.Second)

  if isRefreshToken == false {
    return true, ""
  }

  id, _ := strconv.Atoi(claims.Id)
  newToken, _ := util.GenerateToken(id)

  return true, newToken
}

// 验证登录账户密码
func Login(c *gin.Context) {
  type Params struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
  }

  cG := common.Gin{C: c}
  params := &Params{}

  paramsErr := cG.ScanRequestBody(params)

  if paramsErr != nil {
    return
  }

  isExist := models.CheckEmailExist(params.Email)

  if isExist == false {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_USER, nil)
    return
  }

  loginStatus, id := models.VerifyPassword(params.Email, params.Password)

  if loginStatus == true {
    token, _ := util.GenerateToken(id)
    util.SetCookie(c, "Token", token)
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.ERROR_PASSWORD_FAIL, nil)
  }
}

func GetLoginStatus(c *gin.Context) {
  cG := common.Gin{C: c}
  isValid, token := VerifyTokenWithRefresh(c)

  if isValid == true {
    if token != "" {
      util.SetCookie(c, "Token", token)
    }
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.ERROR_NOT_LOGGED, nil)
  }
}
