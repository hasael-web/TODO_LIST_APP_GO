package service

import (
	"todo_app/dto"
	"todo_app/models"
	"todo_app/repository"

	"github.com/google/uuid"
)

type TodoService interface {
	GetAllTodo(p *Sort, s *Search) (*dto.TodoResponse, error)
	NewTodo(payload dto.TodoRequest) (*dto.TodoPostResponse, error)
	GetDetail(ID dto.TodoRequestID) (*dto.TodoResponseDetail, error)
	UpdateTodo(ID dto.TodoRequestID, payload models.TodoEntity) (*dto.TodoResponseUpdate, error)
	DeletetTodo(ID dto.TodoRequestID) (*dto.TodoResponseDeletet, error)
}

type TodoServiceImpl struct {
	repository *repository.TodoRepositoryImpl
}

func NewTodoService(tr *repository.TodoRepositoryImpl) *TodoServiceImpl {
	return &TodoServiceImpl{
		repository: tr,
	}
}

type Sort struct {
	Page     int
	PageSize int
}

type Search struct {
	Title       string
	Description string
}

func (tls *TodoServiceImpl) GetAllTodo(p *Sort, s *Search) (*dto.TodoResponse, error) {
	if p.Page == 0 {
		page := 0
		p.Page = page
	}

	if p.PageSize == 0 {
		pageSize := 0
		p.PageSize = pageSize
	}

	if s.Description == "" {
		searchByTitle := ""
		s.Description = searchByTitle
	}

	if s.Title == "" {
		searchByTitle := ""
		s.Title = searchByTitle
	}

	data, err := tls.repository.GetAllTodo(p.Page, p.PageSize, s.Title, s.Description)

	var respons dto.TodoResponse = dto.TodoResponse{
		Code: 200,
		Data: *data,
	}
	if err != nil {
		return nil, err
	}
	return &respons, nil
}

func (tls *TodoServiceImpl) NewTodo(payload dto.TodoRequest) (*dto.TodoPostResponse, error) {

	var new_todo models.TodoEntity = models.TodoEntity{
		ID:          uuid.New(),
		Title:       payload.Data.Title,
		Description: payload.Data.Title,
		Files:       payload.Data.Files,
	}

	data, err := tls.repository.NewTodo(new_todo)

	if err != nil {
		return nil, err
	}

	var respone dto.TodoPostResponse = dto.TodoPostResponse{
		Code: 201,
		Data: *data,
	}

	return &respone, err
}

func (tls *TodoServiceImpl) GetDetail(ID dto.TodoRequestID) (*dto.TodoResponseDetail, error) {
	var response dto.TodoResponseDetail

	data, err := tls.repository.GetDetail(ID.ID)
	if err != nil {
		return nil, err
	}

	response = dto.TodoResponseDetail{
		Code: 200,
		Data: *data,
	}
	return &response, nil
}

func (tls *TodoServiceImpl) UpdateTodo(ID dto.TodoRequestID, payload models.TodoEntity) (*dto.TodoResponseUpdate, error) {
	var response dto.TodoResponseUpdate

	data, err := tls.repository.UpdateTodo(ID.ID, payload)

	if err != nil {
		return nil, err
	}

	response = dto.TodoResponseUpdate{
		Code: 200,
		Data: *data,
	}
	return &response, nil
}

func (tls *TodoServiceImpl) DeletetTodo(ID dto.TodoRequestID) (*dto.TodoResponseDeletet, error) {
	var response dto.TodoResponseDeletet

	data, err := tls.repository.DeletetTodo(ID.ID)
	if err != nil {
		return nil, err
	}

	response = dto.TodoResponseDeletet{
		Code: 200,
		Data: *data,
	}

	return &response, nil
}
