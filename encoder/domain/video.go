package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"column:video_id;type:uuid;not null;primarykey;default:uuid_generate_v4()"`
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255);not null"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"foreignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo() *Video {
	return &Video{
		CreatedAt: time.Now(),
	}
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	if err != nil {
		return err
	}
	return nil
}
