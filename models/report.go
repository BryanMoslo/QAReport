package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Report struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	Filename      string    `json:"filename" db:"filename"`
	TotalTime     int       `json:"total_time" db:"total_time"`
	TotalDistance float64   `json:"total_distance" db:"total_distance"`
	From          string    `json:"from" db:"from"`
	To            string    `json:"to" db:"to"`
}

// Reports is not required by pop and may be deleted
type Reports []Report

// String is not required by pop and may be deleted
func (r Reports) String() string {
	jr, _ := json.Marshal(r)
	return string(jr)
}

// Validate gets run everytime you call a "pop.Validate" method.
// This method is not required and may be deleted.
func (r *Report) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: r.Filename, Name: "Filename"},
		&validators.StringIsPresent{Field: r.From, Name: "From"},
		&validators.StringIsPresent{Field: r.To, Name: "To"},
	), nil
}
