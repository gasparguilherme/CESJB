package associate

import "cesjb/domain/entities"

func (s Service) ListAssociates() ([]entities.Associate, error) {
	return s.repository.GetAssociates()
}
