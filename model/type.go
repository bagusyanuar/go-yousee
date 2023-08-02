package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	TypeTableName = "types"
)

type Type struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Icon *string   `json:"icon"`
	common.WithTimestampsModel
}

func (t *Type) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}

func (Type) TableName() string {
	return TypeTableName
}
