package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ItemTableName = "items"
)

type Item struct {
	ID          uuid.UUID `json:"id"`
	CityID      uuid.UUID `json:"city_id"`
	MediaTypeID uuid.UUID `json:"media_type_id"`
	VendorID    uuid.UUID `json:"vendor_id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Position    uint8     `json:"position"`
	Width       float64   `json:"width"`
	Height      float64   `json:"height"`
	common.WithTimestampsModel
	City      *City       `gorm:"foreignKey:CityID" json:"city,omitempty"`
	MediaType *MediaType  `gorm:"foreignKey:MediaTypeID" json:"media_type,omitempty"`
	Vendor    *Vendor     `gorm:"foreignKey:VendorID" json:"vendor,omitempty"`
	Images    []ItemImage `gorm:"foreignKey:ItemID" json:"images,omitempty"`
}

func (item *Item) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	return
}

func (Item) TableName() string {
	return ItemTableName
}
