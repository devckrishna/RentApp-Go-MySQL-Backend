package controllers

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
)

func CreateReview(c *gin.Context) {
	var newReview models.Review
	err := c.BindJSON(&newReview)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	query := "INSERT INTO review (property_id, user_id, body, rating) VALUES(? ,?, ?, ?)"
	_, err = db.Exec(query, newReview.Property_id, newReview.User_id, newReview.Body, newReview.Rating)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Pass"})
}

func GetReviewByPropertyId(c *gin.Context) {
	propertyId := c.Param("id")
	propertyId = strings.ReplaceAll(propertyId, "/", "")
	propertyIdInt, err := strconv.Atoi(propertyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	reviews := []models.Review{}
	query := "SELECT * FROM review WHERE property_id = ?"
	res, err := db.Query(query, propertyIdInt)

	for res.Next() {
		var review models.Review
		err := res.Scan(&review.Id, &review.Property_id, &review.User_id, &review.Body, &review.Rating)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		reviews = append(reviews, review)
	}
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}
