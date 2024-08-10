package model

import (
	"time"

	"github.com/google/uuid"
)

type ExoplanetsType string

// Types of exoplanets
const (
	TerrestrialType ExoplanetsType = "Terrestrial"
	GasGiantType    ExoplanetsType = "GasGiant"
)

// Optional: Method to check if a value is valid
func (e ExoplanetsType) IsValid() bool {
	switch e {
	case TerrestrialType, GasGiantType:
		return true
	}
	return false
}

// Exoplanet -
type Exoplanet struct {
	ID          uuid.UUID
	Name        *string
	Description *string
	Distance    *float64
	Radius      *float64
	Mass        *float64
	Type        *ExoplanetsType
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
