package attachment

type AttachmentCreateRequest struct {
	TodoId  int    `validate:"required,numeric" json:"todo_id"`
	File    string `validate:"required,min=1,max=100" json:"file"`
	Caption string `validate:"required,min=1,max=100" json:"caption"`
}
