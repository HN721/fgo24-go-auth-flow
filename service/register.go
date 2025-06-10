package service

import (
	"crypto/md5"
	"fmt"
)

type user struct {
	name     string
	email    string
	ages     int
	password string
}

var Data1 = []user{}

func Register() {
	var name string
	var email string
	var ages int
	var password string
	fmt.Print("Masukan Nama: ")
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Print("Masukan Email: ")
	fmt.Scan(&email)
	fmt.Print("Masukan Umur: ")
	fmt.Scan(&ages)
	fmt.Print("Masukan Password: ")
	fmt.Scan(&password)

	Data := user{
		name:     name,
		email:    email,
		ages:     ages,
		password: encrypt(password),
	}
	fmt.Println(password)
	Data1 = append(Data1, Data)
}
func encrypt(psswd string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(psswd))

	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
