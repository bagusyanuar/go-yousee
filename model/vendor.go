package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	VendorTableName = "vendors"
)

type Vendor struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Phone   string    `json:"phone"`
	Brand   string    `json:"brand"`
	common.WithTimestampsModel
}

func (vendor *Vendor) BeforeCreate(tx *gorm.DB) (err error) {
	vendor.ID = uuid.New()
	vendor.CreatedAt = time.Now()
	vendor.UpdatedAt = time.Now()
	return
}

func (Vendor) TableName() string {
	return VendorTableName
}
