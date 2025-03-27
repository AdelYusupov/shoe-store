// handlers/auth.go
package handlers

import (
	"net/http"
	"time"

	"shoe-store-api/database"
	"shoe-store-api/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var jwtKey = []byte("6cfd7525324e3fc4baf2782319827555b71e49ce94cefd0dc83b68819f5e2680a09bf0080a6605c6d2d214c156c8af4485fef9c958720924de435db7aeac87b14976b1ab8cb4bf1110eb298ed9ccf9dc40f02ae2dd2bf0d505db8e622f34eabe1c1981292444d28c4dadd25b80320fc76e84ed6e8acb0b74769308fca34c3ae0ff09b452320def56232bac0bc79f61ec563db0fdc752511ec7f6db4f7b2c2c453921dc322e5660215a4d5dffd3b6cab640a9ab63558fec338bfc39a7480941c50ebfbab5b05a3b567a0402773cc1b2da4e75d5f245d82054ac86e6db7416bc64cdc5ed0b13d1ae8b9fe8684ce0df547c59bc4255b815cc8ad46fba74f940d64d") // Замените на реальный секретный ключ

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.AdminUser
	if err := database.DB.Where("username = ?", req.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.PasswordHash), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   admin.Username,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"expires": expirationTime,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		claims := &jwt.StandardClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		c.Set("username", claims.Subject)
		c.Next()
	}
}
