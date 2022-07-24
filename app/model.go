package app


type Todo struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	CreateAt int64 `json:"created_at"`
}