package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type (
	ItemService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		GetDataByID(id string) (*model.Item, error)
		Create(request request.ItemRequest) (*model.Item, error)
		Patch(id string, request request.ItemRequest) (*model.Item, error)
		Delete(id string) error
	}

	Item struct {
		itemRepository repositories.ItemRepository
	}
)

// Delete implements ItemService.
func (svc *Item) Delete(id string) error {
	return svc.itemRepository.Delete(id)
}

// GetDataByID implements ItemService.
func (svc *Item) GetDataByID(id string) (*model.Item, error) {
	return svc.itemRepository.GetDataByID(id)
}

// Patch implements ItemService.
func (svc *Item) Patch(id string, request request.ItemRequest) (*model.Item, error) {
	cityID, _ := uuid.Parse(request.CityID)
	vendorID, _ := uuid.Parse(request.VendorID)
	mediaTypeID, _ := uuid.Parse(request.MediaTypeID)

	entity := model.Item{
		CityID:      cityID,
		VendorID:    vendorID,
		MediaTypeID: mediaTypeID,
		Name:        cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Address:     request.Address,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		Position:    request.Position,
		Width:       request.Width,
		Height:      request.Height,
	}
	return svc.itemRepository.Patch(id, entity)
}

// Create implements ItemService.
func (svc *Item) Create(request request.ItemRequest) (*model.Item, error) {
	cityID, _ := uuid.Parse(request.CityID)
	vendorID, _ := uuid.Parse(request.VendorID)
	mediaTypeID, _ := uuid.Parse(request.MediaTypeID)

	entity := model.Item{
		CityID:      cityID,
		VendorID:    vendorID,
		MediaTypeID: mediaTypeID,
		Name:        cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		Address:     request.Address,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		Position:    request.Position,
		Width:       request.Width,
		Height:      request.Height,
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
