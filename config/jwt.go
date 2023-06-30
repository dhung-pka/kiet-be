package config

import "github.com/go-chi/jwtauth/v5"

func createTokenAuth() {
	tokenAuth = jwtauth.New("HS256", []byte("ke_that_hua"), nil)
}

func GetJWT() *jwtauth.JWTAuth {
	return tokenAuth
}
