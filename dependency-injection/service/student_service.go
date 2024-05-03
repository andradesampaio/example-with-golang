package service

import (
	"dependency-injection/repository"

	"log"
)

// Service is a simple service
type StudentService struct {
	repository *repository.StudentRepository
	publisher Publisher
}

// NewService creates a new main service
func NewMyService(publisher Publisher, repository *repository.StudentRepository) *StudentService {
	return &StudentService{publisher: publisher, repository: repository}
}

func (service *StudentService) Run() {
	student := service.repository.Find()
    service.publisher.Publish(student)
	log.Println("Running the my service")
}