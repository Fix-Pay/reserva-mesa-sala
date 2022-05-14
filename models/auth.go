package models

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UserClaims struct {
	User User `json:"user"`
	jwt.StandardClaims
}
