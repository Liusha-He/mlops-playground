package services

import (
	"context"
	"errors"
	"net/http"
	"time"
	db "users-service/src/database"
	"users-service/src/dto"
	"users-service/src/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = db.NewCollection(db.Client, "user")
var validate = validator.New()

func GetAllUsersService(ctx *gin.Context)

func GetUserService(ctx *gin.Context) {
	var input dto.GetUserRequest
	var user models.User

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := input.User_Id
	userType := input.User_Type

	if userType != "ADMIN" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("Unauthorised to access the user data"),
		})
		return
	}

	c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	err := userCollection.FindOne(c, bson.M{"user_id": userId}).Decode(&user)
	defer cancel()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
