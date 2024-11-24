package utilities

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(userName string, admin bool) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": userName,
			"exp":      time.Now().Add(time.Hour * 24 * 8).Unix(),
			"admin":    admin,
		},
	)

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(hashPassword), nil
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))

	if err != nil {
		return false
	}

	return true
}

func VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return false, errors.New("error occured while parsing token")
	}

	if !token.Valid {
		return false, errors.New("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return false, errors.New("session expired")
	}

	_ = claims

	return true, nil
}

func GetUserSessionDetails(c *gin.Context) error {
	encodedJwt, _ := c.Cookie("auth_token")

	if encodedJwt == "" {
		c.Header("HX-Redirect", "/")
		c.Status(200)
		return nil
	}

	claims, err := VerifyToken(encodedJwt)

	if err != nil {
		return err
	}

	if !claims {
		return errors.New("session expired")
	}

	return nil
}

func username(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return "", errors.New("error occured while parsing token")
	}

	if !token.Valid {
		return "", errors.New("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("session expired")
	}

	username := claims["username"].(string)

	return username, nil

}

func GetUserName(c *gin.Context) (string, error) {
	encodedJwt, _ := c.Cookie("auth_token")

	if encodedJwt == "" {
		c.Header("HX-Redirect", "/")
		c.Status(200)
		return "", nil
	}

	username, err := username(encodedJwt)

	if err != nil {
		return "", err
	}

	return username, nil

}
