package controller

import (
	"myapp/app/errors"
	"myapp/app/model"
	"myapp/app/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(userUseCase usecase.UserUseCase) UserController {
	return UserController{
		userUseCase: userUseCase,
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.userUseCase.List()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, users)
}

func (c *UserController) GetUserDetail(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid user ID"))
		return
	}

	user, err := c.userUseCase.GetByID(uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, user)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(errors.NewValidationError("Invalid input"))
		return
	}

	result, err := c.userUseCase.Create(user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(201, result)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid user ID"))
		return
	}

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.Error(errors.NewValidationError("Invalid input"))
		return
	}

	result, err := c.userUseCase.Update(uint(id), user)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, result)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.Error(errors.NewValidationError("Invalid user ID"))
		return
	}

	if err := c.userUseCase.Delete(uint(id)); err != nil {
		ctx.Error(err)
		return
	}
	ctx.Status(204)
}
