package authentication

import (
	"cesjb/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5" // ← nova lib
)

// CreateToken retorna um token assinado com as permissões do usuário
func CreateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"exp":        jwt.NewNumericDate(time.Now().Add(time.Hour * 6)), // ← tipo correto para exp
		"userID":     userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.SecretKey) // ← SecretKey já é []byte, não precisa converter
}

// ValidateToken verifica se o token passado na requisição é válido
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	if tokenString == "" {
		return fmt.Errorf("token não fornecido")
	}

	token, err := jwt.Parse(tokenString, returnKeyVerify,
		jwt.WithValidMethods([]string{"HS256"}), // ← valida o algoritmo de forma segura
	)
	if err != nil {
		return fmt.Errorf("erro ao fazer parse do token: %w", err)
	}

	if !token.Valid {
		return fmt.Errorf("token inválido")
	}
	return nil
}

// extractToken extrai o token do header Authorization
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
	// A validação do algoritmo agora é feita via WithValidMethods acima,
	// mas mantemos a verificação aqui como segunda camada de segurança
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
