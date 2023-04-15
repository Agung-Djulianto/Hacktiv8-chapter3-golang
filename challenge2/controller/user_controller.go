package controller

import (
	"chap3-challenge2/model"
	"chap3-challenge2/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	newUser := model.UserRegisterRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	res, err := uc.UserService.AddUser(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.UserRegisterResponse{
		ID:    res.ID,
		Email: res.Email,
		Role:  res.Role,
	})
	return
}

func (uc *UserController) CreateAdmin(ctx *gin.Context) {
	newUser := model.UserRegisterRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	res, err := uc.UserService.AddAdmin(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.UserRegisterResponse{
		ID:    res.ID,
		Email: res.Email,
		Role:  res.Role,
	})
	return
}

func (uc *UserController) Login(ctx *gin.Context) {
	request := model.UserLoginRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	res, err := uc.UserService.Login(request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.UserLoginResponse{
		Token: res.Token,
	})
	return

}
