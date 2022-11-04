package models

import "github.com/golang-jwt/jwt"

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Login struct {
	Username string `json:"loginUsername"`
	Password string `json:"loginPass"`
}

type Claims struct {
	Username string `json:"loginUsername"`
	jwt.StandardClaims
}
