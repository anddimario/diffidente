package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"os"

	"strings"
)

func adminPolicies(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r, "admin") {

		if r.Method == "GET" { // Get user informations
			r.ParseForm() // parse arguments, you have to call this by yourself
			if r.FormValue("policy") != "" {
				// Get user informations
				// Read file to byte slice
				data, err := ioutil.ReadFile("./data/" + r.FormValue("policy"))
				if err != nil {
					log.Println(err)
					fmt.Fprintf(w, err.Error())
				} else {
					fmt.Fprintf(w, string(data))
				}
			} else { // Get policies list
				files, _ := ioutil.ReadDir("./data/")
				var policies string
				for _, f := range files {
					var split = strings.Split(f.Name(), "_")
					if split[0] == "policy" {
						if policies != "" {
							policies += ","
						}
						policies += f.Name()
					}
				}
				fmt.Fprintf(w, policies)
			}

		} else if r.Method == "POST" { // Add new user
			r.ParseForm() // parse arguments, you have to call this by yourself
			// Store user informations
			err := ioutil.WriteFile("./data/policy_"+r.FormValue("username")+"_"+r.FormValue("app"), []byte(r.FormValue("keys")), 0666)
			if err != nil {
				log.Println(err)
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "added")
			}
		} else if r.Method == "DELETE" { // Remove user
			err := os.Remove("./data/" + r.FormValue("policy"))
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
