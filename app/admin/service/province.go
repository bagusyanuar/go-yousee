package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ProvinceService interface {
	GetData(name string, page, perPage int) (common.Pagination, error)
	GetDataByID(id string) (*model.Province, error)
	Create(request request.ProvinceRequest) (*model.Province, error)
	Patch(id string, request request.ProvinceRequest) (*model.Province, error)
	Delete(id string) error
}

type Province struct {
	provinceRepository repositories.ProvinceRepository
}

// Delete implements ProvinceService.
func (svc *Province) Delete(id string) error {
	return svc.provinceRepository.Delete(id)
}

// Patch implements ProvinceService.
func (svc *Province) Patch(id string, request request.ProvinceRequest) (*model.Province, error) {
	entity := model.Province{
		Name: cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Code: request.Code,
	}
	return svc.provinceRepository.Patch(id, entity)
}

// GetDataByID implements ProvinceService.
func (svc *Province) GetDataByID(id string) (*model.Province, error) {
	return svc.provinceRepository.GetDataByID(id)
}

// Create implements ProvinceService.
func (svc *Province) Create(request request.ProvinceRequest) (*model.Province, error) {
	entity := model.Province{
		Name: cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Code: request.Code,
	}
	return svc.provinceRepository.Create(entity)
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
