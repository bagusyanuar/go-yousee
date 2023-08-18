package model

import (
	"time"

	"github.com/bagusyanuar/go-yousee/common"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const (
	ProjectTableName = "projects"
)

type Project struct {
	ID           uuid.UUID           `json:"id"`
	Name         string              `json:"name"`
	ClientName   string              `json:"client_name"`
	RequestDate  datatypes.Date      `json:"request_date"`
	Qty          uint                `json:"qty"`
	Description  string              `json:"description"`
	Duration     uint                `json:"duration"`
	DurationUnit common.DurationUnit `json:"duration_unit"`
	IsLightOn    bool                `json:"is_light_on"`
	Status       uint8               `json:"status"`
	common.WithTimestampsModel
	Items []*ProjectItem `gorm:"foreignKey:ProjectID" json:"items"`
}

func (project *Project) BeforeCreate(tx *gorm.DB) (err error) {
	project.ID = uuid.New()
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()
	return
}

func (Project) TableName() string {
	return ProjectTableName
}
