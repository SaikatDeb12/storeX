package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func ParseJSON(body io.Reader, out interface{}) error {
	return json.NewDecoder(body).Decode(out)
}

func EncodeJSON(w http.ResponseWriter, out interface{}) error {
	return json.NewEncoder(w).Encode(out)
}

const SecretKey = "SECRET"

type ErrorModel struct {
	Error      string
	Message    string
	StatusCode int
}

func RespondError(w http.ResponseWriter, statusCode int, err error, message string) {
	w.WriteHeader(statusCode)
	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	NewError := ErrorModel{
		Message:    message,
		Error:      errStr,
		StatusCode: statusCode,
	}

	if err := EncodeJSON(w, NewError); err != nil {
		fmt.Printf("error: %+v", err)
	}
}

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		err := EncodeJSON(w, body)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
	}
}

func HashedPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateJWT(userID, sessionID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userID,
		"session_id": sessionID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(SecretKey))
}
