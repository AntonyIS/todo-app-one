package main

import (
	"log"
	"os"

	h "github.com/AntonyIS/todo-app-one/api/http"
	svc "github.com/AntonyIS/todo-app-one/app/core"
	"github.com/gorilla/mux"

	"github.com/AntonyIS/todo-app-one/app"
	pq "github.com/AntonyIS/todo-app-one/repository/postgresql"
	rd "github.com/AntonyIS/todo-app-one/repository/redis"
	"github.com/gin-gonic/gin"
)

func main() {

	todoRepo := todoRepo()
	todoService := svc.NewTodoService(todoRepo)
	todoHandler := h.NewHandler(todoService)

	userRepo := userRepo()
	userService := svc.NewUserService(userRepo)
	userHandler := h.NewUserHandler(userService)

	ginRouter := gin.Default()
	muxRouter := mux.NewRouter()
	ginRouter.POST("/", todoHandler.Create)
	ginRouter.GET("/:id", todoHandler.Read)
	ginRouter.GET("/", todoHandler.Index)
	ginRouter.PUT("/", todoHandler.Update)
	ginRouter.DELETE("/:id", todoHandler.Delete)

	muxRouter.HandleFunc("/", userHandler.CreateUser)
	muxRouter.HandleFunc("/:id", userHandler.ReadUser)
	muxRouter.HandleFunc("/", userHandler.ReadAllUsers)
	muxRouter.HandleFunc("/", userHandler.UpdateUser)
	muxRouter.HandleFunc("/:id", userHandler.DeleteUser)

	ginRouter.Run(port())

}

func port() string {
	port := ":8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func todoRepo() app.TodoRepository {

	repo, err := rd.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected: ", err)
		return nil
	}
	return repo
}

func userRepo() app.UserRepository {

	repo, err := pq.NewPostgresRepository()

	if err != nil {
		log.Fatal("redis server not connected: ", err)
		return nil
	}
	return repo
}
