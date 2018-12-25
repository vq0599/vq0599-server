package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    // "fmt"
    _"github.com/go-sql-driver/mysql"
    "vq0599/models"
    "vq0599/common"
)


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
  type Params struct {
    Title string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
  }

  cG := common.Gin{C: c}
  params := &Params{}

  paramsErr := cG.Request(params)

  if paramsErr != nil {
    return
  }

  id, err := models.AddArticle(params.Title, params.Content)

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

  type Params struct {
    Title string `json:"title" binding:"required"`
    Content string `json:"content" binding:"required"`
  }

  params := &Params{}
  paramsErr := cG.Request(params)
  if paramsErr != nil {
    return
  }

  articleModels := models.Article{
    Id:       id,
    Title:    params.Title,
    Content:   params.Content,
  }

  isExist := models.CheckArticleExist(id)

  if isExist == false {
    cG.Response(http.StatusOK, common.ERROR_NOT_EXIST_ARTICLE, nil)
    return
  }

  resultErr := articleModels.UpdateArticle()

  if resultErr != nil {
    cG.Response(http.StatusOK, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, nil)
  }
}
