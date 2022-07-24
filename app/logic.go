package app

import (
	"errors"
	"time"

	"github.com/teris-io/shortid"
)

var (
	ErrTodoNotFound = errors.New("todo not found")
	ErrInvalidTodo  = errors.New("invalid todo")
	ErrNoTodos      = errors.New("todos not available todo")
)

type todoService struct {
	todoRepo TodoRepository
}

func NewTodoService(todoRepo TodoRepository) TodoService {
	return &todoService{
		todoRepo: todoRepo,
	}
}

func (t *todoService) Create(todo *Todo) (*Todo, error) {
	todo.CreateAt = time.Now().UTC().Unix()
	todo.ID = shortid.MustGenerate()
	return t.todoRepo.Create(todo)
}

func (t *todoService) Read(id string) (*Todo, error) {
	return t.todoRepo.Read(id)
}
func (t *todoService) ReadAll() (*[]Todo, error) {
	return t.todoRepo.ReadAll()
}

func (t *todoService) Update(todo *Todo) (*Todo, error) {
	return t.todoRepo.Update(todo)
}

func (t *todoService) Delete(id string) error {
	return t.todoRepo.Delete(id)
}
