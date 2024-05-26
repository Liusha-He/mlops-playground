package services

import (
	"log"
	"os"
	"time"
	"users-service/src/models"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type signeddetails struct {
	email      string
	first_name string
	last_name  string
	uid        string
	user_type  string
	jwt.StandardClaims
}

func generateTokens(user *models.User) (string, string) {
	secret_key := os.Getenv("SECRET_KEY")

	claims := &signeddetails{
		email:      user.Email,
		first_name: user.First_Name,
		last_name:  user.Last_Name,
		uid:        user.User_ID,
		user_type:  user.User_Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &signeddetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret_key))
	refresh_token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secret_key))

	if err != nil {
		log.Panic(err)
		return "", ""
	}
	return token, refresh_token
}

func hashPassword(pwd string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func verifyPassword(pwd, current_pwd string) (bool, string) {
	check := true
	msg := ""

	if err := bcrypt.CompareHashAndPassword(
		[]byte(current_pwd),
		[]byte(pwd),
	); err != nil {
		check = false
		msg = "email or password incorrect!"
	}

	return check, msg
}

func updateAllTokens(token, refreshToken, uid string) {

}
