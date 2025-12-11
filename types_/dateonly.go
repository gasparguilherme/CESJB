package types_

import (
	"fmt"
	"strings"
	"time"
)

type DateOnly time.Time

// Serializa no formato yyyy-mm-dd
func (d DateOnly) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	formatted := fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Faz parse do JSON -> DateOnly
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" {
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

// Converte DateOnly -> time.Time
func (d DateOnly) ToTime() time.Time {
	return time.Time(d)
}

// Verifica se a data é zero
func (d DateOnly) IsZero() bool {
	return time.Time(d).IsZero()
}
