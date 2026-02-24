package controllers

import (
	"net/http"
	"usersList/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User inserted successfully!"})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.UpdateUser(userId, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully!"})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	err := models.DeleteUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully!"})
}

func DeleteAllUsers(c *gin.Context) {
	err := models.DeleteAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete all users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "All users deleted successfully!"})
}

func ListAllUsers(c *gin.Context) {
	users := models.ListAll()
	c.JSON(http.StatusOK, users)
}

func FindUserByName(c *gin.Context) {
	username := c.Param("name")
	user := models.FindUser(username)
	if user.Name == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func FindAllUsersByName(c *gin.Context) {
	username := c.Param("name")
	users := models.FindAllUsers(username)
	if len(users) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found"})
		return
	}
	c.JSON(http.StatusOK, users)
}

func InsertMultipleUsers(c *gin.Context) {
	var users []models.User
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.InsertManyUsers(users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert users"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "users inserted successfully!"})
}
