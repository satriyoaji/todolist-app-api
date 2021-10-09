package service

import (
	"context"
	"satriyoaji/todolist-app-api/model/web/attachment"
)

type AttachmentService interface {
	Create(ctx context.Context, request attachment.AttachmentCreateRequest) attachment.AttachmentResponse
	Update(ctx context.Context, request attachment.AttachmentUpdateRequest) attachment.AttachmentResponse
	Delete(ctx context.Context, attachmentId int)
	FindById(ctx context.Context, attachmentId int) attachment.AttachmentResponse
	FindAll(ctx context.Context) []attachment.AttachmentResponse
}
