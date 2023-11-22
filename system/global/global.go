package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	Version = "1.0.0"
)

var (
	DB     *gorm.DB
	Router *gin.Engine
)
