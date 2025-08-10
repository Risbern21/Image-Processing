package images

import (
	"fmt"
	"image/models/files"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
		return c.Status(fiber.StatusBadRequest).JSON(files.Errors{
			Error: "invalid file type",
		})
	}

	//create images dir
	saveDir := "./images"
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "unable to create directory",
		})
	}

	//save the uploaded file
	savePath := filepath.Join(saveDir, fileName)
	out, err := os.Create(savePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "internal server error",
		})
	}
	defer out.Close()

	//saves the file
	c.SaveFile(file, savePath)

	//Detect MIME type
	savedFile, err := os.Open(savePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "unable to open file for reading",
		})
	}
	defer savedFile.Close()

	buffer := make([]byte, 512)
	if _, err := savedFile.Read(buffer); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "Unable to read file",
		})
	}
	contentType := http.DetectContentType(buffer)

	//Get file info
	fileInfo, err := savedFile.Stat()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{})
	}

	i := files.NewImageProcessingSchemaImage()
	imageId, err := i.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "internal server error",
		})
	}

	iO := files.NewImageOptions()
	iO.ImageId = imageId
	if err := iO.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "internal server error",
		})
	}

	//Create imgae url
	imageUrl := fmt.Sprintf("%s://%s/%s", c.Request().URI().Scheme(), c.Request().Host(), fileName)

	return c.Status(fiber.StatusOK).JSON(files.Image{
		Message: "image successfully saved",
		Url:     imageUrl,
		MetaData: files.MetaData{
			Name:        fileName,
			SizeBytes:   fileInfo.Size(),
			ContentType: contentType,
		},
	})
}

type TransformationPayload struct {
	Resize *struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"resize,omitempty"`

	Crop *struct {
		Width  int `json:"width"`
		Height int `json:"height"`
		X      int `json:"x"`
		Y      int `json:"y"`
	} `json:"crop,omitempty"`

	Rotate *int `json:"rotate,omitempty"`
	Format *int `json:"format,omitempty"`

	Filter *struct {
		GrayScale *bool `json:"grayscale,omitempty"`
		Sepia     *bool `json:"sepia,omitempty"`
	} `json:"filter,omitempty"`
}

type imageTransformationsReq struct {
	Transformations TransformationPayload `json:"transformations"`
}

func TransformImage(c *fiber.Ctx) error {
	var req imageTransformationsReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(files.Errors{
			Error: "invalid image options",
		})
	}

	iO := files.NewImageOptions()
	//iO.ID=someid
	if err := iO.Get(); err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON("image not found")
		}
		return c.Status(fiber.StatusInternalServerError).JSON(files.Errors{
			Error: "internal server error",
		})
	}

	return c.Status(fiber.StatusOK).JSON("image trasnformed")
}

func DownloadImage(c *fiber.Ctx) error {
	fileName := c.Params("filename")
	filePath := "./images/" + fileName

	c.Response().Header.Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	return c.SendFile(filePath)
}
