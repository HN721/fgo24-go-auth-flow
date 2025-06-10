package service

import (
	"auth/utils"
	"fmt"
)

type Profile interface {
	HandleLogin()
	Forgot()
}
type LoginData struct {
	email    string
	password string
}
type newPassword struct {
	email    string
	password string
}

var choose int

func (a LoginData) HandleLogin() {
	found := false
	for _, user := range Data1 {
		if a.email == user.email && a.password == user.password {
			fmt.Println("Login Berhasil")
			found = true
			utils.User()
		}
	}
	if !found {
		fmt.Println("Login Failed: Wrong email or password")
		fmt.Println("Forgot your Password? Tekan 1 untuk reset:")
		fmt.Scan(&choose)
		if choose == 1 {
			NewPassword()
		}
	}

}
func (s newPassword) Forgot() string {
	var result string
	found := false
	for x := range Data1 {
		if Data1[x].email == s.email {
			Data1[x].password = s.password
			result = s.password
			fmt.Println("Password berhasil diubah!")
			found = true
		}
	}
	if !found {
		fmt.Println("Harap masukkan email yang benar")
	}
	return result

}
func Login() {
	fmt.Println("-----Login----")
	var email string

	var password string

	fmt.Print("Masukan Email: ")
	fmt.Scan(&email)
	fmt.Print("Masukan Password: ")
	fmt.Scan(&password)

	result := LoginData{
		email:    email,
		password: encrypt(password),
	}
	result.HandleLogin()

}
func NewPassword() {
	var password string
	var email string
	fmt.Println("Masukan Email:")
	fmt.Scan(&email)
	fmt.Print("Masukan Password: ")
	fmt.Scan(&password)
	result := &newPassword{
		email:    email,
		password: encrypt(password),
	}
	result.Forgot()
}
