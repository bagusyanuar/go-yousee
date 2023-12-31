package service

import (
	"math"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
)

type (
	VendorService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		GetDataByID(id string) (*model.Vendor, error)
		Create(request request.VendorRequest) (*model.Vendor, error)
		Patch(id string, request request.VendorRequest) (*model.Vendor, error)
		Delete(id string) error
	}

	Vendor struct {
		vendorRepository repositories.VendorRepository
	}
)

// Delete implements VendorService.
func (svc *Vendor) Delete(id string) error {
	return svc.vendorRepository.Delete(id)
}

// GetDataByID implements VendorService.
func (svc *Vendor) GetDataByID(id string) (*model.Vendor, error) {
	return svc.vendorRepository.GetDataByID(id)
}

// Patch implements VendorService.
func (svc *Vendor) Patch(id string, request request.VendorRequest) (*model.Vendor, error) {
	entity := model.Vendor{
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
		Brand:   request.Brand,
	}
	return svc.vendorRepository.Patch(id, entity)
}

// Create implements VendorService.
func (svc *Vendor) Create(request request.VendorRequest) (*model.Vendor, error) {
	entity := model.Vendor{
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
		Brand:   request.Brand,
	}
	return svc.vendorRepository.Create(entity)
}

// GetData implements VendorService.
func (svc *Vendor) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.vendorRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.Vendor{}
		return pagination, err
	}

	data, err := svc.vendorRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.Vendor{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

func NewVendor(vendorRepo repositories.VendorRepository) VendorService {
	return &Vendor{
		vendorRepository: vendorRepo,
	}
}
