package handlers

import (
	"net/http"

	"shoe-store-api/database"
	"shoe-store-api/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

type OrderRequest struct {
	FullName    string             `json:"full_name"`
	Phone       string             `json:"phone"`
	City        string             `json:"city"`
	Address     string             `json:"address"`
	Delivery    string             `json:"delivery"`
	PickupPoint string             `json:"pickup_point"`
	Comment     string             `json:"comment"`
	Contact     string             `json:"contact"`
	Items       []models.OrderItem `json:"items"`
	Total       float64            `json:"total"`
}

func CreateOrder(c *gin.Context) {
	var request OrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		FullName:    request.FullName,
		Phone:       request.Phone,
		City:        request.City,
		Address:     request.Address,
		Delivery:    request.Delivery,
		PickupPoint: request.PickupPoint,
		Comment:     request.Comment,
		Contact:     request.Contact,
		Items:       request.Items,
		Total:       request.Total,
		Status:      "new",
	}

	if err := database.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Order created successfully",
		"order_id": order.ID,
	})
}
