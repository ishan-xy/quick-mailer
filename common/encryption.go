package common

import (
	"time"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	utils "github.com/ItsMeSamey/go_utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, utils.WithStack(err)
}

func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}

func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(Cfg.JWTExpiration).Unix()

	tokenString, err := token.SignedString([]byte(Cfg.JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(Cfg.JWTSecret), nil
	})
}

func GenerateAPIKey(secret string) (string, string, error) {
	serialNumber := uuid.New().String()

	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate random bytes: %v", utils.WithStack(err))
	}
	randomBits := hex.EncodeToString(randomBytes)

	data := fmt.Sprintf("%s:%s", serialNumber, randomBits)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))

	apiKey := fmt.Sprintf("%s:%s:%s", serialNumber, randomBits, signature)
	encodedKey := base64.URLEncoding.EncodeToString([]byte(apiKey))

	return encodedKey, serialNumber,nil
}