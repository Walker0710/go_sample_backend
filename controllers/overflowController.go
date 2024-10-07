package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"Backend/models"
	"Backend/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

// Get all overflows
func GetOverflows(c *gin.Context) {
	var overflows []models.Overflow
	cursor, err := config.Client.Database("your_db").Collection("overflows").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching overflows"})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &overflows); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching overflows"})
		return
	}
	c.JSON(http.StatusOK, overflows)
}

// Get a specific overflow
func GetOverflow(c *gin.Context) {
	id := c.Param("id")
	var overflow models.Overflow
	err := config.Client.Database("your_db").Collection("overflows").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&overflow)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Overflow not found"})
		return
	}
	c.JSON(http.StatusOK, overflow)
}

// Create a new overflow
func CreateOverflow(c *gin.Context) {
	var overflow models.Overflow
	if err := c.ShouldBindJSON(&overflow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := config.Client.Database("your_db").Collection("overflows").InsertOne(context.TODO(), overflow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create overflow"})
		return
	}
	c.JSON(http.StatusCreated, overflow)
}

// Comment on an overflow
func CommentOnOverflow(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	_, err := config.Client.Database("your_db").Collection("overflows").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$push": bson.M{"comments": comment}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not comment on overflow"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}

// Get user's overflows
func GetUserOverflows(c *gin.Context) {
	username := c.Param("username")
	var overflows []models.Overflow
	cursor, err := config.Client.Database("your_db").Collection("overflows").Find(context.TODO(), bson.M{"username": username})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user's overflows"})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &overflows); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user's overflows"})
		return
	}
	c.JSON(http.StatusOK, overflows)
}
