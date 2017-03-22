package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"strings"
)

func userPolicies(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r, "user") {

		if r.Method == "GET" { // Get user informations
			r.ParseForm() // parse arguments, you have to call this by yourself
			if r.FormValue("policy") != "" {
				// Check if authorized
				var policyOwner = strings.Split(r.FormValue("policy"), "_")
				if policyOwner[1] == getUserFromAuth(w, r) {
					// Get user informations
					// Read file to byte slice
					data, err := ioutil.ReadFile("./data/" + r.FormValue("policy"))
					if err != nil {
						log.Println(err)
						fmt.Fprintf(w, err.Error())
					} else {
						fmt.Fprintf(w, string(data))
					}
				} else {
					http.NotFound(w, r)
					return
				}
			} else { // Get policies list
				files, _ := ioutil.ReadDir("./data/")
				var policies string
				for _, f := range files {
					var split = strings.Split(f.Name(), "_")
					// Only user policy, get user from basic credentials
					if split[0] == "policy" && split[1] == getUserFromAuth(w, r) {
						if policies != "" {
							policies += ","
						}
						policies += f.Name()
					}
				}
				fmt.Fprintf(w, policies)
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
