package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
)

type ProvinceService interface {
	GetData(name string, page, perPage int) (common.Pagination, error)
}

type Province struct {
	provinceRepository repositories.ProvinceRepository
}

// GetData implements ProvinceService.
func (svc *Province) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.provinceRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.Province{}
		return pagination, err
	}

	data, err := svc.provinceRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.Province{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewProvince(provinceRepo repositories.ProvinceRepository) ProvinceService {
	return &Province{
		provinceRepository: provinceRepo,
	}
}
