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

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type SubTodoHandlerImpl struct {
	service *service.SubTodoServiceImpl
}

func NewSubTodoHandler(s *service.SubTodoServiceImpl) *SubTodoHandlerImpl {
	return &SubTodoHandlerImpl{
		service: s,
	}
}

// POST (CREATED) SUB TODO LIST
// @Summary Add sub lists to the todo list
// @Descripton create a new sub todo list
// @Tags SUB TODO LIST
// @Accept json
// @Produce json
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param files formData file true "Files"
// @Param listid formData string true "List ID"
// @Success 201 {object} dto.SubTodoPostResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/sublist [post]
func (sh *SubTodoHandlerImpl) Created(e echo.Context) error {
	title := e.FormValue("title")
	description := e.FormValue("description")
	listId, err := uuid.Parse(e.FormValue("listid"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}
	// file
	form, err := e.MultipartForm()

	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	files := form.File["sub-files"]

	var payload dto.SubTodoPostRequest = dto.SubTodoPostRequest{
		Data: models.SubTodoEntity{
			Title:       title,
			Description: description,
			ListId:      listId,
			Files:       nil,
		},
	}

	for _, file := range files {
		filePath, err := config.UploadFileToLocalDir(file)
		if err != nil {
			return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		}

		ext := strings.ToLower(filepath.Ext(filePath)[1:])

		if ext != "pdf" && ext != "txt" {
			config.DeleteUploadedFile(filePath)
			return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "Invalid file extension. Only PDF or TXT allowed."})
		}

		payload.Data.Files = append(payload.Data.Files, filePath)

	}

	data, err := sh.service.NewSubTodo(payload.Data)
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusCreated, data)
}

// GET ALL SUB TODOS LIST
// @Summary Get All sub todo list
// @Descripton See more detail sub todo list by id todo list
// @Tags SUB TODO LIST
// @Accept json
// @Produce json
// @Param list_id path string true "List id"
// @Param page query int false "Page number for pagination (default is 1)"
// @Param page_size query int false "Number of items per page (default is 10)"
// @Param title query string false "Search by title (case-insensitive)"
// @Param description query string false "Search by description (case-insensitive)"
// @Success 200 {array} dto.SubTodoGetResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/sublists/{list_id} [get]
func (sh *SubTodoHandlerImpl) GetAll(e echo.Context) error {
	page, _ := strconv.Atoi(e.QueryParam("page"))
	pageSize, _ := strconv.Atoi(e.QueryParam("page_size"))
	searchByTitle := e.QueryParam("title")
	searchByDescription := e.QueryParam("description")
	listID := e.Param("list_id")

	if listID == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id params cannot be empty"})
	}

	data, err := sh.service.GetAllSubTodo(page, pageSize, searchByTitle, searchByDescription, listID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, data)
}

// GET SUB TODO LIST DETAIL
// @Summary Detail Sub Todo
// @Descripton see more detail sub todo list
// @Tags SUB TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Sub Todo List ID"
// @Success 200 {object} dto.SubTodoGetDetail
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/sublist/update/{id} [get]
func (sh *SubTodoHandlerImpl) Detail(e echo.Context) error {
	id := e.Param("id")
	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id cannot be empty"})
	}
	result, err := sh.service.GetDetailSubTodo(id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, result)
}

// UPDATE SUB TODO LIST
// @Summary Update sub todo list by id sub todo
// @Description update sub todo list
// @Tags SUB TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Sub Todo ID"
// @Param title formData string true "Title"
// @Param description formData string true "Description"
// @Param files formData file true "Files"
// @Success 200 {object} dto.SubTodoUpdateResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/sublist/update/{id} [patch]
func (h *SubTodoHandlerImpl) Update(e echo.Context) error {
	id := e.Param("id")

	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id param cannot be empty"})
	}

	var updateData models.SubTodoEntity

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
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
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

	result, err := h.service.UpdateSubTodo(id, updateData)

	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}

	return e.JSON(http.StatusOK, result)

}

// DELETE SUB TODO LIST
// @Summary Delete sub todo list
// @Descripton delete sub todo list
// @Tags SUB TODO LIST
// @Accept json
// @Produce json
// @Param id path string true "Sub Todo ID"
// @Success 200 {object} dto.SubTodoDeleteResponse
// @Failure 400 {object} dto.ErrorResponse
// @Router /todo/sublist/delete/{id} [delete]
func (sh *SubTodoHandlerImpl) Delete(e echo.Context) error {
	id := e.Param("id")

	if id == "" {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: "id params cannot be empty"})
	}

	result, err := sh.service.DeletetSubTodo(id)

	if err != nil {
		return e.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
	}

	return e.JSON(http.StatusOK, result)
}
