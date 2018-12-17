package controller
import (
  "net/http"
  "net/url"
  "github.com/gin-gonic/gin"
  "strconv"
  // "fmt"
  "time"
  "vq0599/models"
  "vq0599/common"
  "vq0599/util"
  "vq0599/conf"
)

func setCookie(c *gin.Context, name, value string) {
  http.SetCookie(c.Writer, &http.Cookie{
    Name:     name,
    Value:    url.QueryEscape(value),
    MaxAge:   conf.COOKIES_MAXAGE,
    Path:     "/",
    Domain:   conf.COOKIES_DOMAIN,
    Secure:   false,
    HttpOnly: true,
  })
}

// 验证TOKEN是否合法
func verifyToken(token string) bool {
  claims, TokenErr := util.ParseToken(token)

  if TokenErr != nil {
    return false
  }

  idString := claims.Id
  id, idErr := strconv.Atoi(idString)

  if idErr != nil {
    return false
  }

  return models.VerifyToken(id, token)
}

// 验证TOKEN剩余有效期，并判断是否刷新
func verifyTokenRefresh(token string) (bool, string) {
  // 剩余有效时间
  claims, _ := util.ParseToken(token)
  id, _ := strconv.Atoi(claims.Id)

  validTime := (claims.ExpiresAt - time.Now().Unix())

  if validTime < int64((conf.JWT_MAXAGE - conf.JWT_REFRESH) / time.Second) {
    newToken, _ := util.GenerateToken(id)
    models.UpdateToken(newToken, id)
    return true, newToken
  }

  return false, ""
}

// 验证登录账户密码
func Login(c *gin.Context) {
  type Params struct {
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
  }

  cG := common.Gin{C: c}
  params := &Params{}

  paramsErr := cG.Request(params)

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
    models.UpdateToken(token, id)
    setCookie(c, "Token", token)
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.ERROR_LOGIN_PASSWORD, nil)
  }
}

// 获取登录状态
func GetLoginStatus(c *gin.Context) {
  cG := common.Gin{C: c}
  cookie, _ := c.Request.Cookie("Token")
  token := cookie.Value

  isValidToken := verifyToken(token)

  if isValidToken == true {
    isRefresh, newToken := verifyTokenRefresh(token)
    if (isRefresh) {
      setCookie(c, "Token", newToken)
      cG.Response(http.StatusOK, common.SUCCESS, newToken)
      return
    }
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.NOT_LOGGED, nil)
  }
}