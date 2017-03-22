package main

import "net/http"

func main() {
	http.HandleFunc("/admins/users", users)
	http.HandleFunc("/admins/policies", adminPolicies)
	http.HandleFunc("/users/policies", userPolicies)

	http.ListenAndServe(":3000", nil)
}
