package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/types"
	"time"
)

const (
	salt       = "dfghuiehrgyeg674hgijdnjkashwegf7"
	tokenTTL   = 10 * time.Minute
	signingKey = "dijuaehrguerygnfkjvnaskfjqhwyfgr654fsdfa"
	Validate   = "user not registered in database"
)

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return a.db.CreateUser(user)
}
func (a *AuthService) GenerateToken(username, password string) (string, error) {
	//get User from DB
	user, err := a.db.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}
func (a *AuthService) Validate(username, password string) (string, error) {
	_, err := a.db.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return Validate, nil
	}
	return "You are already registered", err
}