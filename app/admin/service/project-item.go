package service

import (
	"fmt"
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"github.com/google/uuid"
)

type (
	ProjectItemService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		GetDataByID(id string) (*model.ProjectItem, error)
		Create(request request.ProjectItemRequest) (*model.ProjectItem, error)
		Patch(id string, request request.ProjectItemRequest) (*model.ProjectItem, error)
		Delete(id string) error
	}

	ProjectItem struct {
		projectItemRepository repositories.ProjectItemRepository
	}
)

// Create implements ProjectItemService.
func (svc *ProjectItem) Create(request request.ProjectItemRequest) (*model.ProjectItem, error) {
	cityID, _ := uuid.Parse(request.CityID)
	reqItemID, _ := uuid.Parse(request.ItemID)
	ProjectID, _ := uuid.Parse(request.ProjectID)
	PicID, _ := uuid.Parse(request.PicID)

	itemID := new(uuid.UUID)
	fmt.Println(itemID)
	if reqItemID != uuid.Nil {
		itemID = &reqItemID
	} else {
		itemID = nil
	}
	entity := model.ProjectItem{
		CityID:      cityID,
		ItemID:      itemID,
		ProjectID:   &ProjectID,
		PicID:       PicID,
		VendorPrice: request.VendorPrice,
	}
	return svc.projectItemRepository.Create(entity)
}

// Delete implements ProjectItemService.
func (svc *ProjectItem) Delete(id string) error {
	panic("unimplemented")
}

// GetData implements ProjectItemService.
func (svc *ProjectItem) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.projectItemRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.ProjectItem{}
		return pagination, err
	}

	data, err := svc.projectItemRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.ProjectItem{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

// GetDataByID implements ProjectItemService.
func (svc *ProjectItem) GetDataByID(id string) (*model.ProjectItem, error) {
	return svc.projectItemRepository.GetDataByID(id)
}

// Patch implements ProjectItemService.
func (svc *ProjectItem) Patch(id string, request request.ProjectItemRequest) (*model.ProjectItem, error) {
	panic("unimplemented")
}

func NewProjectItem(projectItemRepo repositories.ProjectItemRepository) ProjectItemService {
	return &ProjectItem{
		projectItemRepository: projectItemRepo,
	}
}
