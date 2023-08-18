package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ProjectItemTableName = "project_items"
)

type ProjectItem struct {
	ID          uuid.UUID `json:"id"`
	ProjectID   uuid.UUID `json:"project_id"`
	CityID      uuid.UUID `json:"city_id"`
	ItemID      uuid.UUID `json:"item_id"`
	PicID       uuid.UUID `json:"pic_id"`
	VendorPrice int64     `json:"vendor_price"`
	common.WithTimestampsModel
	Project *Project `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	City    *City    `gorm:"foreignKey:CityID" json:"city,omitempty"`
	Item    *Item    `gorm:"foreignKey:ItemID" json:"item,omitempty"`
	Pic     *User    `gorm:"foreignKey:PicID" json:"pic,omitempty"`
}

func (projectItem *ProjectItem) BeforeCreate(tx *gorm.DB) (err error) {
	projectItem.ID = uuid.New()
	projectItem.CreatedAt = time.Now()
	projectItem.UpdatedAt = time.Now()
	return
}

func (ProjectItem) TableName() string {
	return ProjectItemTableName
}
