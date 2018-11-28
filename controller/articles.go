package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "fmt"
    "strconv"
    _"github.com/go-sql-driver/mysql"
    "vq0599/pkg/response"
    "vq0599/models"
)


func GetArticles (c *gin.Context) {
  rG := r.Gin{C: c}
  results, err := models.GetArticles()
  if (err != nil) {
    rG.Response(http.StatusOK, r.ERROR_NOT_EXIST_ARTICLE, nil)
  } else {
    rG.Response(http.StatusOK, r.SUCCESS, results)
  }
}

func GetArticle (c *gin.Context) {
  rG := r.Gin{C: c}
  idString := c.Param("id")
  id, err := strconv.Atoi(idString)

  if (err != nil) {
    rG.Response(http.StatusBadRequest, r.INVALID_PARAMS, nil)
    return
  }

  result, err := models.GetArticle(id)

  if (err != nil) {
    fmt.Println(err)
    rG.Response(http.StatusOK, r.ERROR_NOT_EXIST_ARTICLE, nil)
    } else {
    rG.Response(http.StatusOK, r.SUCCESS, result)
  }
}

func AddArticle (c *gin.Context) {

}


func Count (c *gin.Context) {
  type Params struct {
    Id int `json:"id"`
  }
  data := &Params{}
  paramsErr := c.ShouldBindJSON(data)

  rG := r.Gin{C: c}

  if (paramsErr != nil) {
    rG.Response(http.StatusBadRequest, r.INVALID_PARAMS, nil)
    return
  }

  isExist := models.CheckArticleExist(data.Id)

  if isExist == false {
    rG.Response(http.StatusOK, r.ERROR_NOT_EXIST_ARTICLE, nil)
    return
  }

  resultErr := models.Count(data.Id)

  if resultErr != nil {
    rG.Response(http.StatusOK, r.ERROR, nil)
  } else {
    rG.Response(http.StatusOK, r.SUCCESS, nil)
  }
}
