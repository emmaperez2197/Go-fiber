package jwt

import (
	"time"

	"github.com/emmaperez2197/go-fiber/config"
	"github.com/emmaperez2197/go-fiber/models"
	"github.com/golang-jwt/jwt/v5"
)

func GeneratetToken(user models.User) (string, error) {

	secret := config.Config("SECRET_KEY")
	claims := jwt.MapClaims{
		"name":     user.Email,
		"lastname": user.LastName,
		"email":    user.Email,
		"active":   user.Active,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))

	return t, err

}
