// Package handlers
package handlers

import (
	"net/http"
)

type MyUser struct{}

func (user *MyUser) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
	} else {
		panic("The method no is valid")
	}
}

func (user *MyUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
	} else {
		panic("The method no is valid")
	}
}
