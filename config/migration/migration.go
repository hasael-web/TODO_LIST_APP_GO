package migration

import (
	"fmt"
	"todo_app/models"

	"gorm.io/gorm"
)

type MigrationDB struct {
	DB *gorm.DB
}

func MigrationInit() *MigrationDB {
	return &MigrationDB{}
}

func (md *MigrationDB) AutoMigration() error {

	err := md.DB.AutoMigrate(
		&models.SubTodoEntityModel{},
		&models.TodoEntityModel{},
	)

	if err != nil {
		fmt.Println("Migration field", err)
	}

	fmt.Println("Migration completed successfully.")
	return nil
}
