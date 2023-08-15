package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	MediaTypeRepository interface {
		GetData(name string, limit, offset int) ([]model.MediaType, error)
		GetDataByID(id string) (*model.MediaType, error)
		Count(name string) (int64, error)
		Create(entity model.MediaType) (*model.MediaType, error)
		Patch(id string, entity model.MediaType) (*model.MediaType, error)
		Delete(id string) error
	}

	MediaType struct {
		database *gorm.DB
	}
)

// Patch implements MediaTypeRepository.
func (r *MediaType) Patch(id string, entity model.MediaType) (*model.MediaType, error) {
	if err := r.database.Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Delete implements MediaTypeRepository.
func (r *MediaType) Delete(id string) error {
	if err := r.database.Where("id = ?", id).Delete(&model.MediaType{}).Error; err != nil {
		return err
	}
	return nil
}

// Count implements TypeRepository.
func (r *MediaType) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.MediaType{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements TypeRepository.
func (r *MediaType) Create(entity model.MediaType) (*model.MediaType, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements TypeRepository.
func (r *MediaType) GetData(name string, limit int, offset int) ([]model.MediaType, error) {
	n := "%" + name + "%"
	var data []model.MediaType
	if err := r.database.
		Where("name LIKE ?", n).Offset(offset).Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements TypeRepository.
func (r *MediaType) GetDataByID(id string) (*model.MediaType, error) {
	data := new(model.MediaType)
	if err := r.database.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewMediaType(db *gorm.DB) MediaTypeRepository {
	return &MediaType{
		database: db,
	}
}
