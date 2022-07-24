package api

import (
	"net/http"

	"github.com/AntonyIS/todo-app-one/app"
	"github.com/gin-gonic/gin"
)

type TodoHandler interface {
	Create(*gin.Context)
	Read(*gin.Context)
	Index(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

type handler struct {
	todoService app.TodoService
}

func NewHandler(todoService app.TodoService) TodoHandler {
	return &handler{
		todoService: todoService,
	}
}

func (h *handler) Create(c *gin.Context) {
	var todo app.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}

	newTodo, err := h.todoService.Create(&todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": newTodo,
	})
}

func (h *handler) Read(c *gin.Context) {
	todoID := c.Param("id")
	todo, err := h.todoService.Read(todoID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"err": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": todo,
	})
}

func (h *handler) Index(c *gin.Context) {

	todos, err := h.todoService.ReadAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": todos,
	})
}

func (h *handler) Update(c *gin.Context) {
	var todo app.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrInvalidTodo,
		})

		return
	}
	updatedTodo, err := h.todoService.Update(&todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrInvalidTodo,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"todo": updatedTodo,
	})

}

func (h *handler) Delete(c *gin.Context) {
	todoID := c.Param("id")
	if err := h.todoService.Delete(todoID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": app.ErrInvalidTodo,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo deleted successfuly",
	})

}
