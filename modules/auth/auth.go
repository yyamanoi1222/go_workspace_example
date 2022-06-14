package auth

import (
  jwt "github.com/golang-jwt/jwt/v4"
  "github.com/yyamanoi1222/go_workspace_example/modules/user"
)

func GenerateJwt(user user.User) (tokenString string, err error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "id": user.ID,
  })

  tokenString, err = token.SignedString([]byte("test"))
  return
}
