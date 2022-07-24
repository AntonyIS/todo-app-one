package app

type TodoSerializer interface {
	Decode(input []byte) (*Todo, error)
	Encode(input *Todo) ([]byte, error)
}
