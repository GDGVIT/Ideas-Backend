package main

import (
	"github.com/GDGVIT/Ideas-Backend/pkg/configs"
	"github.com/GDGVIT/Ideas-Backend/pkg/middleware"
	"github.com/GDGVIT/Ideas-Backend/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber Config
	config := configs.FiberConfig()

	// Define new app with Fiber config
	app := fiber.New(config)

	// use middleware
	middleware.FiberMiddleware(app)

	utils.StartServer(app)
}
