package controllers

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
)

func CreateRequest(c *gin.Context) {
	var newRequest models.Request
	err := c.BindJSON(&newRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	query := "INSERT INTO request (user_id, property_id) VALUES(?, ?)"
	_, err = db.Exec(query, newRequest.User_id, newRequest.Property_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Pass"})
}

func GetRequestByPropertyId(c *gin.Context) {
	propertyId := c.Param("id")
	propertyId = strings.ReplaceAll(propertyId, "/", "")
	propertyIdInt, err := strconv.Atoi(propertyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	requests := []models.Request{}
	query := "SELECT * FROM request WHERE property_id = ?"
	res, err := db.Query(query, propertyIdInt)

	for res.Next() {
		var request models.Request
		err := res.Scan(&request.Id, &request.User_id, &request.Property_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		requests = append(requests, request)
	}
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func GetAllRequestsByUser(c *gin.Context) {
	userId := c.Param("id")
	userId = strings.ReplaceAll(userId, "/", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	requests := []models.Request{}
	query := "SELECT * FROM request WHERE user_id = ?"
	res, err := db.Query(query, userIdInt)

	for res.Next() {
		var request models.Request
		err := res.Scan(&request.Id, &request.User_id, &request.Property_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		requests = append(requests, request)
	}
	c.JSON(http.StatusOK, gin.H{"requests": requests})
}
