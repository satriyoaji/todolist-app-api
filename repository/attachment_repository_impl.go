package repository

import (
	"context"
	"database/sql"
	"errors"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"time"
)

type AttachmentRepositoryImpl struct {
}

func NewAttachmentRepository() AttachmentRepository {
	return &AttachmentRepositoryImpl{}
}

func (repository *AttachmentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, attachment domain.Attachment) domain.Attachment {
	SQL := "insert into attachments(todo_id, file, caption, created_at, updated_at) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, attachment.TodoId, attachment.File, attachment.Caption, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	attachment.Id = int(id)
	return attachment
}

func (repository *AttachmentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, attachment domain.Attachment) domain.Attachment {
	SQL := "update attachments set todo_id = ?, file = ?, caption = ?, updated_at = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, attachment.TodoId, attachment.File, attachment.Caption, time.Now(), attachment.Id)
	helper.PanicIfError(err)

	return attachment
}

func (repository *AttachmentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, attachment domain.Attachment) {
	SQL := "delete from attachments where id = ?"
	_, err := tx.ExecContext(ctx, SQL, attachment.Id)
	helper.PanicIfError(err)
}

func (repository *AttachmentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, attachmentId int) (domain.Attachment, error) {
	SQL := "select id, todo_id, file, caption from attachments where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, attachmentId)
	helper.PanicIfError(err)
	defer rows.Close()

	attachment := domain.Attachment{}
	if rows.Next() {
		err := rows.Scan(&attachment.Id, &attachment.TodoId, &attachment.File, &attachment.Caption)
		helper.PanicIfError(err)
		return attachment, nil
	} else {
		return attachment, errors.New("attachment is not found")
	}
}

func (repository *AttachmentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Attachment {
	SQL := "select id, todo_id, file, caption from attachments"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var attachments []domain.Attachment
	for rows.Next() {
		attachment := domain.Attachment{}
		err := rows.Scan(&attachment.Id, &attachment.TodoId, &attachment.File, &attachment.Caption)
		helper.PanicIfError(err)
		attachments = append(attachments, attachment)
	}
	return attachments
}
