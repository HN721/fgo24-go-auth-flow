package main

import (
	"auth/service"
	"fmt"
	"os"
)

var Choice int

func main() {

	fmt.Println("Selamat Datang Di Aplikasi Koda Uhuy")
	fmt.Println("----------Pilih Menu------")
	fmt.Println("1.Login")
	fmt.Println("2.Register")
	fmt.Println("0. Exit")
	fmt.Scanln(&Choice)
	if Choice == 1 {
		service.Login()

	}
	if Choice == 2 {
		service.Register()

	}
	defer os.Exit(0)
}
