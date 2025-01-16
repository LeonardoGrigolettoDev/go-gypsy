package main

import (
	"log"
	"os"

	"github.com/LeonardoGrigolettoDev/go-gypsy.git/cmd/internal/config"
	"github.com/LeonardoGrigolettoDev/go-gypsy.git/cmd/internal/domain/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Configurações de ambiente
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	// Conexão com MongoDB
	if err := config.ConnectMongoDB(mongoURI); err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}
	defer func() {
		if err := config.MongoClient.Disconnect(nil); err != nil {
			log.Printf("Erro ao desconectar do MongoDB: %v", err)
		}
	}()

	// Configuração do Gin
	router := gin.Default()

	// Configura rotas
	if err := routes.SetupRoutes(router); err != nil {
		log.Fatalf("Erro ao configurar rotas: %v", err)
	}

	// Iniciar servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
