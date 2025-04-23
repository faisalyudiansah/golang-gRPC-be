package repository

import "database/sql"

type ExampleRepositoryInterface interface {
	Get() error
}

type ExampleRepositoryImplementation struct {
	db *sql.DB
}

func NewExampleRepositoryImplementation(db *sql.DB) *ExampleRepositoryImplementation {
	return &ExampleRepositoryImplementation{
		db: db,
	}
}

func (rs *ExampleRepositoryImplementation) Get() error {
	return nil
}
