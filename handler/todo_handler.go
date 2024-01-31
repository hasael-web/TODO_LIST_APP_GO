package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"todo_app/config"
	"todo_app/dto"
	"todo_app/models"
	"todo_app/service"

	"github.com/labstack/echo/v4"
)

type TodoHandlerImpl struct {
	service *service.TodoServiceImpl
}

func NewTodoHandler(s *service.TodoServiceImpl) *TodoHandlerImpl {
	return &TodoHandlerImpl{
		service: s,
	}
}

// GET ALL TODOS LIST
// @Summary All Todo List
// @Descripton Displays All Todolist Data
// @Tags TODO LIST
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination (default is 1)"
// @Param page_size query int false "Number of items per page (default is 10)"
// @Param title query string false "Search by title (case-insensitive)"
// @Param description query string false "Search by description (case-insensitive)"
// @Success 200 {array} dto.TodoResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/lists [get]
func (h *TodoHandlerImpl) GetAll(e echo.Context) error {
	page, _ := strconv.Atoi(e.QueryParam("page"))
	pageSize, _ := strconv.Atoi(e.QueryParam("page_size"))
	searchByTitle := e.QueryParam("title")
	searchByDescription := e.QueryParam("description")

	sort := service.Sort{Page: page, PageSize: pageSize}
	search := service.Search{Title: searchByTitle, Description: searchByDescription}

	data, err := h.service.GetAllTodo(&sort, &search)

	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, data)

}

// POST (CREATED) NEW TODO LIST
// @Summary create new todo list
// @Descripton create a new todo list
// @Tags TODO LIST
// @Accept json
// @Produce json
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param files formData file true "Files"
// @Success 201 {object} dto.TodoResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/list [post]
func (h *TodoHandlerImpl) Created(e echo.Context) error {
	title := e.FormValue("title")
	descripton := e.FormValue("descripton")
	// files
	form, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	files := form.File["files"]

	var payload dto.TodoRequest = dto.TodoRequest{
		Data: models.TodoEntityModel{
			TodoEntity: models.TodoEntity{
				Title:       title,
				Description: descripton,
				Files:       nil,
			},
		},
	}

	for _, file := range files {
		filePath, err := config.UploadFileToLocalDir(file)

		if err != nil {

			return e.JSON(http.StatusInternalServerError, err)
		}
		ext := strings.ToLower(filepath.Ext(filePath)[1:])

		if ext != "pdf" && ext != "txt" {
			config.DeleteUploadedFile(filePath)
			return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid file extension. Only PDF or TXT allowed."})
		}
		payload.Data.TodoEntity.Files = append(payload.Data.TodoEntity.Files, filePath)
	}

	data, err := h.service.NewTodo(payload)

	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}
	return e.JSON(http.StatusCreated, data)

}

// GET TODO BY ID
// @Summary Detail Todo
// @Descripton see more detail todo list
// @Tags TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} dto.TodoResponseDetail
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/list/{id} [get]
func (h *TodoHandlerImpl) Detail(e echo.Context) error {
	var ID dto.TodoRequestID

	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id params cannot be empty"})
	}

	ID = dto.TodoRequestID{
		ID: id,
	}

	result, err := h.service.GetDetail(ID)

	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, result)
}

// DELETE TODO LIST
// @Summary Delete todo list by id list
// @Descripton delete list
// @Tags TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} dto.TodoResponseDeletet
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/list/delete/{id} [delete]
func (h *TodoHandlerImpl) Deletet(e echo.Context) error {
	var ID dto.TodoRequestID

	id := e.Param("id")

	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id params cannot be empty"})
	}

	ID = dto.TodoRequestID{
		ID: id,
	}

	result, err := h.service.DeletetTodo(ID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, result)
}

// UPDATE TODO LIST
// @Summary Update todo list by id todo
// @Descripton update todo list
// @Tags TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param files formData file true "Files"
// @Success 200 {object} dto.TodoResponseUpdate
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/list/update/{id} [patch]
func (h *TodoHandlerImpl) Update(e echo.Context) error {
	var ID dto.TodoRequestID

	id := e.Param("id")

	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id params cannot be empty"})
	}

	ID = dto.TodoRequestID{
		ID: id,
	}

	var updateData models.TodoEntity

	title := e.FormValue("title")
	descripton := e.FormValue("description")

	if title != "" {
		updateData.Title = title
	}

	if descripton != "" {
		updateData.Description = descripton
	}

	// files
	form, err := e.MultipartForm()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	files := form.File["files"]

	if len(files) > 0 {

		for _, file := range files {
			filePath, err := config.UploadFileToLocalDir(file)

			if err != nil {

				return e.JSON(http.StatusInternalServerError, err)
			}
			ext := strings.ToLower(filepath.Ext(filePath)[1:])

			if ext != "pdf" && ext != "txt" {
				config.DeleteUploadedFile(filePath)
				return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid file extension. Only PDF or TXT allowed."})
			}
			updateData.Files = append(updateData.Files, filePath)
		}
	}

	result, err := h.service.UpdateTodo(ID, updateData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, result)

}
