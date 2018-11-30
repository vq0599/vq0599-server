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
  idString := c.Param("id")
  id, err := strconv.Atoi(idString)

  if (err != nil) {
    cG.Response(http.StatusBadRequest, common.INVALID_PARAMS, nil)
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
  type Params struct {
    Id int `json:"id" binding:"required"`
    Title string `json:"title"`
    Content string `json:"content"`
  }

  params := &Params{}
  cG := common.Gin{C: c}

  paramsErr := cG.Request(params)

  if paramsErr != nil {
    return
  }

  articleModels := models.Article{
    Id:     params.Id,
    Title:     params.Title,
    Content:     params.Content,
  }

  isExist := models.CheckArticleExist(params.Id)

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
