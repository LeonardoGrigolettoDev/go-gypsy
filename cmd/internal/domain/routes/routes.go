package routes

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configura as rotas do servidor.
func SetupRoutes(router *gin.Engine) error {
	if router == nil {
		return errors.New("o roteador é nulo")
	}

	// Rota de ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	return nil
}
