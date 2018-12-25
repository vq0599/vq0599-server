package common

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
)

func (g *Gin) ScanRequestBody(params interface{}) error {
  paramsErr := g.C.ShouldBindJSON(params)

  if (paramsErr != nil) {
    g.C.JSON(http.StatusBadRequest, gin.H{
      "code": ERROR_INVALID_PARAMS,
      "msg": GetMsg(ERROR_INVALID_PARAMS),
      "data": nil,
    })
  }
  return paramsErr
}

func (g *Gin) GetParamFromURI(key string) (int, error) {
  idString := g.C.Param("id")
  id, err := strconv.Atoi(idString)

  if (err != nil) {
    g.C.JSON(http.StatusBadRequest, gin.H{
      "code": ERROR_INVALID_PARAMS,
      "msg": GetMsg(ERROR_INVALID_PARAMS),
      "data": nil,
    })

    return 0, err
  }

  return id, err
}