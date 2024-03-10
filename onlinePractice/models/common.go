package models

import (
	"database/sql"
	"time"
)

type model struct {
	ID        uint `gorm:"column:id;primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
