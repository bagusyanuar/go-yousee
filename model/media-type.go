package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	MediaTypeTableName = "media_types"
)

type MediaType struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Slug string    `json:"slug"`
	Icon *string   `json:"icon"`
	common.WithTimestampsModel
}

func (mt *MediaType) BeforeCreate(tx *gorm.DB) (err error) {
	mt.ID = uuid.New()
	mt.CreatedAt = time.Now()
	mt.UpdatedAt = time.Now()
	return
}

func (MediaType) TableName() string {
	return MediaTypeTableName
}
