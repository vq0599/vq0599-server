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

// 验证TOKEN是否合法(JWT验证 + 数据库验证)
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

// 判断TOKEN是否需要刷新并自动更新TOKEN
func refreshToken(token string) (bool, string) {
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

// 任何更新资源的操作之前，完整的走一遍Token流程
func VerifyTokenWithRefresh(c *gin.Context) (bool, string) {
  cookie, cookieErr := c.Request.Cookie("Token")

  if cookieErr != nil {
    return false, ""
  }

  token := cookie.Value
  isValidToken := verifyToken(token)

  if isValidToken == true {
    _, newToken := refreshToken(token)
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
    models.UpdateToken(token, id)
    setCookie(c, "Token", token)
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
      setCookie(c, "Token", token)
    }
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.ERROR_NOT_LOGGED, nil)
  }
}
