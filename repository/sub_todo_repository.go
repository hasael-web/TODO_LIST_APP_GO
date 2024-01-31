package repository

import (
	"errors"
	"todo_app/config/migration"
	"todo_app/models"

	"gorm.io/gorm"
)

type SubTodoRepository interface {
	GetAllSubTodo(page int, pageSize int, searchByTitle string, searchByDescription string, ListID string) (*[]models.SubTodoEntityModel, error)
	NewSubTodo(payload models.SubTodoEntity) (*models.SubTodoEntityModel, error)
	GetDetailSubTodo(ID string) (*models.SubTodoEntityModel, error)
	DeletetSubTodo(ID string) (*models.SubTodoEntityModel, error)
	UpdateSubTodo(ID string, payload models.SubTodoEntity) (*models.SubTodoEntityModel, error)
}

type SubTodoRepositoryImpl struct {
	migration.MigrationDB
	DB *gorm.DB
}

func NewSubTodoRepository(md *migration.MigrationDB) *SubTodoRepositoryImpl {
	return &SubTodoRepositoryImpl{
		MigrationDB: *md,
		DB:          md.DB,
	}
}

func (str *SubTodoRepositoryImpl) GetAllSubTodo(page int, pageSize int, searchByTitle string, searchByDescription string, ListID string) (*[]models.SubTodoEntityModel, error) {
	var data []models.SubTodoEntityModel

	query := str.DB.Where("list_id =?", ListID)

	if searchByTitle != "" {
		query = query.Where("title LIKE ?", "%"+searchByTitle+"%")
	}

	if searchByDescription != "" {
		query = query.Where("description ?", "%"+searchByDescription+"%")
	}

	query = query.Order("id DESC")

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	query = query.Limit(pageSize).Offset((page - 1) * pageSize)

	if err := query.Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return nil, err
	}
	return &data, nil
}

func (str *SubTodoRepositoryImpl) NewSubTodo(payload models.SubTodoEntity) (*models.SubTodoEntityModel, error) {
	var create_sub_todo models.SubTodoEntityModel = models.SubTodoEntityModel{
		SubTodoEntity: models.SubTodoEntity{
			ID:          payload.ID,
			Title:       payload.Title,
			Description: payload.Description,
			Files:       payload.Files,
			ListId:      payload.ListId,
		},
	}

	err := str.DB.Create(&create_sub_todo).Error

	if err != nil {
		return nil, err
	}

	return &create_sub_todo, nil

}

func (str *SubTodoRepositoryImpl) GetDetailSubTodo(ID string) (*models.SubTodoEntityModel, error) {
	var sub_todo models.SubTodoEntityModel

	err := str.DB.Where("id = ?", ID).First(&sub_todo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	return &sub_todo, nil
}

func (str *SubTodoRepositoryImpl) DeletetSubTodo(ID string) (*models.SubTodoEntityModel, error) {
	var sub_todo models.SubTodoEntityModel

	err := str.DB.Where("id = ?", ID).First(&sub_todo).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	err = str.DB.Delete(&sub_todo).Error

	if err != nil {
		return nil, err
	}

	return &sub_todo, nil

}

func (str *SubTodoRepositoryImpl) UpdateSubTodo(ID string, payload models.SubTodoEntity) (*models.SubTodoEntityModel, error) {
	var sub_todo models.SubTodoEntityModel
	err := str.DB.Where("id = ?", ID).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}

	if payload.Description != "" {
		sub_todo.SubTodoEntity.Description = payload.Description
	}

	if payload.Title != "" {
		sub_todo.SubTodoEntity.Title = payload.Title
	}

	if len(payload.Files) > 0 {
		sub_todo.SubTodoEntity.Files = payload.Files
	}

	err = str.DB.Save(&sub_todo).Error
	if err != nil {
		return nil, err
	}

	return &sub_todo, nil

}
