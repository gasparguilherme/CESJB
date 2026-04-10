package payment

import (
	"cesjb/types_"
	"context"
	"fmt"
	"time"
)

func (r Repository) GetMonthlyTotal(competence types_.DateOnly) (float64, error) {
	query := `
		SELECT COALESCE(SUM(value), 0)
		FROM payments
		WHERE DATE_TRUNC('month', competence) = DATE_TRUNC('month', $1)
		AND status = true
	`

	var total float64
	err := r.connectionInstance.QueryRow(context.TODO(), query, time.Time(competence)).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("erro ao buscar total mensal: %w", err)
	}

	return total, nil
}
