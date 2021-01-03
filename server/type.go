package server

import "time"

type fileModel struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Encoded string `json:"encoded"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
