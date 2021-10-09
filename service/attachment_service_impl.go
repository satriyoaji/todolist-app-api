package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"satriyoaji/todolist-app-api/exception"
	"satriyoaji/todolist-app-api/helper"
	"satriyoaji/todolist-app-api/model/domain"
	"satriyoaji/todolist-app-api/model/web/attachment"
	"satriyoaji/todolist-app-api/repository"
)

type AttachmentServiceImpl struct {
	AttachmentRepository repository.AttachmentRepository
	DB                   *sql.DB
	Validate             *validator.Validate
}

func NewAttachmentService(attachmentRepository repository.AttachmentRepository, DB *sql.DB, validate *validator.Validate) AttachmentService {
	return &AttachmentServiceImpl{
		AttachmentRepository: attachmentRepository,
		DB:                   DB,
		Validate:             validate,
	}
}

func (service *AttachmentServiceImpl) Create(ctx context.Context, request attachment.AttachmentCreateRequest) attachment.AttachmentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attachment := domain.Attachment{
		TodoId:  request.TodoId,
		Caption: request.Caption,
		File:    request.File,
	}

	attachment = service.AttachmentRepository.Save(ctx, tx, attachment)

	return helper.ToAttachmentResponse(attachment)
}

func (service *AttachmentServiceImpl) Update(ctx context.Context, request attachment.AttachmentUpdateRequest) attachment.AttachmentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attachment, err := service.AttachmentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	attachment.TodoId = request.TodoId
	attachment.Caption = request.Caption
	attachment.File = request.File

	attachment = service.AttachmentRepository.Update(ctx, tx, attachment)

	return helper.ToAttachmentResponse(attachment)
}

func (service *AttachmentServiceImpl) Delete(ctx context.Context, attachmentId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attachment, err := service.AttachmentRepository.FindById(ctx, tx, attachmentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.AttachmentRepository.Delete(ctx, tx, attachment)
}

func (service *AttachmentServiceImpl) FindById(ctx context.Context, attachmentId int) attachment.AttachmentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attachment, err := service.AttachmentRepository.FindById(ctx, tx, attachmentId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAttachmentResponse(attachment)
}

func (service *AttachmentServiceImpl) FindAll(ctx context.Context) []attachment.AttachmentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	attachments := service.AttachmentRepository.FindAll(ctx, tx)

	return helper.ToAttachmentResponses(attachments)
}
