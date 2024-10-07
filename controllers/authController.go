package controllers

import (
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"Backend/models"
	"Backend/config"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte("your_secret_key") // Change this to a strong secret

// Register a new user
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)
	_, err := config.Client.Database("your_db").Collection("users").InsertOne(context.TODO(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

// Login user
func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := config.Client.Database("your_db").Collection("users").FindOne(context.TODO(), bson.M{"username": input.Username}).Decode(&user)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Get user profile
func Profile(c *gin.Context) {
	userID := c.MustGet("userID").(string)
	var user models.User
	err := config.Client.Database("your_db").Collection("users").FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
