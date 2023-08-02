package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	GetData(name string, limit, offset int) ([]model.Province, error)
	GetDataByID(id string) (*model.Province, error)
	Count(name string) (int64, error)
}

type Province struct {
	database *gorm.DB
}

// Count implements ProvinceRepository.
func (r *Province) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.Province{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// GetData implements ProvinceRepository.
func (r *Province) GetData(name string, limit int, offset int) ([]model.Province, error) {
	n := "%" + name + "%"
	var data []model.Province
	if err := r.database.
		Where("name LIKE ?", n).Offset(offset).Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements ProvinceRepository.
func (r *Province) GetDataByID(id string) (*model.Province, error) {
	data := new(model.Province)
	if err := r.database.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewProvince(db *gorm.DB) ProvinceRepository {
	return &Province{
		database: db,
	}
}
