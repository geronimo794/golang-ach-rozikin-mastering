package claim

import (
	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	Username string `json:"name"`
	jwt.StandardClaims
}
