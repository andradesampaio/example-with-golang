package main

import (
	"dependency-injection/repository"
	"dependency-injection/service"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(service.NewMyService),
		fx.Provide(repository.NewStudentRepository),
		fx.Provide(repository.NewPostgreSQLDatabase),
		fx.Provide(
			fx.Annotate(
				service.NewAwsPublisher,
				fx.As(new(service.Publisher)),
			),
		),
		fx.Invoke(func(service *service.StudentService) {
			service.Run()
		}),
	).Run()
}
