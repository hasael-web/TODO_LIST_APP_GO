package models

import (
	"database/sql/driver"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Files []string

type TodoEntity struct {
	ID          uuid.UUID            `json:"id"  gorm:"type:uuid;primaryKey"`
	Title       string               `json:"title" validate:"required,min=5,max=100" gorm:"type:varchar(100)"`
	Description string               `json:"description" validate:"required,min=5,max=100" gorm:"type:text"`
	Files       Files                `json:"files"  gorm:"type:varchar(255)[]"`
	SubList     []SubTodoEntityModel `json:"sublist" gorm:"foreignKey:ListId"`
}

type Entity struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type TodoEntityModel struct {
	TodoEntity
	Entity
}

type RTdodoEntitiModel struct {
	ID          uuid.UUID       `json:"id"  gorm:"type:uuid;primaryKey"`
	Title       string          `json:"title" validate:"required,min=5,max=100" gorm:"type:varchar(100)"`
	Description string          `json:"description" validate:"required,min=5,max=100" gorm:"type:text"`
	Files       Files           `json:"files"  gorm:"type:varchar(255)[]"`
	SubList     []SubTodoEntity `json:"sublist" gorm:"foreignKey:ListId"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt  `json:"-" gorm:"index"`
}

func (TodoEntityModel) TableName() string {
	return "todo_list"
}

func (f *Files) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*f = strings.Split(string(v), ",")
		return nil
	case string:
		*f = strings.Split(v, ",")
		return nil
	default:
		return errors.New("src value cannot cast to []string")
	}
}

func (f Files) Value() (driver.Value, error) {
	if len(f) == 0 {
		return nil, nil
	}

	return "{" + strings.Join(f, ",") + "}", nil
}
