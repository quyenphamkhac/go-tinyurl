package entities

import "time"

type User struct {
	ID             string    `json:"id,omitempty"`
	Username       string    `json:"username,omitempty"`
	Name           string    `json:"name,omitempty"`
	HashedPassword string    `json:"-"`
	Email          string    `json:"email,omitempty"`
	CreationDate   time.Time `json:"creation_date,omitempty"`
	LastLogin      time.Time `json:"last_login,omitempty"`
}
