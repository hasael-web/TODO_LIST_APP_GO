package dto

import "todo_app/models"

type TodoRequest struct {
	Data models.TodoEntityModel
}

type TodoResponse struct {
	Code int                      `json:"code"`
	Data []models.TodoEntityModel `json:"data"`
}

type TodoPostResponse struct {
	Code int                    `json:"code"`
	Data models.TodoEntityModel `json:"data"`
}

type TodoRequestID struct {
	ID string
}

type TodoResponseDetail struct {
	Code int                    `json:"code"`
	Data models.TodoEntityModel `json:"data"`
}

type TodoResponseDeletet struct {
	Code int                    `json:"code"`
	Data models.TodoEntityModel `json:"data"`
}

type TodoResponseUpdate struct {
	Code int                    `json:"code"`
	Data models.TodoEntityModel `json:"data"`
}
