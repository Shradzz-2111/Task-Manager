package helpers

import (
	// "log"
	"time"
	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte("VISHWAAI")

type CustomClaims struct {
    UserID uint `json:"user_id"` // Your custom claim
    jwt.StandardClaims
}

func GenerateTokens(userID uint)(string, error){
	claims := jwt.MapClaims{
		"user_id":userID,
		"exp":time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}


func ValidateToken(tokenString string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(
        tokenString,
        &CustomClaims{},
        func(token *jwt.Token) (interface{}, error) {
            return jwtSecret, nil
        },
    )

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, jwt.ErrInvalidKey
}