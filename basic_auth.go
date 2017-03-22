package main

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func checkAuth(w http.ResponseWriter, r *http.Request, role string) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	// Get informations from file
	data, err := ioutil.ReadFile("./data/" + role + "_" + pair[0])
	if err != nil {
		return false
	} else {
		fields := strings.Split(string(data), ":")
		storedPassword := strings.Replace(fields[1], "\n", "", 1)
		err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(pair[1]))
		if err != nil {
			return false
		} else {
			// Return user
			return true
		}
	}
}

func getUserFromAuth(w http.ResponseWriter, r *http.Request) string {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return ""
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return ""
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return ""
	}
	return pair[0]
}
