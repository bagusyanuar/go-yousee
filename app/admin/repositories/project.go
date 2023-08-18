package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ProjectRepository interface {
		GetData(name string, limit, offset int) ([]model.Project, error)
		GetDataByID(id string) (*model.Project, error)
		Count(name string) (int64, error)
		Create(entity model.Project) (*model.Project, error)
	}

	Project struct {
		database *gorm.DB
	}
)

// Count implements ProjectRepository.
func (r *Project) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.Project{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements ProjectRepository.
func (r *Project) Create(entity model.Project) (*model.Project, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements ProjectRepository.
func (r *Project) GetData(name string, limit int, offset int) ([]model.Project, error) {
	n := "%" + name + "%"
	var data []model.Project
	if err := r.database.
		Where("name LIKE ?", n).
		Preload("Items").
		Offset(offset).
		Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements ProjectRepository.
func (r *Project) GetDataByID(id string) (*model.Project, error) {
	data := new(model.Project)
	if err := r.database.Where("id = ?", id).
		Preload("Items").
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProject(db *gorm.DB) ProjectRepository {
	return &Project{
		database: db,
	}
}
