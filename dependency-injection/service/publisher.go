package service

import "dependency-injection/domain"

type Publisher interface {
	Publish(studet *domain.Student)
}