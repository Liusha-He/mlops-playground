package controllers

import (
	"users-service/src/services"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary login
// @Schemes
// @Description user login
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/auth/login [post]
func Login(ctx *gin.Context) {
	services.LoginService(ctx)
}

// Signup godoc
// @Summary signup
// @Schemes
// @Description new user register
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/signup [post]
func Signup(ctx *gin.Context) {
	services.SignupService(ctx)
}

// GetAllUsers godoc
// @Summary GetAllUsers
// @Schemes
// @Description get all users
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Router /api/v1/users [post]
func GetAllUsers(ctx *gin.Context) {
	services.GetAllUsersService(ctx)
}

// GetUser godoc
// @Summary GetUser
// @Schemes
// @Description get detail of a user by user id and user type
// @Tags auth
// @Accept json
// @Produce json
// @Success 200
// @Router v1//users/{user_id} [get]
func GetUser(ctx *gin.Context) {
	services.GetUserService(ctx)
}
