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

// GetUsers godoc
// @Summary ユーザー一覧取得
// @Description すべてのユーザー情報を取得します
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func (c *UserController) GetUsers(ctx *gin.Context) {
	users, err := c.userUseCase.List()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, users)
}

// GetUserDetail godoc
// @Summary ユーザー詳細取得
// @Description 指定したIDのユーザー情報を取得します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
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

// CreateUser godoc
// @Summary ユーザー作成
// @Description 新しいユーザーを作成します
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "ユーザー情報"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
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

// UpdateUser godoc
// @Summary ユーザー更新
// @Description 指定したIDのユーザー情報を更新します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Param user body model.User true "ユーザー情報"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
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

// DeleteUser godoc
// @Summary ユーザー削除
// @Description 指定したIDのユーザーを削除します
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ユーザーID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
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
