package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	CityTableName = "cities"
)

type City struct {
	ID         uuid.UUID `json:"id"`
	ProvinceID uuid.UUID `json:"province_id"`
	Code       string    `json:"code"`
	Name       string    `json:"name"`
	common.WithTimestampsModel
	Province Province `gorm:"foreignKey:ProvinceID" json:"province"`
}

func (city *City) BeforeCreate(tx *gorm.DB) (err error) {
	city.ID = uuid.New()
	city.CreatedAt = time.Now()
	city.UpdatedAt = time.Now()
	return
}

func (City) TableName() string {
	return CityTableName
}
