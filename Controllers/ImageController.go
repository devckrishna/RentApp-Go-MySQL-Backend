package controllers

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
)

func AddImage(c *gin.Context) {
	var newImage models.Image
	err := c.BindJSON(&newImage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	query := "INSERT INTO image (property_id, url) VALUES(?, ?)"
	_, err = db.Exec(query, newImage.Property_id, newImage.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Pass"})
}

func GetPropertyImages(c *gin.Context) {
	propertyId := c.Param("id")
	propertyId = strings.ReplaceAll(propertyId, "/", "")
	propertyIdInt, err := strconv.Atoi(propertyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	images := []models.Image{}
	query := "SELECT * FROM image WHERE property_id = ?"
	res, err := db.Query(query, propertyIdInt)

	for res.Next() {
		var image models.Image
		err := res.Scan(&image.Id, &image.Property_id, &image.Url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		images = append(images, image)
	}
	c.JSON(http.StatusOK, gin.H{"images": images})
}
