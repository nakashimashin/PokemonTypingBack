package controller

import (
	"back/model"
	"back/usecase"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	Signup(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type userController struct {
	uu usecase.UserUseCase
}

func NewUserController(uu usecase.UserUseCase) UserController {
	return &userController{uu}
}

func (uc *userController) Signup(c *gin.Context) {
	user := model.User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	err := uc.uu.Signup(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, nil)
}

func (uc *userController) Login(c *gin.Context) {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	tokenString, err := uc.uu.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(1 * time.Hour),
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, tokenString)
}

func (uc *userController) Logout(c *gin.Context) {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		Path:     "/",
		Domain:   "localhost",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, cookie)
	c.JSON(http.StatusOK, nil)
}
