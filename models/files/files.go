package files

import (
	"image/internal/database"
	"time"
)

type MetaData struct {
	Name        string `json:"name"`
	SizeBytes   int64  `json:"size_bytes"`
	ContentType string `json:"content_type"`
}

type Image struct {
	Message  string   `json:"message"`
	Url      string   `json:"url"`
	MetaData MetaData `json:"metadata"`
}

type Errors struct {
	Error string `json:"error"`
}

type ImageProcessingSchemaImage struct {
	ID        int32     `gorm:"primaryKey;autoIncrement" json:"id"`
	Url       string    `gorm:"not null;default=null" json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ImageOptions struct {
	ID           int32  `json:"id"`
	ResizeWidth  int32  `json:"resize_width"`
	ResizeHeight int32  `json:"resize_height"`
	CropWidth    int32  `json:"crop_width"`
	CropHeight   int32  `json:"crop_height"`
	CropX        int32  `json:"crop_x"`
	CropY        int32  `json:"crop_y"`
	Rotate       int32  `json:"rotate"`
	Format       string `json:"format"`
	GrayScale    bool   `json:"grayscale"`
	Sepia        bool   `json:"sepia"`
	ImageId      int32  `json:"image_id"`
}

func NewImageProcessingSchemaImage() *ImageProcessingSchemaImage {
	return &ImageProcessingSchemaImage{}
}

func (i *ImageProcessingSchemaImage) Create() (int32, error) {
	if err := database.Client().Create(&i).Error; err != nil {
		return 0, err
	}

	return i.ID, nil
}

func (i *ImageProcessingSchemaImage) Get() error {
	if err := database.Client().First(i, i.ID).Error; err != nil {
		return err
	}
	return nil
}

func (i *ImageProcessingSchemaImage) Update() error {
	if err := database.Client().Save(i).Error; err != nil {
		return err
	}
	return nil
}

func (i *ImageProcessingSchemaImage) Delete() error {
	if err := database.Client().Delete(i, i.ID).Error; err != nil {
		return err
	}
	return nil
}

func NewImageOptions() *ImageOptions {
	return &ImageOptions{}
}

func (i *ImageOptions) Create() error {
	if err := database.Client().Create(&i).Error; err != nil {
		return err
	}

	return nil
}

func (i *ImageOptions)Get()error{
	if err:=database.Client().First(i,i.ID).Error;err!=nil{
		return err
	}
	return nil
}