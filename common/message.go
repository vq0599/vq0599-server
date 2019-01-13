package common

var msgFlags = map[string]string{
  SUCCESS:                         "成功",
  ERROR:                           "未知错误",
  ERROR_INVALID_PARAMS:            "请求参数错误",
  ERROR_NOT_EXIST_ARTICLE:         "该文章不存在",
  ERROR_ADD_ARTICLE_FAIL:          "添加文章失败",
  ERROR_GET_ARTICLE_FAIL:          "获取文章失败",
  ERROR_NOT_EXIST_USER:            "用户不存在",
  ERROR_PASSWORD_FAIL:             "密码错误",
  ERROR_NOT_LOGGED:                "未登录",
  ERROR_AUTHENTICATION_FAIL:       "鉴权失败",
}

func GetMsg(code string) string {
  msg, ok := msgFlags[code]
  if ok {
    return msg
  }

  return msgFlags[ERROR]
}