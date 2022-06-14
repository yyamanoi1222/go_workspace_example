package main

import (
  "fmt"
  "log"
  "github.com/yyamanoi1222/go_workspace_example/modules/auth"
  "github.com/yyamanoi1222/go_workspace_example/modules/user"
)

func main() {
  user := user.User{Name: "test", ID: "1"}
  token, err := auth.GenerateJwt(user)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("token is %s", token)
}
