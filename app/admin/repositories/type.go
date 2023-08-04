package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	TypeRepository interface {
		GetData(name string, limit, offset int) ([]model.Type, error)
		GetDataByID(id string) (*model.Type, error)
		Count(name string) (int64, error)
		Create(entity model.Type) (*model.Type, error)
	}

	Type struct {
		database *gorm.DB
	}
)

// Count implements TypeRepository.
func (r *Type) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.Type{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements TypeRepository.
func (r *Type) Create(entity model.Type) (*model.Type, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements TypeRepository.
func (r *Type) GetData(name string, limit int, offset int) ([]model.Type, error) {
	n := "%" + name + "%"
	var data []model.Type
	if err := r.database.
		Where("name LIKE ?", n).Offset(offset).Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements TypeRepository.
func (r *Type) GetDataByID(id string) (*model.Type, error) {
	data := new(model.Type)
	if err := r.database.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewType(db *gorm.DB) TypeRepository {
	return &Type{
		database: db,
	}
}
