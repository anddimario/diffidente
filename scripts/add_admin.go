package main

import (
	"fmt"
	"io/ioutil"

	"os"

	"golang.org/x/crypto/bcrypt"
)

// Run: go run add_admin.go username password
func main() {
	password := []byte(os.Args[2]) // transform string in byte for bcrypt
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	// Store user informations
	err = ioutil.WriteFile("../data/admin_"+os.Args[1], []byte("PASSWORD:"+string(hashedPassword)+"\n"), 0666)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added")
	}
}
