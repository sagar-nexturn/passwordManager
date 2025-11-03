package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/sagar-nexturn/passwordManager/internal/models"
)

type PostgresPasswordRepo struct {
	db *sql.DB
}

func (m *PostgresPasswordRepo) GetPasswordByID(id string) (*models.Password, error) {
	query := `
        SELECT id, name, username, secret, nonce, created_at, updated_at
        FROM passwords
        WHERE id = $1
    `
	row := m.db.QueryRow(query, id)

	var p models.Password

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Username,
		&p.Secret,
		&p.Nonce,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("password with id %s not found", id)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &p, nil
}

func (m *PostgresPasswordRepo) GetAllPasswords() ([]models.Password, error) {
	rows, err := m.db.Query(`SELECT id, name, username, secret, nonce, created_at, updated_at FROM passwords`)
	if err != nil {
		return nil, fmt.Errorf("failed to query passwords: %v", err)
	}
	defer rows.Close()

	var passwords []models.Password
	for rows.Next() {
		var p models.Password
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Username,
			&p.Secret,
			&p.Nonce,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		passwords = append(passwords, p)
	}
	return passwords, nil
}

func (m *PostgresPasswordRepo) GetPasswordByName(name string) (*models.Password, error) {
	query := `
        SELECT id, name, username, secret, nonce, created_at, updated_at
        FROM passwords
        WHERE name = $1
    `
	row := m.db.QueryRow(query, name)

	var p models.Password

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Username,
		&p.Secret,
		&p.Nonce,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("password with name %s not found", name)
		}
		return nil, fmt.Errorf("query error: %v", err)
	}

	return &p, nil
}

func (m *PostgresPasswordRepo) UpdatePassword(password *models.Password) error {
	password.UpdatedAt = time.Now()

	query := `
	UPDATE passwords 
	SET name=$1, username=$2, secret=$3, nonce=$4, updated_at=$5 
	WHERE id=$6
	`

	res, err := m.db.Exec(query, password.Name, password.Username, password.Secret, password.Nonce, password.UpdatedAt, password.ID)
	if err != nil {
		return fmt.Errorf("failed to update password: %v", err)
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("no record found to update")
	}

	return nil
}

func (m *PostgresPasswordRepo) DeletePassword(id string) error {
	query := `DELETE FROM passwords WHERE id = $1`
	res, err := m.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete password: %v", err)
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		return fmt.Errorf("no record found to delete")
	}

	return nil
}

func (m *PostgresPasswordRepo) AddPassword(password *models.Password) error {
	query := `
		INSERT INTO passwords (id, name, username, secret, nonce, createdAt, updatedAt)
		VALUES ($1,$2,$3,$4,$5,$6,$7)
		`
	_, err := m.db.Exec(query, password.ID, password.Name, password.Username, password.Secret, password.Nonce, password.CreatedAt, password.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert password info: %w", err)
	}

	return nil
}

func NewPostgresPasswordRepo(db *sql.DB) PasswordDbRepo {
	return &PostgresPasswordRepo{db: db}
}
