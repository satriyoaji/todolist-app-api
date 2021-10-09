package attachment

type AttachmentUpdateRequest struct {
	Id       int    `validate:"required"`
	TodoId   int    `validate:"required,numeric" json:"todo_id"`
	Location string `validate:"required,min=1,max=100" json:"location"`
	Caption  string `validate:"required,min=1,max=100" json:"caption"`
}
