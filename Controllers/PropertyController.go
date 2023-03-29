package controllers

import (
	"net/http"
	"strconv"
	"strings"

	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
)

func CreateProperty(c *gin.Context) {
	var newProperty models.Property
	err := c.BindJSON(&newProperty)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	query := "INSERT INTO property (title, owner_id, city, country, total_rooms, total_area, rating, nei_details,price, avy_living_cost, facilities) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, newProperty.Title, newProperty.Owner_id, newProperty.City, newProperty.Country, newProperty.Total_rooms, newProperty.Total_area, newProperty.Rating, newProperty.Nei_details, newProperty.Price, newProperty.Avg_living_cost, newProperty.Facilities)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Pass"})
}

func GetAllProperties(c *gin.Context) {
	query := "SELECT * FROM property"
	res, err := db.Query(query)

	defer res.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	properties := []models.Property{}

	for res.Next() {
		var property models.Property
		err := res.Scan(&property.Id, &property.Title, &property.Owner_id, &property.City, &property.Country, &property.Total_rooms, &property.Total_area, &property.Rating, &property.Nei_details, &property.Price, &property.Avg_living_cost, &property.Facilities)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		properties = append(properties, property)
	}
	c.JSON(http.StatusOK, gin.H{"properties": properties})
}

func GetPropertyById(c *gin.Context) {
	propertyId := c.Param("id")
	propertyId = strings.ReplaceAll(propertyId, "/", "")
	propertyIdInt, err := strconv.Atoi(propertyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var property models.Property
	query := "SELECT * FROM property WHERE id = ?"
	err = db.QueryRow(query, propertyIdInt).Scan(&property.Id, &property.Title, &property.Owner_id, &property.City, &property.Country, &property.Total_rooms, &property.Total_area, &property.Rating, &property.Nei_details, &property.Price, &property.Avg_living_cost, &property.Facilities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"property": property})
}

func GetTopRatedProperties(c *gin.Context) {
	query := "SELECT * FROM property WHERE RATING >= 4.5"
	res, err := db.Query(query)

	defer res.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	properties := []models.Property{}

	for res.Next() {
		var property models.Property
		err := res.Scan(&property.Id, &property.Title, &property.Owner_id, &property.City, &property.Country, &property.Total_rooms, &property.Total_area, &property.Rating, &property.Nei_details, &property.Price, &property.Avg_living_cost, &property.Facilities)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		properties = append(properties, property)
	}
	c.JSON(http.StatusOK, gin.H{"properties": properties})
}

func GetPropertiesNearYou(c *gin.Context) {
	cityName := c.Param("city")
	cityName = strings.ReplaceAll(cityName, "/", "")

	var property models.Property
	query := `SELECT * FROM property WHERE City = ?`
	err := db.QueryRow(query, cityName).Scan(&property.Id, &property.Title, &property.Owner_id, &property.City, &property.Country, &property.Total_rooms, &property.Total_area, &property.Rating, &property.Nei_details, &property.Price, &property.Avg_living_cost, &property.Facilities)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"property": property})
}
