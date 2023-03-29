package controllers

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
)

func AddToSaved(c *gin.Context) {
	var newSaved models.Saved
	err := c.BindJSON(&newSaved)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	query := "INSERT INTO saved (user_id, property_id) VALUES(?, ?)"
	_, err = db.Exec(query, newSaved.User_id, newSaved.Property_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Pass"})
}

func GetSaved(c *gin.Context) {
	userId := c.Param("id")
	userId = strings.ReplaceAll(userId, "/", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	savedList := []models.Saved{}
	query := "SELECT * FROM saved WHERE user_id = ?"
	res, err := db.Query(query, userIdInt)

	for res.Next() {
		var saved models.Saved
		err := res.Scan(&saved.Id, &saved.User_id, &saved.Property_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		savedList = append(savedList, saved)
	}
	c.JSON(http.StatusOK, gin.H{"savedList": savedList})
}
