package main

import (
	"log"

	"github.com/badimalex/goshop/config"
	"github.com/badimalex/goshop/internal/handlers"
	"github.com/badimalex/goshop/internal/repositories"
	"github.com/gin-gonic/gin"

	"github.com/badimalex/goshop/internal/database"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productRepo)

	jwtSecret := cfg.JwtSecret
	userRepo := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo, jwtSecret)

	r := gin.Default()

	r.GET("/products", productHandler.SearchProducts)
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products/:id", productHandler.GetProduct)
	r.PUT("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	address := cfg.Server.Host + ":" + cfg.Server.Port

	log.Printf("Starting server on %s", address)
	log.Fatal(r.Run(address))
}
