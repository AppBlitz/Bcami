// Package handlers
package handlers

import (
	"net/http"

	"github.com/jeisvalen745/Bcami/internal/models"
)

type MyUser struct {
	*models.User
}

func (userLogin *MyUser) Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Solution sucessfully"))
}
