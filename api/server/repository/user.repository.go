package repository

import (
	"bookz/api/server/models"

	"github.com/gin-gonic/gin"
)

type UserRepository interface {
	Login(ctx *gin.Context)
	Register(*gin.Context, models.User)
	IsLogin(ctx *gin.Context)
}
