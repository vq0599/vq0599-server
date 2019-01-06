package common

import (
  "net/http"
  "io/ioutil"
	"github.com/gin-gonic/gin"
  "bytes"
)

func NewRequestForward(c *gin.Context, url string) (*http.Response, error) {
  requestBody, _ := ioutil.ReadAll(c.Request.Body)

  proxyRequest, _ := http.NewRequest(c.Request.Method, url, bytes.NewReader(requestBody))
  proxyRequest.Header = c.Request.Header

  return http.DefaultClient.Do(proxyRequest)
}