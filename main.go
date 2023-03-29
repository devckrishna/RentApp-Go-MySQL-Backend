package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	database "github.com/devckrishna/rentapp/Database"
	routes "github.com/devckrishna/rentapp/Routes"
)

var db *sql.DB

func main() {

	db = database.GetDB()

	server := gin.Default()
	server.Static("/images", "./images")
	// server.Use(helpers.CORSMiddleware())
	server.Use(gin.Logger())
	routes.UserRoutes(server)
	routes.PropertyRoutes(server)

	server.POST("/api/upload", imageUpload)

	server.GET("/api", func(ctx *gin.Context) {
		ctx.JSON(200, "hye!!!")
	})

	server.Run()
}

func imageUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	filename := header.Filename
	fmt.Println(header.Filename)
	var pathName = time.Now().String() + filename
	pathName = strings.ReplaceAll(pathName, " ", "")
	var name = "./images/" + pathName
	out, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{"url": "/images/" + pathName})
}
