package service

import "fmt"

type Profile interface {
	HandleLogin() string
	Forgot() int
}
type LoginData struct {
	email    string
	password string
}

func (a LoginData) HandleLogin() {
	if a.email == Data1[0].email && a.password == Data1[0].password {
		fmt.Println("Login Berhasil")
	} else {
		fmt.Println("Login Failed Wrong email or Password")
	}

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
		password: password,
	}
	defer result.HandleLogin()
}
