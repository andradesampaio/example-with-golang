package repository

import (
	"dependency_injection/domain"
)


// StudentRepository is a simple repository
type StudentRepository struct {
	db *PostgreSQLDatabase
}

// NewStudentRepository creates a new student repository
func NewStudentRepository(db *PostgreSQLDatabase) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) Find() *domain.Student {
	return &domain.Student{ID: 1, Name: "John Doe"}
}