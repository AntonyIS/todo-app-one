package core

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
	"github.com/AntonyIS/todo-app-one/app"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidTodo  = errors.New("invalid todo")
	ErrNoTodos      = errors.New("todos not available todo")
)

type todoService struct {
	todoRepo app.TodoRepository
}



func NewTodoService(todoRepo app.TodoRepository) app.TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) Create(todo *app.Todo) (*app.Todo, error) {
	todo.CreateAt = time.Now().UTC().Unix()
	todo.ID = shortid.MustGenerate()
	return t.todoRepo.Create(todo)
}

func (t *todoService) Read(id string) (*app.Todo, error) {
	return t.todoRepo.Read(id)
}
func (t *todoService) ReadAll() (*[]app.Todo, error) {
	return t.todoRepo.ReadAll()
}

func (t *todoService) Update(todo *app.Todo) (*app.Todo, error) {
	return t.todoRepo.Update(todo)
}

func (t *todoService) Delete(id string) error {
	return t.todoRepo.Delete(id)
}
