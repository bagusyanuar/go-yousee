package common

import (
	"github.com/bagusyanuar/go-yousee/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

type JWTSignReturn struct {
	UserID uuid.UUID `json:"user_id"`
}

func CreateAccessToken(cfg *config.JWT, jwtSign *JWTSignReturn) (accessToken string, err error) {
	JWTSigninMethod := jwt.SigningMethodHS256
	claims := JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: cfg.Issuer,
		},
		UserID: jwtSign.UserID,
	}
	token := jwt.NewWithClaims(JWTSigninMethod, claims)
	return token.SignedString([]byte(cfg.SignatureKey))
}
