package cnf

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("bfiu32hguri9g3h12951y9083jnmED32n")

type JWTClaim struct {
	UserName string
	jwt.RegisteredClaims
}