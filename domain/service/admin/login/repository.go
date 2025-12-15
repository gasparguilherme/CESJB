package login

import "cesjb/domain/entities"

type Repository interface {
	FindAdminByEmail(email string) (*entities.Admin, error)
}
