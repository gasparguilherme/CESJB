package associated

import "cesjb/domain/entities"

type Repository interface {
	SaveAssociated(data entities.Associated) (int, error)
}
