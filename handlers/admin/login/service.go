package login

import (
	"cesjb/domain/entities"
	"cesjb/domain/service/admin/login"
)

type Service interface {
	FindAdminByEmail(login.Login) (*entities.Admin, error)
}
