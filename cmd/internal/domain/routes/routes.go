package routes

import (
	"errors"
	"log"

	"github.com/LeonardoGrigolettoDev/go-gypsy.git/cmd/internal/services/loader"
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

	router.POST("/data/load/:file", func(c *gin.Context) {
		typeLoad := c.Param("file")
		perm := c.Query("permission")

		if len((typeLoad)) == 0 {
			c.JSON(400, gin.H{
				"error": "Type load is invalid.",
			})
			return
		}
		if perm != "public" && perm != "private" {
			if len(perm) == 0 {
				c.JSON(400, gin.H{
					"error": "Type permission is not in query params.",
				})
				return
			}
			c.JSON(400, gin.H{
				"error": "Type permission is invalid.",
			})
			return
		}
		switch typeLoad {
		case "csv":
			{
				file, err := c.FormFile("file")
				if err != nil {
					log.Println("Erro ao obter arquivo:", err)
					c.JSON(400, gin.H{
						"error": "Failed to get file from request.",
					})
					return
				}
				log.Printf("Arquivo recebido: %s (tamanho: %d bytes)\n", file.Filename, file.Size)
				err = loader.CreateFile(*file, "./files/temp/"+perm+"/"+file.Filename)
				if err != nil {
					log.Println(err)
					c.JSON(500, gin.H{
						"error": "Could not create file in current path.",
					})
					return
				}
				c.JSON(201, gin.H{
					"error": "Created new data (not implemented).",
				})
				return
			}
		default:
			{
				c.JSON(501, gin.H{
					"error": "Type of load is not implemented.",
				})
			}
		}
	})

	return nil

}
