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
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Filename  string    `json:"filename" db:"filename"`
	Distance  float64   `json:"distance" db:"distance"`
	ImageSrc  string    `json:"image_src" db:"image_src"`
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
		&validators.StringIsPresent{Field: r.ImageSrc, Name: "ImageSrc"},
	), nil
}
