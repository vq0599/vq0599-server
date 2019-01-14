package common

import (
  "net/http"
  "io/ioutil"
	"github.com/gin-gonic/gin"
  "bytes"
  "strings"
  "regexp"
)

func NewRequestForward(c *gin.Context, url string) (*http.Response, error) {
  requestBody, _ := ioutil.ReadAll(c.Request.Body)

  proxyRequest, _ := http.NewRequest(c.Request.Method, url, bytes.NewReader(requestBody))
  proxyRequest.Header = c.Request.Header

  return http.DefaultClient.Do(proxyRequest)
}

// golang 切割空字符串依然会得到包含一个空元素的数组 => [""]
func Split(str, sep string) []string {
  if str != "" {
    return strings.Split(str, sep)
  } else {
    return make([]string, 0)
  }
}

func HtmlToPureText (html string) string {
  reHTML, _ := regexp.Compile("<[^>]*>")
  return reHTML.ReplaceAllString(html, "")
}

func SubString (src string, start, end int) string {
  return string([]rune(src)[start: end])
}
