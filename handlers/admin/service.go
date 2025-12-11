package admin

import "cesjb/domain/entities"

type Service interface {
	CreateAdmin(name, email, password string) (*entities.Admin, error)
}
