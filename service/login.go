package service

import (
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
	if len(Data1) >= 1 {
		for x := range Data1 {
			if a.email == Data1[x].email && a.password == Data1[x].password {
				fmt.Println("Login Berhasil")
			} else {
				fmt.Println("Login Failed Wrong email or Password")
				fmt.Println("Forgot your Password?")
				fmt.Scan(&choose)
				if choose == 1 {
					NewPassword()
				}
			}
		}

	} else {
		fmt.Println("Silahkan Registerasi Terlebih Dahulu")
		Register()
	}

}
func (s newPassword) Forgot() string {

	for x := range Data1 {
		if Data1[x].email == s.email {
			Data1[x].password = s.password

		} else {
			fmt.Println("Harap Masukan Email Yang benars")
		}
	}
	return s.password
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
	defer Login()
}
