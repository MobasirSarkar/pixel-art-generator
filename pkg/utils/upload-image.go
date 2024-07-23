package imageurl

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(c *fiber.Ctx) error {
	log.Println("Recieved request:", c.Request().Header) // creating a file
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Error in uploading Image:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": 500, "message": "server error", "data": nil})
	}

	uniqueId := uuid.New() // generate a uuid

	// filename and extension function
	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	// save the image
	err = c.SaveFile(file, fmt.Sprintf("C:/users/abdul.mobasir/Desktop/pixel-art-generator/.images/%s.%s", filename, fileExt))
	if err != nil {
		log.Println("Error in saving Image:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "Error while saving image",
			"data":    nil,
		})
	}

	image := fmt.Sprintf("%s.%s", filename, fileExt)

	imageUrl := fmt.Sprintf("%s", image)

	data := map[string]interface{}{
		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    file.Header,
		"size":      file.Size,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "Image uploaded Successfully",
		"data":    data,
	})
}
