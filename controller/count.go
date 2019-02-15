package controller

import (
  "net/http"
  "github.com/gin-gonic/gin"
  // "fmt"
  "vq0599/models"
  "vq0599/common"
)

func UpdatePvs(c *gin.Context) {
  cG := common.Gin{C: c}
  id, idErr := cG.GetParamFromURI("id")

  if idErr != nil {
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