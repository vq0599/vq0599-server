# 基于Golang Gin的RESTful API 实践


## 鉴权体系
- 登录 => 登录成功后生成TOKEN(JSON Web Token)
  - setCookie，有效期7天
  - TOKEN入库

- 鉴权 => 判断TOKEN创建时间
  - 创建时间 < 1小时 => Next
  - 创建时间 > 1小时 => TOKEN VS 数据库 & 更新TOKEN
