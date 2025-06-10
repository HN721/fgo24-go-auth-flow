package main

import (
	"auth/service"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var Choice int

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
func main() {

	for {
		ClearScreen()
		fmt.Println("Selamat Datang Di Aplikasi Koda Uhuy")
		fmt.Println("----------Pilih Menu------")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Forgot Password")
		fmt.Println("0. Exit")
		fmt.Print("Pilih: ")
		fmt.Scanln(&Choice)

		switch Choice {
		case 1:
			service.Login()
		case 2:
			service.Register()
		case 3:
			service.NewPassword()
		case 0:
			fmt.Println("Keluar dari aplikasi...")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid")
		}

		fmt.Println("Tekan ENTER untuk melanjutkan...")
		fmt.Scanln()
	}

}
