package handler

import (
	"bookz/api/server/dto"
	"bookz/api/server/helper"
	"bookz/api/server/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserRepository repository.UserRepository
}

func (u *UserHandler) Login(ctx *gin.Context) {
	req := dto.RegisterRequestDTO{}

	defer ctx.Request.Body.Close()
	ctx.Bind(&req)

	ctx.JSON(http.StatusOK, helper.Response{
		StatusCode: 1,
		Message:    "",
		Data:       req,
	})
}

func (u *UserHandler) Register(ctx *gin.Context) {
	req := dto.RegisterRequestDTO{}

	defer ctx.Request.Body.Close()
	ctx.Bind(&req)

	ctx.JSON(http.StatusOK, helper.Response{
		StatusCode: 1,
		Message:    "",
		Data:       req,
	})

}
