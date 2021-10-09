package helper

import (
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/attachment"
)

func ToAttachmentResponse(value domain.Attachment) attachment.AttachmentResponse {
	return attachment.AttachmentResponse{
		Id:        value.Id,
		TodoId:    value.TodoId,
		Location:  value.Location,
		Caption:   value.Caption,
		CreatedAt: value.CreatedAt,
		UpdatedAt: value.UpdatedAt,
	}
}

func ToAttachmentResponses(attachments []domain.Attachment) []attachment.AttachmentResponse {
	var attachmentResponses []attachment.AttachmentResponse
	for _, attachment := range attachments {
		attachmentResponses = append(attachmentResponses, ToAttachmentResponse(attachment))
	}
	return attachmentResponses
}
