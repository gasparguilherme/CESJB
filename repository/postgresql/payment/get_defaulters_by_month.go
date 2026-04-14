package payment

import (
	"cesjb/domain/entities"
	"cesjb/types_"
	"context"
	"fmt"
	"time"
)

// GetDefaultersByMonth retorna os associados ativos que ainda não pagaram no mês informado
func (r Repository) GetDefaultersByMonth(competence types_.DateOnly) ([]entities.Associate, error) {
	query := `
		SELECT
			a.id,
			a.name,
			a.cpf
		FROM associates a
		LEFT JOIN payments p
			ON p.associate_id = a.id
			AND DATE_TRUNC('month', p.competence::date) = DATE_TRUNC('month', $1::date)
		WHERE a.status = true
		AND p.id IS NULL
		ORDER BY a.name ASC
	`

	rows, err := r.connectionInstance.Query(context.TODO(), query, time.Time(competence))
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar inadimplentes: %w", err)
	}
	defer rows.Close()

	// inicializa slice vazio para evitar retorno null quando não há inadimplentes
	defaulters := []entities.Associate{}

	for rows.Next() {
		var a entities.Associate
		err := rows.Scan(&a.ID, &a.Name, &a.CPF)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler inadimplente: %w", err)
		}
		defaulters = append(defaulters, a)
	}

	return defaulters, nil
}
