package admin

import (
	"cesjb/domain/entities"
	"cesjb/domain/service/admin"
)

type Service interface {
	CreateAdmin(name, email, password string) (*entities.Admin, error)
	Login(data admin.Login) (*entities.Admin, error)
}
