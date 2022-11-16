package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var app = fiber.New()

	app.Listen(":3000")
	fmt.Println("server port 3000")
}
