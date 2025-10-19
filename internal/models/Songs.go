package models

import (
	"time"
)

type Song struct {
	ID                    string    `json:"id"`
	Title                 string    `json:"title"`
	Artist                Artist    `json:"Artist"`
	DateSongCreationAge   time.Time `json:"DateSongCreationAge"`
	DurationSongInMinutes uint8     `json:"DurationSongInMinutes"`
}
