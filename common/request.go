package common

import (
  "strconv"
  "errors"
  // "fmt"
  "vq0599/conf"
)

func (g *Gin) ScanRequestBody(params interface{}) error {
  paramsErr := g.C.ShouldBindJSON(params)

  if (paramsErr != nil) {
    g.ResponseParamError()
  }
  return paramsErr
}

func (g *Gin) GetParamFromURI(key string) (int, error) {
  idString := g.C.Param("id")
  id, err := strconv.Atoi(idString)

  if (err != nil) {
    g.ResponseParamError()
    return 0, err
  }

  return id, err
}

func (g *Gin) GetNumberQuery(key string, defaultValue int) int {
  numString := g.C.Query(key)
  numQuery, err := strconv.Atoi(numString)

  if (err != nil) {
    return defaultValue
  }
  return numQuery
}

func (g *Gin) GetPage() (int, error) {
  page := g.GetNumberQuery("page", 1)
  if page < 0 {
    g.ResponseParamError()
    return 0, errors.New("ERROR_INVALID_PARAMS")
  }

  return page, nil
}

func (g *Gin) GetPageSize() (int, error) {
  pageSize := g.GetNumberQuery("page_size", conf.DEFAULT_PER_SIZE)
  if pageSize < 0 {
    g.ResponseParamError()
    return 0, errors.New("ERROR_INVALID_PARAMS")
  }

  return pageSize, nil
}