package admin

import (
	"cesjb/domain/entities"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (s Service) CreateAdmin(name, email, password string) (*entities.Admin, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("erro ao gerar hash da senha: %w", err)
	}

	newAdmin := &entities.Admin{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	id, err := s.repository.SaveAdmin(*newAdmin)
	if err != nil {
		return nil, fmt.Errorf("erro ao salvar admin no banco: %w", err)
	}

	newAdmin.ID = id

	return newAdmin, nil
}
