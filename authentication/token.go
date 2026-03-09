package authentication

import (
	"cesjb/config"
	"fmt"
	"net/http"
	"strings"
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

// validar token verifica se o token passsado na requisicao é valido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	if tokenString == "" {
		return fmt.Errorf("token não fornecido")
	}

	token, err := jwt.Parse(tokenString, returnKeyVerify)
	if err != nil {
		return fmt.Errorf("erro ao fazer parse do token: %w", err)
	}

	if !token.Valid {
		return fmt.Errorf("token inválido")
	}
	return nil
}

// pega o token do header
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func returnKeyVerify(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf(
			"método de assinatura inesperado: %v",
			token.Header["alg"],
		)
	}

	return []byte(config.SecretKey), nil
}
