package controller

import (
  "net/http"
  "net/url"
  "path"
  "github.com/gin-gonic/gin"
  // "fmt"
  "vq0599/conf"
  "vq0599/common"
  "github.com/tencentyun/cos-go-sdk-v5"
  "github.com/teris-io/shortid"
)

func UploadImage(c *gin.Context) {
  cG := common.Gin{C: c}
  file, header, ParamErr := c.Request.FormFile("file")

  if ParamErr != nil {
    cG.ResponseParamError()
    return
  }

  bucketUrl, _ := url.Parse(conf.COS_API_DOMAIN)
  client := cos.NewClient(&cos.BaseURL{BucketURL: bucketUrl}, &http.Client{
    Transport: &cos.AuthorizationTransport{
      SecretID:  conf.COS_SECRET_ID,
      SecretKey: conf.COS_SECRET_KEY,
    },
  })

  sid, _ := shortid.New(conf.SHORT_WORKER, shortid.DefaultABC, conf.SHORT_SEED)
  key, _ := sid.Generate()
  suffix := path.Ext(header.Filename)
  filePath := "images/"+ key + suffix

  resp, err := client.Object.Put(c, filePath, file, nil)

  if err != nil {
    cG.Response(http.StatusInternalServerError, common.ERROR, nil)
  } else if resp.StatusCode != http.StatusOK {
    cG.Response(resp.StatusCode, common.ERROR, nil)
  } else {
    cG.Response(http.StatusOK, common.SUCCESS, conf.COS_STATIC_DOMAIN + "/" + filePath)
  }
}