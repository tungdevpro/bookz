package handler

import (
	"bookz/api/server/dto"
	"bookz/api/server/helper"
	"bookz/api/server/middleware"
	"bookz/api/server/models"
	"bookz/api/server/repository"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	if err := ctx.Bind(&req); err != nil {
		ctx.JSON(http.StatusAccepted, helper.ErrorResponse{
			StatusCode: -1,
			Message:    "",
		})
		return
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.ErrorResponse{
			StatusCode: -1,
			Message:    err.Error(),
		})
		return
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	helper.Fatal(err)

	req.Password = string(bs[:])

	user := models.User{
		Name:     req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Role:     "guest",
	}

	u.UserRepository.Register(ctx, user)

	token, err := middleware.GenToken(user)
	helper.Fatal(err)

	user.AccessToken = token

	ctx.JSON(http.StatusOK, helper.Response{
		StatusCode: 1,
		Message:    "",
		Data:       user,
	})
}
