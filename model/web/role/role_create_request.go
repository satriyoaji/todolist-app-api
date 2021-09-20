package role

type RoleCreateRequest struct {
	Name        string `validate:"required,min=1,max=100" json:"name"`
	Description string `validate:"required,min=1,max=250" json:"description"`
}
