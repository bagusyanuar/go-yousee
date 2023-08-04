package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"github.com/google/uuid"
)

type (
	CityService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		Create(request request.CityRequest) (*model.City, error)
	}

	City struct {
		cityRepository repositories.CityRepository
	}
)

// Create implements CityService.
func (svc *City) Create(request request.CityRequest) (*model.City, error) {
	provinceID, _ := uuid.Parse(request.ProvinceID)
	entity := model.City{
		ProvinceID: provinceID,
		Name:       request.Name,
		Code:       request.Code,
	}
	return svc.cityRepository.Create(entity)
}

// GetData implements CityService.
func (svc *City) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.cityRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.City{}
		return pagination, err
	}

	data, err := svc.cityRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.City{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewCity(cityRepo repositories.CityRepository) CityService {
	return &City{
		cityRepository: cityRepo,
	}
}
