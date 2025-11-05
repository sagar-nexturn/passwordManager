package repository

type InitDbRepo interface {
	CreatePasswordsTableIfNotExist() error
	InsertSampleData() error
}
