package models

type Rol string

const (
	Admin Rol = "Admininistration"
	Users Rol = "User"
)

type User struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Rol      Rol    `json:"rol"`
}
