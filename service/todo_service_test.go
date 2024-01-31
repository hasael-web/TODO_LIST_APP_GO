package service

import (
	"testing"
	"todo_app/config/migration"
	"todo_app/repository"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTodo(t *testing.T) {

	realRepo := repository.NewTodoRepository(&migration.MigrationDB{})

	todoService := NewTodoService(realRepo)

	result, err := todoService.GetAllTodo(&Sort{Page: 0, PageSize: 0}, &Search{Title: "", Description: ""})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
