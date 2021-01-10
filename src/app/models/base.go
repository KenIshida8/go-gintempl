package models

import (
	"gintemplate/config"
	"time"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Text   string
	Status string
}


func DbInit() {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		panic("データベース開けず！（dbInit）")
	}
	db.AutoMigrate(&Todo{})
	defer db.Close()
}
