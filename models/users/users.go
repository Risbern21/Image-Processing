package users

import (
	"image/internal/database"
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null;default:null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New() *Users {
	return &Users{}
}

func (u *Users) Create() error {
	if err := database.Client().Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *Users) Get() error {
	if err := database.Client().First(u, u.ID).Error; err != nil {
		return err
	}
	return nil
}

func (u *Users) Update() error {
	if err := database.Client().Save(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *Users) Delete() error {
	if err := database.Client().Delete(u).Error; err != nil {
		return err
	}
	return nil
}
