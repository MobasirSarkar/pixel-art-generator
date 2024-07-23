package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	imageurl "pixel-art/pkg/utils"
)

func main() {
	fmt.Print("pixel art generator")
	app := fiber.New()
	app.Post("/upload", imageurl.UploadImage)
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting server:%v", err)
	}
}
