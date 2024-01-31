package router

import (
	"todo_app/handler"

	"github.com/labstack/echo/v4"
)

type SubTodoRouteImpl struct {
	handler *handler.SubTodoHandlerImpl
}

func NewSubTodoRoute(h *handler.SubTodoHandlerImpl) *SubTodoRouteImpl {
	return &SubTodoRouteImpl{
		handler: h,
	}
}

func (r *SubTodoRouteImpl) Route(e *echo.Group) {
	e.GET("/todo/sublists/:list_id", r.handler.GetAll)
	e.GET("/todo/sublist/:id", r.handler.Detail)
	e.POST("/todo/sublist", r.handler.Created)
	e.DELETE("/todo/sublist/delete/:id", r.handler.Delete)
	e.PATCH("/todo/sublist/update/:id", r.handler.Update)
}
