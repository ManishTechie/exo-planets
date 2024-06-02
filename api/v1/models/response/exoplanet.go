package response

import (
	"time"

	"github.com/google/uuid"
)

// Exoplanet -
type Exoplanet struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Distance    float64    `json:"distance"`
	Radius      float64    `json:"radius"`
	Mass        float64    `json:"mass"`
	Type        string     `json:"type"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
