package app


type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreateAt int64 `json:"created_at"`
}

type User struct {
	ID string `json:"id"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
	Avater string `json:"avater"`
	Todo []Todo
}