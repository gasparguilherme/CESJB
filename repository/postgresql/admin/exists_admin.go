package admin

import "context"

func (r Repository) ExistsAdmin() (bool, error) {
	var exists bool

	err := r.connectionInstance.
		QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM admins)").
		Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
