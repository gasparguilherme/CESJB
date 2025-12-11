package admin

import (
	"cesjb/domain/entities"
)

func (s Service) CreateAdmin(name, email, password string) (*entities.Admin, error) {
	newAdmin := &entities.Admin{
		Name:     name,
		Email:    email,
		Password: password,
	}
	id, err := s.repository.SaveAdmin(*newAdmin)
	if err != nil {
		return nil, err
	}
	newAdmin.ID = id
	return newAdmin, nil
}
