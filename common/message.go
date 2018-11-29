package common

var msgFlags = map[int]string{
  SUCCESS:                         "SUCCESS",
  ERROR:                           "FAIL",
  INVALID_PARAMS:                  "请求参数错误",
  ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
  ERROR_ADD_ARTICLE_FAIL:          "添加文章失败",
}


func GetMsg(code int) string {
  msg, ok := msgFlags[code]
  if ok {
    return msg
  }

  return msgFlags[ERROR]
}