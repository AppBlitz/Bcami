package main

import (
	"log"
	"net/http"

	"github.com/jeisvalen745/Bcami/internal/handlers"
)

func main() {
	handlersUsers := &handlers.MyUser{}
	handlersRegister := http.NewServeMux()
	handlersRegister.HandleFunc("/login/user", handlersUsers.Login)

	log.Fatal(http.ListenAndServe(":8080", handlersRegister))
}
