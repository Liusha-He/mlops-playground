package services

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"users-service/src/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignupService(ctx *gin.Context) {
	var user models.User
	var count int64
	c, cancel := context.WithTimeout(context.Background(), 100*time.Second)

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := validate.Struct(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	count, err = userCollection.CountDocuments(c, bson.M{"email": user.Email})
	defer cancel()
	if err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred when checking email"})
	}

	password := hashPassword(user.Password)
	user.Password = password

	count, err = userCollection.CountDocuments(c, bson.M{"phone": user.Phone})
	defer cancel()
	if err != nil {
		log.Panic(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred when checking phone"})
	}

	if count > 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "the user email or phone already exist."})
	}

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.User_ID = user.ID.Hex()

	token, refresh_token := generateTokens(&user)
	user.Token = token
	user.Refresh_Token = refresh_token

	result, err := userCollection.InsertOne(c, user)
	if err != nil {
		msg := fmt.Sprintf("user item was not created")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	}
	defer cancel()

	ctx.JSON(http.StatusCreated, result)
}

func LoginService(ctx *gin.Context) {
	c, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	var user models.User
	var foundUser models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err := userCollection.FindOne(c, bson.M{"emil": user.Email}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "email or password incorrect!"})
	}

	pwdIsValid, msg := verifyPassword(user.Password, foundUser.Password)
	defer cancel()

	if pwdIsValid != true {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
	}

	token, refreshTolen := generateTokens(&foundUser)
	updateAllTokens(token, refreshTolen, foundUser.User_ID)

	ctx.JSON(http.StatusOK, foundUser)
}
