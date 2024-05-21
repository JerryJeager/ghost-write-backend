package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/JerryJeager/ghost-write-backend/models/users"
	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateToken(user users.User) (string, error) {

    tokenLifespan, err := strconv.Atoi("24")

    if err != nil {
        return "", err
    }

    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["id"] = user.ID
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    return token.SignedString([]byte(os.Getenv("API_SECRET")))

}
