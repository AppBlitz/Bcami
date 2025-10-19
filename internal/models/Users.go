package models

type Rol string

type User struct {
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Rol      Rol    `json:"rol"`
}
