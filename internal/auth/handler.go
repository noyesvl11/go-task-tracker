package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-project/internal/db"
	user "rest-project/internal/models"
)

func Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid"})
		return
	}

	var u user.User
	db.DB.Where("username = ? ", req.Username).First(&u)
	if u.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login or password incorrect"})
		return
	}

	token, _ := GenerateJWT(u.ID, u.Role)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     string `json:"role"` // admin, teacher, student
	}

	if err := c.ShouldBindJSON(&req); err != nil || req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.Role == "" {
		req.Role = "student" // по умолчанию
	} else if req.Role != "admin" && req.Role != "teacher" && req.Role != "student" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	// Проверка, существует ли пользователь
	var existing user.User
	if err := db.DB.Where("username = ?", req.Username).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already taken"})
		return
	}

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Создание нового пользователя
	u := user.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Role:     req.Role,
	}
	if err := db.DB.Create(&u).Error; err != nil {
		// Логируем ошибку для более подробной информации
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
