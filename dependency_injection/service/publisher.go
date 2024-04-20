package service

import "dependency_injection/domain"

type Publisher interface {
	Publish(studet *domain.Student)
}