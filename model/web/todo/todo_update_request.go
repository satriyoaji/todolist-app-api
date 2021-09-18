package todo

type TodoUpdateRequest struct {
	Id   int    `validate:"required"`
	UserId int `validate:"required,numeric" json:"user_id"`
	Title string `validate:"required,min=1,max=100" json:"title"`
}
