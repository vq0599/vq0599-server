package controller
import (
  "net/http"
  "github.com/gin-gonic/gin"
  // "fmt"
  // _"github.com/go-sql-driver/mysql"
  "vq0599/models"
  // "fmt"
  "vq0599/common"
)

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

  loginStatus := models.VerifyPassword(params.Email, params.Password)

  if loginStatus == true {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  } else {
    cG.Response(http.StatusOK, common.ERROR_LOGIN_PASSWORD, nil)
  }
}