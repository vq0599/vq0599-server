package controller

import (
  "net/http"
  "github.com/gin-gonic/gin"
  // "fmt"
  // _"github.com/go-sql-driver/mysql"
  "vq0599/models"
  "vq0599/common"
  // "strings"
  "strconv"
)

func UpdatePvs(c *gin.Context) {
  idString := c.Param("id")
  id, err := strconv.Atoi(idString)

  if err != nil {
    return
  }

  if models.UpdatePvs(id) != nil {
    return
  }

  c.Header("Content-Type", "image/gif")
  c.Writer.WriteHeader(http.StatusOK)
}

func UpdateLikes(c * gin.Context) {
  cG := common.Gin{C: c}
  id, idErr := cG.GetParamFromURI("id")

  if idErr != nil {
    return
  }

  resultErr := models.UpdateLikes(id)

  if resultErr != nil {
    cG.Response(http.StatusOK, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  }
}