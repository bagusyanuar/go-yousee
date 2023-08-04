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
	ItemService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		Create(request request.ItemRequest) (*model.Item, error)
	}

	Item struct {
		itemRepository repositories.ItemRepository
	}
)

// Create implements ItemService.
func (svc *Item) Create(request request.ItemRequest) (*model.Item, error) {
	cityID, _ := uuid.Parse(request.CityID)
	vendorID, _ := uuid.Parse(request.VendorID)
	typeID, _ := uuid.Parse(request.TypeID)

	entity := model.Item{
		CityID:    cityID,
		VendorID:  vendorID,
		TypeID:    typeID,
		Name:      request.Name,
		Address:   request.Address,
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
		Position:  request.Position,
		Width:     request.Width,
		Height:    request.Height,
	}
	return svc.itemRepository.Create(entity)
}

// GetData implements ItemService.
func (svc *Item) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.itemRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.Item{}
		return pagination, err
	}

	data, err := svc.itemRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.Item{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewItem(itemRepo repositories.ItemRepository) ItemService {
	return &Item{
		itemRepository: itemRepo,
	}
}
