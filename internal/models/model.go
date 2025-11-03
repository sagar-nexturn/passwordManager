package models

import "time"

type Password struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Secret    []byte    `json:"secret"`
	Nonce     []byte    `json:"nonce"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
