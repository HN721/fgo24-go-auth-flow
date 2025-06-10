package main

import (
	"auth/service"
	"fmt"
)

func main() {
	service.Register()
	fmt.Println("Harap Login")
	defer service.Login()
}
