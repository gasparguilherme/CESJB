package associate

import "cesjb/domain/entities"

func (s Service) GetByID(id int) (*entities.Associate, error) {
	associate, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return associate, nil
}
