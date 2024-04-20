package service

import ("log"
	"dependency_injection/domain"
)

type AwsPublisher struct {
}

func NewAwsPublisher() * AwsPublisher {
	return &AwsPublisher{}
}

func (p *AwsPublisher) Publish(student *domain.Student) {
	log.Println("Publishing student to AWS: ", student)
}