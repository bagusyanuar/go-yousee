package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ItemImageRepository interface {
		Create(entity model.ItemImage) (*model.ItemImage, error)
		GetDataByItemID(id string) ([]model.ItemImage, error)
	}

	ItemImage struct {
		database *gorm.DB
	}
)

// Create implements ItemImageRepository.
func (r *ItemImage) Create(entity model.ItemImage) (*model.ItemImage, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetDataByItemID implements ItemImageRepository.
func (r *ItemImage) GetDataByItemID(id string) ([]model.ItemImage, error) {
	var data []model.ItemImage
	if err := r.database.
		Where("item_id = ?", id).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func NewItemImage(db *gorm.DB) ItemImageRepository {
	return &ItemImage{
		database: db,
	}
}
