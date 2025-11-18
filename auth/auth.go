package auth

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("test123")

func GenerateJWT(UserID uint) (string,error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(secretKey)
}

func ValidadeToken(tokenString string) (*jwt.Token, error){
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		return secretKey, nil
	})
}
