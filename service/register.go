package service

import (
	"crypto/md5"
	"fmt"
	"os"
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
	fmt.Println("-----------Register------------")
	fmt.Print("Masukan Nama: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("Nama Tidak boleh Kosong")
		os.Exit(0)
	}
	fmt.Println(name)
	fmt.Print("Masukan Email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		fmt.Println("Email Tidak boleh Kosong")
		os.Exit(0)
	}
	fmt.Print("Masukan Umur: ")
	_, err = fmt.Scan(&ages)
	if err != nil {
		fmt.Println("Umur Tidak boleh Kosong")
		os.Exit(0)
	}
	fmt.Print("Masukan Password: ")
	_, err = fmt.Scan(&password)
	if err != nil {
		fmt.Println("Password Tidak boleh Kosong")
		os.Exit(0)
	}
	Data := &user{
		name:     name,
		email:    email,
		ages:     ages,
		password: encrypt(password),
	}
	fmt.Println(password)
	Data1 = append(Data1, *Data)
	fmt.Println("Sucessfully Registered")
	fmt.Print(Data1)
	Login()
}
func encrypt(psswd string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(psswd))

	md5 := hash.Sum(nil)
	return fmt.Sprintf("%x", md5)
}
