package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "fmt"
    "strconv"
    _"github.com/go-sql-driver/mysql"
    "vq0599/models"
    "vq0599/common"
)


func GetArticles (c *gin.Context) {
  cG := common.Gin{C: c}
  results, err := models.GetArticles()
  if (err != nil) {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, results)
  }
}

func GetArticle (c *gin.Context) {
  cG := common.Gin{C: c}
  idString := c.Param("id")
  id, err := strconv.Atoi(idString)

  if (err != nil) {
    cG.Response(http.StatusBadRequest, common.INVALID_PARAMS, nil)
    return
  }

  result, err := models.GetArticle(id)

  if (err != nil) {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, result)
  }
}

func AddArticle (c *gin.Context) {
  type Params struct {
    Title string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
  }

  cG := common.Gin{C: c}
  params := &Params{}

  paramsErr := cG.Request(params)

  if (paramsErr == nil) {
    cG.Response(http.StatusOK, common.SUCCESS, params)
  }
}


func Count (c *gin.Context) {
  type Params struct {
    Id int `json:"id" binding:"required"`
  }
  params := &Params{}
  cG := common.Gin{C: c}

  paramsErr := cG.Request(params)

  if paramsErr != nil {
    return
  }

  isExist := models.CheckArticleExist(params.Id)

  if isExist == false {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
    return
  }

  resultErr := models.Count(params.Id)

  if resultErr != nil {
    cG.Response(http.StatusOK, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  }
}
