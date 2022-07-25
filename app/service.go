package app

type TodoService interface {
	Create(todo *Todo) (*Todo, error)
	Read(id string) (*Todo, error)
	ReadAll() (*[]Todo, error)
	Update(todo *Todo) (*Todo, error)
	Delete(id string) error
}

type UserService interface {
	Create(user *User) (*User, error)
	Read(id string) (*User, error)
	ReadAll() (*[]User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
}
