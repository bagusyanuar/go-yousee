package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
)

type (
	TypeService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		Create(request request.TypeRequest) (*model.Type, error)
	}

	Type struct {
		typeRepository repositories.TypeRepository
	}
)

// Create implements TypeService.
func (svc *Type) Create(request request.TypeRequest) (*model.Type, error) {
	entity := model.Type{
		Name: request.Name,
	}
	return svc.typeRepository.Create(entity)
}

// GetData implements TypeService.
func (svc *Type) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.typeRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.Type{}
		return pagination, err
	}

	data, err := svc.typeRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.Type{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewType(typeRepo repositories.TypeRepository) TypeService {
	return &Type{
		typeRepository: typeRepo,
	}
}
