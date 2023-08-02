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
	panic("unimplemented")
}

// GetData implements ProvinceRepository.
func (r *Province) GetData(name string, limit int, offset int) ([]model.Province, error) {
	panic("unimplemented")
}

// GetDataByID implements ProvinceRepository.
func (r *Province) GetDataByID(id string) (*model.Province, error) {
	panic("unimplemented")
}

func NewProvince(db *gorm.DB) ProvinceRepository {
	return &Province{
		database: db,
	}
}
