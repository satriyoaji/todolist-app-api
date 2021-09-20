package role

type RoleUpdateRequest struct {
	Id          int    `validate:"required"`
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=250" json:"description"`
}
