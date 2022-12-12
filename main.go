package main

import (
	"log"

	"github.com/GDGVIT/Ideas-Backend/pkg/configs"
	"github.com/GDGVIT/Ideas-Backend/pkg/middleware"
	"github.com/GDGVIT/Ideas-Backend/pkg/utils"
	"github.com/GDGVIT/Ideas-Backend/platform/database"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber Config
	config := configs.FiberConfig()
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// Define new app with Fiber config
	app := fiber.New(config)
	// use middleware
	middleware.FiberMiddleware(app)

	database.PostgreSQLConnection()

	utils.StartServer(app)
}
