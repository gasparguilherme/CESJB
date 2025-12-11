package types_

import (
	"fmt"
	"strings"
	"time"
)

// DateOnly representa apenas a data (sem horário)
type DateOnly time.Time

// MarshalJSON serializa DateOnly no formato yyyy-mm-dd
func (d DateOnly) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	if t.IsZero() {
		return []byte(`null`), nil
	}
	formatted := fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// UnmarshalJSON faz parse do JSON -> DateOnly
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		*d = DateOnly(time.Time{}) // data zero
		return nil
	}

	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = DateOnly(t)
	return nil
}

// ToTime converte DateOnly -> time.Time (útil para INSERT/UPDATE)
func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}

// IsZero verifica se a data é zero
func (d DateOnly) IsZero() bool {
	return time.Time(d).IsZero()
}

// Scan implementa sql.Scanner para ler do PostgreSQL
func (d *DateOnly) Scan(value any) error {
	if value == nil {
		*d = DateOnly(time.Time{})
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("DateOnly: não é um time.Time: %v", value)
	}
	*d = DateOnly(t)
	return nil
}

// Value implementa driver.Valuer para enviar DateOnly para o PostgreSQL
func (d DateOnly) Value() (any, error) {
	return time.Time(d), nil
}
