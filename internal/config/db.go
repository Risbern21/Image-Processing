package config

import (
	"image/internal/database"
	"image/models/files"
	"image/models/users"
)

func AutoMigrate() {
	database.Client().AutoMigrate(&users.Users{}, &files.ImageOptions{}, &files.ImageProcessingSchemaImage{})
}
