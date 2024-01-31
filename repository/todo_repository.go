package repository

import (
	"errors"
	"todo_app/config/migration"
	"todo_app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodo(page int, pageSize int, searchByTitle string, searchByDescription string) (*[]models.TodoEntityModel, error)
	NewTodo(payload models.TodoEntity) (*models.TodoEntityModel, error)
	GetDetail(ID string) (*models.TodoEntityModel, error)
	DeletetTodo(ID string) (*models.TodoEntityModel, error)
	UpdateTodo(ID string, payload models.TodoEntity) (*models.TodoEntityModel, error)
}

type TodoRepositoryImpl struct {
	migration.MigrationDB
	DB *gorm.DB
}

func NewTodoRepository(tr *migration.MigrationDB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		MigrationDB: *tr,
		DB:          tr.DB,
	}
}

func (tlr *TodoRepositoryImpl) GetAllTodo(page int, pageSize int, searchByTitle string, searchByDescription string) (*[]models.TodoEntityModel, error) {
	var data []models.TodoEntityModel

	if page <= 0 {
		page = 1
	}

	if pageSize <= 0 {
		pageSize = 10
	}

	query := tlr.DB.Limit(pageSize).Offset((page - 1) * pageSize)

	if searchByTitle != "" {
		query = tlr.DB.Where("title LIKE ?", "%"+searchByTitle+"%")
	}

	if searchByDescription != "" {
		query = query.Where("description LIKE ?", "%"+searchByDescription+"%")
	}

	if err := query.Preload("SubList").Find(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (tlr *TodoRepositoryImpl) NewTodo(payload models.TodoEntity) (*models.TodoEntityModel, error) {
	// var create_todo models.TodoEntityModel = models.TodoEntityModel{
	// 	TodoEntity: models.TodoEntity{
	// 		ID:          uuid.New(),
	// 		Title:       payload.Title,
	// 		Description: payload.Description,
	// 		Files:       payload.Files,
	// 		SubList:     payload.SubList,
	// 	},
	// }

	var create_todo models.TodoEntityModel

	create_todo.TodoEntity.ID = uuid.New()
	create_todo.TodoEntity.Title = payload.Title
	create_todo.TodoEntity.Description = payload.Description
	create_todo.TodoEntity.Files = payload.Files
	create_todo.TodoEntity.SubList = payload.SubList

	err := tlr.DB.Create(&create_todo).Error
	if err != nil {
		return nil, err
	}

	return &create_todo, nil
}

func (tlr *TodoRepositoryImpl) GetDetail(ID string) (*models.TodoEntityModel, error) {

	var data models.TodoEntityModel

	err := tlr.DB.Where("id = ?", ID).Preload("SubList").First(&data).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(" todo list not found")
		}
		return nil, err
	}

	return &data, err
}

func (tlr *TodoRepositoryImpl) DeletetTodo(ID string) (*models.TodoEntityModel, error) {

	var todo models.TodoEntityModel

	err := tlr.DB.Where("id = ?", ID).First(&todo).Error

	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo list not found")
		}

		return nil, err
	}

	err = tlr.DB.Delete(&todo).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (tlr *TodoRepositoryImpl) UpdateTodo(ID string, payload models.TodoEntity) (*models.TodoEntityModel, error) {

	var todo models.TodoEntityModel

	err := tlr.DB.Where("id = ?", ID).First(&todo).Error

	if err != nil {
		return nil, err
	}

	if payload.Description != "" {
		todo.TodoEntity.Description = payload.Description
	}
	if payload.Title != "" {
		todo.TodoEntity.Title = payload.Title
	}

	if len(payload.Files) > 0 {
		todo.TodoEntity.Files = payload.Files
	}

	err = tlr.DB.Save(&todo).Error
	if err != nil {
		return nil, err
	}

	return &todo, nil
}
