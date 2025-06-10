package main

import (
	"auth/service"
	"fmt"
	"os"
)

var choice int

func main() {
	fmt.Println("Selamat Datang Di Aplikasi Koda Uhuy")
	fmt.Println("----------Pilih Menu------")
	fmt.Println("1.Login")
	fmt.Println("2.Register")
	fmt.Println("3.Forgot Password")
	fmt.Println("0. Exit")
	fmt.Scanln(&choice)
	if choice == 1 {
		service.Register()
	}
	if choice == 2 {
		service.Login()
	}
	defer os.Exit(0)
}
