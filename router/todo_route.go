package router

import (
	"todo_app/handler"

	"github.com/labstack/echo/v4"
)

type TodoRouteImpl struct {
	handler *handler.TodoHandlerImpl
}

func NewTodoRoute(h *handler.TodoHandlerImpl) *TodoRouteImpl {
	return &TodoRouteImpl{
		handler: h,
	}
}

func (r *TodoRouteImpl) Route(e *echo.Group) {
	e.GET("/todo/lists", r.handler.GetAll)
	e.POST("/todo/list", r.handler.Created)
	e.GET("/todo/list/:id", r.handler.Detail)
	e.DELETE("/todo/list/delete/:id", r.handler.Deletet)
	e.PATCH("/todo/list/update/:id", r.handler.Update)
}
