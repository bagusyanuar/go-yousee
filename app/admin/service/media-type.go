package service

import (
	"fmt"
	"math"
	"path/filepath"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

const (
	IconPath = "assets/media-type"
)

type (
	MediaTypeService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		Create(request request.MediaTypeRequest) (*model.MediaType, error)
		GetDataByID(id string) (*model.MediaType, error)
		Patch(id string, request request.MediaTypeRequest) (*model.MediaType, error)
		Delete(id string) error
	}

	MediaType struct {
		mediaTypeRepository repositories.MediaTypeRepository
	}
)

// Patch implements MediaTypeService.
func (svc *MediaType) Patch(id string, request request.MediaTypeRequest) (*model.MediaType, error) {
	icon := new(string)
	entity := model.MediaType{
		Name: request.Name,
		Slug: slug.Make(request.Name),
	}

	if request.Icon != nil {
		fileSystem := common.FileSystem{
			File: request.Icon,
		}
		if err := fileSystem.CheckPath(IconPath); err != nil {
			return nil, err
		}

		ext := filepath.Ext(request.Icon.Filename)
		fileName := fmt.Sprintf("%s/%s%s", IconPath, uuid.New().String(), ext)
		icon = &fileName
		err := fileSystem.Upload(fileName)
		if err != nil {
			return nil, err
		}
		entity.Icon = icon
	}

	return svc.mediaTypeRepository.Patch(id, entity)
}

// Delete implements MediaTypeService.
func (svc *MediaType) Delete(id string) error {
	return svc.mediaTypeRepository.Delete(id)
}

// GetDataByID implements MediaTypeService.
func (svc *MediaType) GetDataByID(id string) (*model.MediaType, error) {
	return svc.mediaTypeRepository.GetDataByID(id)
}

// Create implements TypeService.
func (svc *MediaType) Create(request request.MediaTypeRequest) (*model.MediaType, error) {
	icon := new(string)
	if request.Icon != nil {
		fileSystem := common.FileSystem{
			File: request.Icon,
		}
		if err := fileSystem.CheckPath(IconPath); err != nil {
			return nil, err
		}

		ext := filepath.Ext(request.Icon.Filename)
		fileName := fmt.Sprintf("%s/%s%s", IconPath, uuid.New().String(), ext)
		icon = &fileName
		err := fileSystem.Upload(fileName)
		if err != nil {
			return nil, err
		}
	}
	entity := model.MediaType{
		Name: request.Name,
		Icon: icon,
		Slug: slug.Make(request.Name),
	}
	return svc.mediaTypeRepository.Create(entity)
}

// GetData implements TypeService.
func (svc *MediaType) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.mediaTypeRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.MediaType{}
		return pagination, err
	}

	data, err := svc.mediaTypeRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.MediaType{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewMediaType(mediaTypeRepo repositories.MediaTypeRepository) MediaTypeService {
	return &MediaType{
		mediaTypeRepository: mediaTypeRepo,
	}
}
