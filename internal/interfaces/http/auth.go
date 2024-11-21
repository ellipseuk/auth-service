package http

import (
	"auth-service/internal/entities"
	"auth-service/internal/usecases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	RegisterUseCase *usecases.RegisterUserUseCase
}

func (h *AuthHandler) Register(c *gin.Context) {
	var user entities.User

	fmt.Println("Incoming request to /auth/register")

	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// Print parsed user
	fmt.Printf("Parsed user: %+v\n", user)

	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
		return
	}

	if err := h.RegisterUseCase.Execute(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
