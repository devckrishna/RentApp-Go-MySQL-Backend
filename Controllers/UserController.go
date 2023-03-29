package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "github.com/devckrishna/rentapp/Database"
	models "github.com/devckrishna/rentapp/Models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB = database.GetDB()

type LoginRequest struct {
	Email    string
	Password string
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}

	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("email of password is incorrect")
		check = false
	}
	return check, msg
}

func GetUsers(c *gin.Context) {
	query := "SELECT * FROM users"
	res, err := db.Query(query)

	defer res.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	users := []models.User{}

	for res.Next() {
		var user models.User
		err := res.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Age, &user.Gender, &user.Marital_status, &user.Photo, &user.Is_host)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		users = append(users, user)
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	userId := c.Param("id")
	userId = strings.ReplaceAll(userId, "/", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	var user models.User
	query := "SELECT * FROM users WHERE id = ?"
	err = db.QueryRow(query, userIdInt).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Age, &user.Gender, &user.Marital_status, &user.Photo, &user.Is_host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, user)
}

func SignUp(c *gin.Context) {
	var newUser models.User
	err := c.BindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	newUser.Password = HashPassword(newUser.Password)
	query := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"
	res, err := db.Exec(query, newUser.Name, newUser.Email, newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	newUser.Id, err = res.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func LoginUser(c *gin.Context) {
	var request LoginRequest
	var user models.User
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	query := "SELECT * FROM users WHERE email = ?"
	err = db.QueryRow(query, request.Email).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Age, &user.Gender, &user.Marital_status, &user.Photo, &user.Is_host)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// fmt.Println(user.Password)
	// fmt.Println(HashPassword(request.Password))
	isVerified, _ := VerifyPassword(request.Password, user.Password)

	if isVerified {
		c.JSON(http.StatusOK, user)
	} else {
		c.Status(http.StatusNotFound)
	}
}

func EnableHost(c *gin.Context) {
	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	query := "UPDATE users SET is_host = true WHERE id = ?"
	_, err = db.Exec(query, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "User enabled as host")
}

func DisableHost(c *gin.Context) {
	userId := c.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	query := "UPDATE users SET is_host = false WHERE id = ?"
	_, err = db.Exec(query, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "User disabled as host")
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")

	userId = strings.ReplaceAll(userId, "/", "")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	query := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(query, userIdInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, "Deleted user")
}
