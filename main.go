package main

import (
	"log"
	"os"

	h "github.com/AntonyIS/todo-app-one/api"
	"github.com/AntonyIS/todo-app-one/app"
	rd "github.com/AntonyIS/todo-app-one/repository/redis"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := repo()
	service := app.NewTodoService(repo)
	handler := h.NewHandler(service)

	r := gin.Default()

	r.POST("/", handler.Create)
	r.GET("/:id", handler.Read)
	r.GET("/", handler.Index)
	r.PUT("/", handler.Update)
	r.DELETE("/:id", handler.Delete)

	r.Run(port())

}

func port() string {
	port := ":8000"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return port
}

func repo() app.TodoRepository {
	repo, err := rd.NewRedisRepository("redis://localhost:6379")

	if err != nil {
		log.Fatal("redis server not connected")
		return nil
	}
	return repo
}
