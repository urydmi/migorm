package migorm

import (
	"time"
)

type migrationDTO struct {
	Id        uint   `gorm:"primary_key;"`
	Name      string `gorm:"type:varchar(150);unique;not null"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	tableName string `gorm:"-"` // for set custom table name
}

func (m migrationDTO) TableName() string {
	return m.tableName
}
