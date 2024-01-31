package models

import (
	"todo_app/utils/abstraction"

	"github.com/google/uuid"
)

type SubTodoEntity struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Title       string    `json:"title"  gorm:"type:varchar(100)"`
	Description string    `json:"description" gorm:"type:text"`
	Files       Files     `json:"files" gorm:"type:varchar(255)[]"`
	ListId      uuid.UUID `json:"list_id"`
}

type SubTodoEntityModel struct {
	SubTodoEntity
	abstraction.Entity
}

func (SubTodoEntityModel) TableName() string {
	return "sub_todo_list"
}
