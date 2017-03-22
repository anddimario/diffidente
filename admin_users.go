package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"

	"strings"

	"golang.org/x/crypto/bcrypt"
)

func users(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r, "admin") {

		if r.Method == "GET" { // Get user informations
			r.ParseForm() // parse arguments, you have to call this by yourself
			if r.FormValue("username") != "" {
				// Get user informations
				// Read file to byte slice
				data, err := ioutil.ReadFile("./data/user_" + r.FormValue("username"))
				if err != nil {
					log.Println(err)
					fmt.Fprintf(w, err.Error())
				} else {
					fmt.Fprintf(w, string(data))
				}
			} else { // Get user list
				files, _ := ioutil.ReadDir("./data/")
				var users string
				for _, f := range files {
					var split = strings.Split(f.Name(), "_")
					if split[0] == "user" {
						if users != "" {
							users += ","
						}
						users += split[1]
					}
				}
				fmt.Fprintf(w, users)
			}

		} else if r.Method == "POST" { // Add new user
			r.ParseForm()                               // parse arguments, you have to call this by yourself
			password := []byte(r.FormValue("password")) // transform string in byte for bcrypt
			// Hashing the password with the default cost of 10
			hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
			}
			// Store user informations
			err = ioutil.WriteFile("./data/user_"+r.FormValue("username"), []byte("PASSWORD:"+string(hashedPassword)+"\n"), 0666)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "added")
			}
		} else if r.Method == "DELETE" { // Remove user
			err := os.Remove("./data/user_" + r.FormValue("username"))
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "deleted")
			}
		} else {
			http.NotFound(w, r)
			return
		}
	} else {
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	}
}
