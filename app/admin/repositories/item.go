package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ItemRepository interface {
		GetData(name string, limit, offset int) ([]model.Item, error)
		GetDataByID(id string) (*model.Item, error)
		Count(name string) (int64, error)
		Create(entity model.Item) (*model.Item, error)
	}

	Item struct {
		database *gorm.DB
	}
)

// Count implements ItemRepository.
func (r *Item) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.Item{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements ItemRepository.
func (r *Item) Create(entity model.Item) (*model.Item, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements ItemRepository.
func (r *Item) GetData(name string, limit int, offset int) ([]model.Item, error) {
	n := "%" + name + "%"
	var data []model.Item
	if err := r.database.
		Where("name LIKE ?", n).
		Preload("City").
		Preload("Vendor").
		Preload("Type").
		Offset(offset).
		Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements ItemRepository.
func (r *Item) GetDataByID(id string) (*model.Item, error) {
	data := new(model.Item)
	if err := r.database.Where("id = ?", id).
		Preload("City").
		Preload("Vendor").
		Preload("Type").
		First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewItem(db *gorm.DB) ItemRepository {
	return &Item{
		database: db,
	}
}
