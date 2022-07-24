package redis

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AntonyIS/todo-app-one/app"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {

	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, err
	}
	client := redis.NewClient(opts)
	_, err = client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func NewRedisRepository(redisURL string) (app.TodoRepository, error) {
	repo := &redisRepository{}

	client, err := newRedisClient(redisURL)

	if err != nil {
		return nil, errors.Wrap(err, "todo.NewRedisReposiory")
	}

	repo.client = client
	return repo, nil
}

func (r redisRepository) Create(todo *app.Todo) (*app.Todo, error) {

	data := map[string]interface{}{
		"id":          todo.ID,
		"title":       todo.Title,
		"description": todo.Description,
		"created_at":  todo.CreateAt,
	}
	rawmsg, err := json.Marshal(data)
	if err != nil {
		log.Fatal("ERROR MARSHALLING")
	}

	d, err := r.client.HSet("todos", todo.ID, rawmsg).Result()

	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Store")
	}
	fmt.Println("RES", d)

	return todo, nil
}

func (r redisRepository) Read(id string) (*app.Todo, error) {
	todos, err := r.client.HGetAll("todos").Result()

	if err != nil {
		return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.Read")
	}
	todo := &app.Todo{}
	err = json.Unmarshal([]byte(todos[id]), todo)
	if err != nil {
		return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.Read")
	}

	return todo, nil
}

func (r redisRepository) ReadAll() (*[]app.Todo, error) {
	all_data, err := r.client.HGetAll("todos").Result()
	todos := []app.Todo{}
	if err != nil {
		return nil, err
	}

	if len(all_data) == 0 {
		return &todos, nil
	}

	for _, todo := range all_data {
		res := app.Todo{}
		err := json.Unmarshal([]byte(todo), &res)
		if err != nil {
			return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.ReadAll")
		}
		todos = append(todos, res)
	}

	return &todos, nil
}

func (r redisRepository) Update(todo *app.Todo) (*app.Todo, error) {
	todos, err := r.client.HGetAll("todos").Result()

	if err != nil {
		return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.Update")
	}
	res := &app.Todo{}
	err = json.Unmarshal([]byte(todos[todo.ID]), res)

	if err != nil {
		return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.Update")
	}

	res.Title = todo.Title
	res.Description = todo.Description

	rawmsg, err := json.Marshal(res)
	if err != nil {
		return nil, errors.Wrap(app.ErrInvalidTodo, "repository.Todo.Update")
	}

	found, err := r.client.HSet("todos", todo.ID, rawmsg).Result()

	if err != nil {
		return nil, errors.Wrap(err, "repository.Redirect.Store")
	}
	if found {
		return res, nil
	}
	return nil, nil

}

func (r redisRepository) Delete(id string) error {
	_, err := r.client.HDel("todos", id).Result()
	if err != nil {
		return errors.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
