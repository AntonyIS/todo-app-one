package app

type TodoRepository interface {
	Create(todo *Todo) (*Todo, error)
	Read(id string) (*Todo, error)
	ReadAll() (*[]Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id string) error
}
