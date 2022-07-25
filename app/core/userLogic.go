package core

import (
	"errors"

	"github.com/AntonyIS/todo-app-one/app"
	"github.com/teris-io/shortid"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrInvalidUser  = errors.New("invalid user")
	ErrNoUser       = errors.New("users not available user")
)

type userService struct {
	userRepo app.UserRepository
}

func NewUserService(userRepo app.UserRepository) app.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Create(user *app.User) (*app.User, error) {
	user.ID = shortid.MustGenerate()
	return u.userRepo.Create(user)
}

func (u *userService) Read(id string) (*app.User, error) {
	return u.userRepo.Read(id)
}
func (u *userService) ReadAll() (*[]app.User, error) {
	return u.userRepo.ReadAll()
}

func (u *userService) Update(user *app.User) (*app.User, error) {
	return u.userRepo.Update(user)
}

func (u *userService) Delete(id string) error {
	return u.userRepo.Delete(id)
}
