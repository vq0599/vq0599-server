package controller

import (
  "net/http"
  "io/ioutil"
  "github.com/gin-gonic/gin"
  // "fmt"
  "vq0599/conf"
  "strings"
  "vq0599/common"
  "encoding/json"
)


type StaticServerResponse struct {
  Code  string `json:"code"`
  Data  string `json:"data"`
  Msg   string `json:"msg"`
}

func requestStaticServer(c *gin.Context, path string) {
  url := conf.STATIC_DOMAIN + path

  resp, _ := common.NewRequestForward(c, url)
  defer resp.Body.Close()

  respData := &StaticServerResponse{}
  respBodyByte, _ := ioutil.ReadAll(resp.Body)
  json.NewDecoder(strings.NewReader(string(respBodyByte))).Decode(respData)

  if resp.StatusCode >= 500 {
    c.JSON(resp.StatusCode, http.StatusText(resp.StatusCode))
  } else {
    c.JSON(resp.StatusCode, respData)
  }
}

func UploadImage(c *gin.Context) {
  requestStaticServer(c, "/upload/image")
}

func UploadVideo(c *gin.Context) {
  requestStaticServer(c, "/upload/video")
}

