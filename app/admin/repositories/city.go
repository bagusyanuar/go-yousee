package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	CityRepository interface {
		GetData(name string, limit, offset int) ([]model.City, error)
		GetDataByID(id string) (*model.City, error)
		Count(name string) (int64, error)
		Create(entity model.City) (*model.City, error)
	}

	City struct {
		database *gorm.DB
	}
)

// Count implements CityRepository.
func (r *City) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.City{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements CityRepository.
func (r *City) Create(entity model.City) (*model.City, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements CityRepository.
func (r *City) GetData(name string, limit int, offset int) ([]model.City, error) {
	n := "%" + name + "%"
	var data []model.City
	if err := r.database.
		Where("name LIKE ?", n).Offset(offset).Limit(limit).Preload("Province").
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements CityRepository.
func (r *City) GetDataByID(id string) (*model.City, error) {
	data := new(model.City)
	if err := r.database.Where("id = ?", id).Preload("Province").First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewCity(db *gorm.DB) CityRepository {
	return &City{
		database: db,
	}
}
