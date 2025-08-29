package main

import (
	"fmt"
	"github.com/caozhong11/go-zero-test/client/user"
)

func main() {
	fmt.Println("this is a server ")
	request := user.UserRequest{}
	fmt.Printf("%+v\n", request)
}
