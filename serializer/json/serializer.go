package json

import (
	"encoding/json"

	"github.com/AntonyIS/todo-app-one/app"
	"github.com/pkg/errors"
)

type Todo struct{}

func (t *Todo) Decode(input []byte) (*app.Todo, error) {
	todo := &app.Todo{}
	if err := json.Unmarshal(input, todo); err != nil {
		return nil, errors.Wrap(err, "serializer.Todo.Decode")
	}
	return todo, nil
}

func (t *Todo) Encode(input *app.Todo) ([]byte, error) {
	rawMsg, err := json.Marshal(input)

	if err != nil {
		return nil, errors.Wrap(err, "serializer.Todo.Encode")
	}

	return rawMsg, nil
}
