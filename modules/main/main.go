package main

import (
  "fmt"
  "github.com/yyamanoi1222/go_workspace_example/modules/user"
)

func main() {
  user := user.User{Name: "test"}

  fmt.Printf("Hello world %v \n", user)
}
