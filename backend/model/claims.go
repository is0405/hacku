package model

import (
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID           int    `json:"id"` 
	Role         string `json:"role"` // Audience (Ex. user, admin)
	jwt.StandardClaims        // expires, issued at, not before
}
