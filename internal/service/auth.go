package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/leeerraaa/backend-app/internal/domain"
	repo "github.com/leeerraaa/backend-app/internal/repository/psql"
)

const (
	salt       = "superhashkey"
	signingKey = "L6e2h3e6gfE4ae93AZMfPLRg782Y"
	tokenTTL   = 18 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}

type Auth struct {
	repo repo.User
}

func NewAuthService(repo repo.User) *Auth {
	return &Auth{
		repo: repo,
	}
}

func (a *Auth) UserInfo(userId string) (domain.User, error) {
	return a.repo.UserInfo(userId)
}

func (a *Auth) GenerateToken(login, password string) (string, error) {
	userId, err := a.repo.GetUser(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(signingKey))
}

func (a *Auth) ParseToken(accessToken string) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
