package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"Backend/models"
	"Backend/config"
	"context"
)

// Get all blogs
func GetBlogs(c *gin.Context) {
	var blogs []models.Blog
	cursor, err := config.Client.Database("your_db").Collection("blogs").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching blogs"})
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &blogs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching blogs"})
		return
	}
	c.JSON(http.StatusOK, blogs)
}

// Get a specific blog
func GetBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	err := config.Client.Database("your_db").Collection("blogs").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(http.StatusOK, blog)
}

// Create a new blog
func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := config.Client.Database("your_db").Collection("blogs").InsertOne(context.TODO(), blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create blog"})
		return
	}
	c.JSON(http.StatusCreated, blog)
}

// Comment on a blog
func CommentOnBlog(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	_, err := config.Client.Database("your_db").Collection("blogs").UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.M{"$push": bson.M{"comments": comment}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not comment on blog"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment added"})
}
