package dto

import "todo_app/models"

type SubTodoGetResponse struct {
	Code int                         `json:"code"`
	Data []models.SubTodoEntityModel `json:"data"`
}

type SubTodoPostResponse struct {
	Code int                       `json:"code"`
	Data models.SubTodoEntityModel `json:"data"`
}

type SubTodoGetDetail struct {
	Code int                       `json:"code"`
	Data models.SubTodoEntityModel `json:"data"`
}

type SubTodoDeleteResponse struct {
	Code int                       `json:"code"`
	Data models.SubTodoEntityModel `json:"data"`
}

type SubTodoUpdateResponse struct {
	Code int                       `json:"code"`
	Data models.SubTodoEntityModel `json:"data"`
}

type SubTodoPostRequest struct {
	Data models.SubTodoEntity
}
