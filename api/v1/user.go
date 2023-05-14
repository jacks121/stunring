package v1

import (
	"errors"
	"net/http"
	"swetelove/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserService(c *gin.Context) (*services.UserService, error) {
	db, ok := c.Get("db")
	if !ok {
		return nil, errors.New("database not found in context")
	}

	userService, err := services.NewUserService(db.(*gorm.DB))
	if err != nil {
		return nil, err
	}

	return userService, nil
}

func Register(c *gin.Context) {
	var registerInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&registerInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userService, err := getUserService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user service"})
		return
	}

	user, err := userService.Register(registerInfo.Username, registerInfo.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "user": user})
}

func Login(c *gin.Context) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userService, err := getUserService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user service"})
		return
	}

	token, err := userService.Login(loginInfo.Username, loginInfo.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func GetUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		return
	}

	userService, err := getUserService(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user service"})
		return
	}

	user, err := userService.GetUserByToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User fetched successfully", "user": user})
}
