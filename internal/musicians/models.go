package musicians

import "time"

// MusicianProfile represents a musician's profile in the database
type MusicianProfile struct {
	ID              int       `json:"id"`
	UserID          int       `json:"user_id"`
	Instruments     []string  `json:"instruments"`
	Genres          []string  `json:"genres,omitempty"`
	ExperienceLevel string    `json:"experience_level,omitempty"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Bio             string    `json:"bio,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
}
