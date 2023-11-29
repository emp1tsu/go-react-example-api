package main

import (
	"go-react-example-api/controller"
	"go-react-example-api/db"
	"go-react-example-api/repository"
	"go-react-example-api/router"
	"go-react-example-api/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	userUseCase := usecase.NewUserUsecase(userRepository)
	taskUseCase := usecase.NewTaskUsecase(taskRepository)
	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)
	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
