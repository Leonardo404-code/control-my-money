package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	os.Setenv("PORT", "3000")

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
