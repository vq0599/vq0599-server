package util

import (
  "github.com/dgrijalva/jwt-go"
  "time"
  "vq0599/conf"
  "strconv"
)

var jwtSecret = []byte(conf.JWT_SECRET)

type Claims struct {
  Id string `json:"id"`
  jwt.StandardClaims
}

func GenerateToken(id int) (string, error) {
  expireTime := time.Now().Add(conf.JWT_MAXAGE)
  idString := strconv.Itoa(id)

  claims := Claims{
    idString,
    jwt.StandardClaims{
      ExpiresAt: expireTime.Unix(),
      Issuer:    "vq0599",
    },
  }

  tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  token, err := tokenClaims.SignedString(jwtSecret)

  return token, err
}

func ParseToken(token string) (*Claims, error) {
  tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
    return jwtSecret, nil
  })

  if tokenClaims != nil {
    if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
      return claims, nil
    }
  }

  return nil, err
}
