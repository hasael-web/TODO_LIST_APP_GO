package service

import (
	"todo_app/dto"
	"todo_app/models"
	"todo_app/repository"

	"github.com/google/uuid"
)

type SubTodoServiceImpl struct {
	repository *repository.SubTodoRepositoryImpl
}

func NewSubTodoService(tr *repository.SubTodoRepositoryImpl) *SubTodoServiceImpl {
	return &SubTodoServiceImpl{
		repository: tr,
	}
}

func (stls *SubTodoServiceImpl) GetAllSubTodo(page int, pageSize int, searchByTitle string, searchByDescription string, ListID string) (*dto.SubTodoGetResponse, error) {
	if page == 0 {
		page = 0
	}

	if pageSize == 0 {
		pageSize = 0
	}

	if searchByDescription == "" {
		searchByDescription = ""
	}

	if searchByTitle == "" {
		searchByTitle = ""
	}

	data, err := stls.repository.GetAllSubTodo(page, pageSize, searchByTitle, searchByDescription, ListID)

	if err != nil {
		return nil, err
	}
	var result dto.SubTodoGetResponse = dto.SubTodoGetResponse{
		Code: 200,
		Data: *data,
	}

	return &result, nil
}

func (stls *SubTodoServiceImpl) NewSubTodo(payload models.SubTodoEntity) (*dto.SubTodoPostResponse, error) {
	var new_sub_todo models.SubTodoEntity = models.SubTodoEntity{
		ID:          uuid.New(),
		Title:       payload.Title,
		Description: payload.Description,
		Files:       payload.Files,
		ListId:      payload.ListId,
	}

	data, err := stls.repository.NewSubTodo(new_sub_todo)

	if err != nil {
		return nil, err
	}

	var result dto.SubTodoPostResponse = dto.SubTodoPostResponse{
		Code: 200,
		Data: *data,
	}

	return &result, nil
}

func (stls *SubTodoServiceImpl) GetDetailSubTodo(ID string) (*dto.SubTodoGetDetail, error) {
	var response dto.SubTodoGetDetail

	data, err := stls.repository.GetDetailSubTodo(ID)
	if err != nil {
		return nil, err
	}

	response = dto.SubTodoGetDetail{
		Code: 200,
		Data: *data,
	}

	return &response, nil
}

func (stls *SubTodoServiceImpl) DeletetSubTodo(ID string) (*dto.SubTodoDeleteResponse, error) {
	var response dto.SubTodoDeleteResponse

	data, err := stls.repository.DeletetSubTodo(ID)
	if err != nil {
		return nil, err
	}

	response = dto.SubTodoDeleteResponse{
		Code: 200,
		Data: *data,
	}

	return &response, err
}

func (stls *SubTodoServiceImpl) UpdateSubTodo(ID string, payload models.SubTodoEntity) (*dto.SubTodoUpdateResponse, error) {
	var respone dto.SubTodoUpdateResponse

	data, err := stls.repository.UpdateSubTodo(ID, payload)

	if err != nil {
		return nil, err
	}

	respone = dto.SubTodoUpdateResponse{
		Code: 200,
		Data: *data,
	}

	return &respone, nil
}
