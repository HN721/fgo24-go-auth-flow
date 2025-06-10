package utils

import (
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
func User() {
	ClearScreen()
	fmt.Println("Login Berhasil")
	fmt.Println("----------Pilih Menu------")
	fmt.Println("1. UserList")
	fmt.Println("3. Forgot Password")
	fmt.Println("0. Exit")
	fmt.Print("Pilih: ")
	fmt.Scanln(&Choice)
}
