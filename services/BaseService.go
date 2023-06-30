package services

import (
	"dj-lets-go/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	BaseService struct {
		Model      *models.GormModel
		Ctx        *gin.Context
		DbConnName string
		DbConn     *gorm.DB
	}
)
