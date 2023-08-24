package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ItemImageTableName = "item_images"
)

type ItemImage struct {
	ID     uuid.UUID `json:"id"`
	ItemID uuid.UUID `json:"item_id"`
	Type   uint8     `json:"type"`
	Image  string    `json:"image"`
	common.WithTimestampsModel
	Item *Item `gorm:"foreignKey:ItemID" json:"item,omitempty"`
}

func (itemImage *ItemImage) BeforeCreate(tx *gorm.DB) (err error) {
	itemImage.ID = uuid.New()
	itemImage.CreatedAt = time.Now()
	itemImage.UpdatedAt = time.Now()
	return
}

func (ItemImage) TableName() string {
	return ItemImageTableName
}
