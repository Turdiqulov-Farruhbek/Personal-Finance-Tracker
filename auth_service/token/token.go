package token

import (
	"errors"
	"fmt"
	"log"

	// "log"
	"time"

	"github.com/golang-jwt/jwt"
	// pb "finance_tracker/auth_service/genproto"
)

const (
	signingKey = "secret_key"
)

func GenerateJWTToken(userID, username, role interface{}) (string, string) {
	accessToken := jwt.New(jwt.SigningMethodHS256)
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	claims := accessToken.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["username"] = username
	claims["role"] = role
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(60 * time.Minute).Unix()
	access, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating access token: ", err)
	}

	rftClaims := refreshToken.Claims.(jwt.MapClaims)
	rftClaims["id"] = userID
	rftClaims["iat"] = time.Now().Unix()
	rftClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	refresh, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		log.Fatal("error while generating refresh token: ", err)
	}
	fmt.Println("Refresh token: ", refresh)

	return access, refresh
}

func ValidateToken(tokenStr string) (bool, error) {
	log.Println("token from heaf=der ", tokenStr)
	_, err := ExtractClaim(tokenStr)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	fmt.Print(token.Claims)
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
func JustExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parsing token: %w", err)
	}
	// fmt.Print(token.Claims)

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
