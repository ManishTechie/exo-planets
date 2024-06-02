package request

import (
	"errors"
	"exo-planets/util"
	"time"

	"github.com/google/uuid"
	"go.uber.org/multierr"
	"gorm.io/gorm"
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

// Optional: String method to print the enum value
func (e ExoplanetsType) String() string {
	return string(e)
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
}

func (e *Exoplanet) BeforeCreate(tx *gorm.DB) (err error) {
	e.ID = uuid.New()
	return
}

func (e Exoplanet) Validate() (err error) {
	if e.Name == nil || util.StringPointerToString(e.Name) == "" {
		err = multierr.Append(err, errors.New("'Name' should be present in a request"))
	}

	if e.Distance == nil || *e.Distance < 0 {
		err = multierr.Append(err, errors.New("'Distance' should be present or greater than 0"))
	}

	if e.Radius == nil || *e.Radius < 0 {
		err = multierr.Append(err, errors.New("'Radius' should be present or greater than 0"))
	}

	if e.Mass == nil || *e.Mass < 0 {
		err = multierr.Append(err, errors.New("'Mass' should be present or greater than 0"))
	}

	if e.Type == nil || !e.Type.IsValid() {
		err = multierr.Append(err, errors.New("'Type' should be present in a request {GasGiant,Terrestrial}"))
	}

	return err
}
