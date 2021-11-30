package models

import (
	"time"
)

// 自增ID主键
type ID struct {
	Id uint `gorm:"primary_key;AUTO_INCREMENT"`
}

// 创建、更新时间
type Timestamps struct {
	CreateTime time.Time `gorm:"not null"`
	UpdateTime time.Time `gorm:"not null"`
}
