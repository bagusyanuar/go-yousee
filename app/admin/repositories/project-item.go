package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ProjectItemRepository interface {
		GetData(name string, limit, offset int) ([]model.ProjectItem, error)
		GetDataByID(id string) (*model.ProjectItem, error)
		Count(name string) (int64, error)
		Create(entity model.ProjectItem) (*model.ProjectItem, error)
	}

	ProjectItem struct {
		database *gorm.DB
	}
)

// Count implements ProjectItemRepository.
func (r *ProjectItem) Count(name string) (int64, error) {
	var totalRows int64
	if err := r.database.
		Model(&model.ProjectItem{}).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements ProjectItemRepository.
func (r *ProjectItem) Create(entity model.ProjectItem) (*model.ProjectItem, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements ProjectItemRepository.
func (r *ProjectItem) GetData(name string, limit int, offset int) ([]model.ProjectItem, error) {
	var data []model.ProjectItem
	if err := r.database.
		Preload("Project").
		Preload("City").
		Preload("Item").
		Preload("Pic").
		Offset(offset).
		Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements ProjectItemRepository.
func (r *ProjectItem) GetDataByID(id string) (*model.ProjectItem, error) {
	data := new(model.ProjectItem)
	if err := r.database.Where("id = ?", id).
		Preload("Project").
		Preload("City").
		Preload("Item").
		Preload("Pic").
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProjectItem(db *gorm.DB) ProjectItemRepository {
	return &ProjectItem{
		database: db,
	}
}
