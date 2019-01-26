package util

import (
  "github.com/dgrijalva/jwt-go"
  "time"
  "vq0599/conf"
  "strconv"
  // "fmt"
)

var jwtSecret = []byte(conf.JWT_SECRET)

func GenerateToken(id int) (string, error) {
  createTime := time.Now()
  expireTime := createTime.Add(conf.JWT_MAXAGE)
  idString := strconv.Itoa(id)

  claims := &jwt.StandardClaims{
    Id: idString,
    ExpiresAt: expireTime.Unix(),
    Issuer: conf.JWT_ISSUER,
    IssuedAt: createTime.Unix(),
  }

  tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  token, err := tokenClaims.SignedString(jwtSecret)

  return token, err
}

func ParseToken(tokenString string) (*jwt.StandardClaims, bool) {
  token, _ := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
    return jwtSecret, nil
  })

  claims, ok := token.Claims.(*jwt.StandardClaims);

  if ok && token.Valid {
    return claims, true
  }

  return nil, false
}
