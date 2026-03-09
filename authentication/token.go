package authentication

import (
	"cesjb/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken retorna um token assiando com as permissoes do usuario
func CreateToken(userID uint64) (string, error) {
	permition := jwt.MapClaims{}
	permition["authorizad"] = true
	permition["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permition["userID"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permition)
	return token.SignedString([]byte(config.SecretKey)) //secret

}
