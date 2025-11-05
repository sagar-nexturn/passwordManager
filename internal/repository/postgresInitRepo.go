package repository

import (
	"database/sql"
	"encoding/hex"

	"github.com/google/uuid"
)

type PostgresInitRepo struct {
	db *sql.DB
}

func (m *PostgresInitRepo) CreatePasswordsTableIfNotExist() error {
	query := `
		CREATE TABLE IF NOT EXISTS passwords (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		username TEXT NOT NULL,
		secret BYTEA NOT NULL,
		nonce BYTEA NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMP NOT NULL DEFAULT NOW()
	);`

	_, err := m.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (m *PostgresInitRepo) InsertSampleData() error {
	query := `
		INSERT INTO passwords (id, name, username, secret, nonce)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO NOTHING
	`

	ID := uuid.New().String()
	secretHex := "9a93b1b9f80e66a8cf96c3efb55d2a31c5f6891c50"
	nonceHex := "2d67ae2cf7db1c6712d6b95d"
	secretBytes, _ := hex.DecodeString(secretHex)
	nonceBytes, _ := hex.DecodeString(nonceHex)

	_, err := m.db.Exec(query, ID, "Git", "Sagar", secretBytes, nonceBytes)
	if err != nil {
		return err
	}

	return nil
}

func NewPostgresInitRepo(db *sql.DB) InitDbRepo {
	return &PostgresInitRepo{
		db: db,
	}
}
