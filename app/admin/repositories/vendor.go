package repositories

import (
	"github.com/bagusyanuar/go-yousee/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	VendorRepository interface {
		GetData(name string, limit, offset int) ([]model.Vendor, error)
		GetDataByID(id string) (*model.Vendor, error)
		Count(name string) (int64, error)
		Create(entity model.Vendor) (*model.Vendor, error)
		Patch(id string, entity model.Vendor) (*model.Vendor, error)
		Delete(id string) error
	}

	Vendor struct {
		database *gorm.DB
	}
)

// Delete implements VendorRepository.
func (r *Vendor) Delete(id string) error {
	if err := r.database.Where("id = ?", id).Delete(&model.Vendor{}).Error; err != nil {
		return err
	}
	return nil
}

// Patch implements VendorRepository.
func (r *Vendor) Patch(id string, entity model.Vendor) (*model.Vendor, error) {
	if err := r.database.Omit(clause.Associations).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Count implements VendorRepository.
func (r *Vendor) Count(name string) (int64, error) {
	n := "%" + name + "%"
	var totalRows int64
	if err := r.database.
		Model(&model.Vendor{}).
		Where("name LIKE ?", n).
		Count(&totalRows).Error; err != nil {
		return 0, err
	}
	return totalRows, nil
}

// Create implements VendorRepository.
func (r *Vendor) Create(entity model.Vendor) (*model.Vendor, error) {
	if err := r.database.Omit(clause.Associations).Create(&entity).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// GetData implements VendorRepository.
func (r *Vendor) GetData(name string, limit int, offset int) ([]model.Vendor, error) {
	n := "%" + name + "%"
	var data []model.Vendor
	if err := r.database.
		Where("name LIKE ?", n).Offset(offset).Limit(limit).
		Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

// GetDataByID implements VendorRepository.
func (r *Vendor) GetDataByID(id string) (*model.Vendor, error) {
	data := new(model.Vendor)
	if err := r.database.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func NewVendor(db *gorm.DB) VendorRepository {
	return &Vendor{
		database: db,
	}
}
