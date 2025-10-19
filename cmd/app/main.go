package main

import (
	"log"
	"net/http"

	"github.com/jeisvalen745/Bcami/internal/handlers"
)

func main() {
	handlersUsers := &handlers.MyUser{}
	handlersRegister := http.NewServeMux()
	handlersRegister.HandleFunc("/user/login", handlersUsers.Login)
	handlersRegister.HandleFunc("/user/update", handlersUsers.UpdateUser)

	log.Fatal(http.ListenAndServe(":8080", handlersRegister))
}
