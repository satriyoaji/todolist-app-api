package repository

import (
	"context"
	"database/sql"
	"satriyoaji/todolist-app-api/model/domain"
)

type AttachmentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, attachment domain.Attachment) domain.Attachment
	Update(ctx context.Context, tx *sql.Tx, attachment domain.Attachment) domain.Attachment
	Delete(ctx context.Context, tx *sql.Tx, attachment domain.Attachment)
	FindById(ctx context.Context, tx *sql.Tx, attachmentId int) (domain.Attachment, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Attachment
}
