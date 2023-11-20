package global

import "gorm.io/gorm"

const (
	Version = "1.0.0"
)

var (
	DB *gorm.DB
)
