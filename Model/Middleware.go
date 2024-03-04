package Model

import "github.com/golang-jwt/jwt/v5"

type AuthClaimJWT struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}
