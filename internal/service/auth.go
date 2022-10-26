package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"github.com/ssharifzoda/levelup/internal/database"
	"github.com/ssharifzoda/levelup/internal/types"
	"time"
)

const (
	salt                = "dfghuiehrgyeg674hgijdnjkashwegf7"
	tokenTTL            = 60 * time.Minute
	signingKey          = "dijuaehrguerygnfkjvnaskfjqhwyfgr654fsdfa"
	refreshKey          = "cfre65645fdg1g5rg52a"
	negativeValidUser   = "user already registered"
	newUser             = "user not registered in database"
	negativePasswordVal = "creating again correct password"
	noCyrillyc          = "dont typing Cyrillyc"
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

type RefreshtokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (a *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return a.db.CreateUser(user)
}
func (a *AuthService) GenerateTokens(username, password string) (string, string, error) {
	//get User from DB
	if len(password) > 25 {
		user, err := a.db.GetUser(username, password)
		if err != nil {
			return "", "", err
		}
		token, refreshToken, err := tokenGenerator(user.Id)
		if err != nil {
			return "", "", err
		}
		return token, refreshToken, nil
	}
	user, err := a.db.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", "", err
	}
	token, refreshToken, err := tokenGenerator(user.Id)
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
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
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type")
	}
	return claims.UserId, nil
}

func (a *AuthService) UserValidate(username, password string) (string, error) {
	_, err := a.db.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return newUser, nil
	}
	return negativeValidUser, err
}

func (a *AuthService) ParseRefreshToken(refreshToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshtokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(refreshKey), nil
	})
	if err != nil {
		return "", "", nil
	}
	claims, ok := token.Claims.(*RefreshtokenClaims)
	if !ok {
		return "", "", errors.New("token claims are not of type")
	}
	username, passwordHash, err := a.db.GetUserById(claims.UserId)
	if err != nil {
		return "", "", err
	}
	return username, passwordHash, nil
}
func tokenGenerator(id int) (string, string, error) {
	first := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("auth.tokenttl")).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, id,
	})
	second := jwt.NewWithClaims(jwt.SigningMethodHS256, &RefreshtokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(viper.GetDuration("auth.refreshtokenttl")).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, id,
	})
	token, err := first.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}
	refreshToken, err := second.SignedString([]byte(refreshKey))
	if err != nil {
		return "", "", err
	}
	return token, refreshToken, nil
}
