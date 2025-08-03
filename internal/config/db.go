package config

import (
	"image/internal/database"
	users "image/models"
)

func AutoMigrate() {
	database.Client().AutoMigrate(&users.Users{})
}
