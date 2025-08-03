package images

import (
	// "fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DownloadImage(c *fiber.Ctx) error {
	return nil
}

func UploadImage(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("file eilna")
	}

	//file extension
	fileName := file.Filename
	ext := strings.ToLower(filepath.Ext(fileName))
	// fmt.Println(ext)

	//allowed file types
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}

	if !allowedExtensions[ext] {
		return c.Status(fiber.StatusBadRequest).JSON("invalid file type")
	}

	//create images dir
	saveDir := "./images"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON("unable to create directory")
	}

	//save the uploaded file
	savePath:=filepath.Join(saveDir,fileName)
	out,err:=os.Create(savePath)
	if err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON("internal server error")
	}
	defer out.Close()

	//saves the file
	return c.SaveFile(file, savePath)
}
