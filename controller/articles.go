package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "fmt"
    _"github.com/go-sql-driver/mysql"
    "vq0599/models"
    "vq0599/common"
    "strings"
)

type ArticleParams struct {
  Title string `json:"title" binding:"required"`
  Source string `json:"source" binding:"required"`
  Tags []string `json:"tags" binding:"required"`
  Html string `json:"html" binding:"required"`
}


func GetArticles(c *gin.Context) {
  cG := common.Gin{C: c}
  results, err := models.GetArticles()
  if (err != nil) {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, results)
  }
}

func GetArticle(c *gin.Context) {
  cG := common.Gin{C: c}
  id, idErr := cG.GetParamFromURI("id")

  if idErr != nil {
    return
  }

  result, err := models.GetArticle(id)

  if (err != nil) {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, result)
  }
}

func AddArticle(c *gin.Context) {
  cG := common.Gin{C: c}
  params := &ArticleParams{}

  if cG.ScanRequestBody(params) != nil {
    return
  }

  tagsString := strings.Join(params.Tags, ",")
  id, err := models.AddArticle(params.Title, params.Source, params.Html, tagsString)

  if err != nil {
    cG.Response(http.StatusOK, common.ERROR_ADD_ARTICLE_FAIL, nil)
    } else {
    cG.Response(http.StatusOK, common.SUCCESS, id)
  }
}

func DeleteArticle(c *gin.Context) {
  cG := common.Gin{C: c}
  id, idErr := cG.GetParamFromURI("id")

  if idErr != nil {
    return
  }

  isExist := models.CheckArticleExist(id)

  if isExist == false {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
    return
  }

  resultErr := models.DeleteArticle(id)

  if resultErr != nil {
    cG.Response(http.StatusOK, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  }
}

func UpdateArticle(c *gin.Context) {
  cG := common.Gin{C: c}
  id, idErr := cG.GetParamFromURI("id")

  if idErr != nil {
    return
  }

  params := &ArticleParams{}
  paramsErr := cG.ScanRequestBody(params)
  if paramsErr != nil {
    return
  }

  isExist := models.CheckArticleExist(id)

  if isExist == false {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
    return
  }

  tagsString := strings.Join(params.Tags, ",")
  resultErr := models.UpdateArticle(
    id,
    params.Title,
    params.Source,
    tagsString,
    params.Html,
  )

  if resultErr != nil {
    cG.Response(http.StatusOK, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  }
}
