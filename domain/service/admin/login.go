package admin

import (
	"cesjb/domain/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) FindAdminByEmail(data Login) (*entities.Admin, error) {
	admin, err := s.repository.FindAdminByEmail(data.Email)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado")

	}
	if bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password)) != nil {
		return nil, fmt.Errorf("senha incorreta")
	}

	return admin, nil
}
