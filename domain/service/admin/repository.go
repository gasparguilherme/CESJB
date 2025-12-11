package admin

import "cesjb/domain/entities"

type Repository interface {
	SaveAdmin(data entities.Admin) (int, error)
}
