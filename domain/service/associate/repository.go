package associate

import "cesjb/domain/entities"

type Repository interface {
	SaveAssociate(data entities.Associate) (int, error)
	GetAssociates() ([]entities.Associate, error)
}
