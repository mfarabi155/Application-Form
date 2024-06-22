package models

import (
	"time"
)

type LogError struct {
	ID           uint      `gorm:"primaryKey"`
	Timestamp    time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ErrorMessage string    `gorm:"type:text"`
}

func (LogError) TableName() string {
	return "log_errors"
}
