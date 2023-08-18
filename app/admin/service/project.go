package service

import (
	"math"
	"time"

	"github.com/bagusyanuar/go-yousee/app/admin/repositories"
	"github.com/bagusyanuar/go-yousee/app/admin/request"
	"github.com/bagusyanuar/go-yousee/common"
	"github.com/bagusyanuar/go-yousee/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/datatypes"
)

type (
	ProjectService interface {
		GetData(name string, page, perPage int) (common.Pagination, error)
		GetDataByID(id string) (*model.Project, error)
		Create(request request.ProjectRequest) (*model.Project, error)
	}

	Project struct {
		projectRepository repositories.ProjectRepository
	}
)

// Create implements ProjectService.
func (svc *Project) Create(request request.ProjectRequest) (*model.Project, error) {
	now := time.Now()
	entity := model.Project{
		Name:         cases.Title(language.Indonesian, cases.Compact).String(request.Name),
		ClientName:   request.ClientName,
		RequestDate:  datatypes.Date(now),
		Qty:          request.Qty,
		Description:  request.Description,
		IsLightOn:    request.IsLightOn,
		Duration:     request.Duration,
		DurationUnit: common.DurationUnit(request.DurationUnit),
		Status:       0,
	}
	return svc.projectRepository.Create(entity)
}

// GetData implements ProjectService.
func (svc *Project) GetData(name string, page int, perPage int) (common.Pagination, error) {
	var pagination common.Pagination
	pagination.Limit = perPage
	pagination.Page = page

	totalRows, err := svc.projectRepository.Count(name)
	if err != nil {
		pagination.Rows = []model.Project{}
		return pagination, err
	}

	data, err := svc.projectRepository.GetData(name, pagination.GetLimit(), pagination.GetOffset())
	if err != nil {
		pagination.Rows = []model.Project{}
		return pagination, err
	}
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalRows = totalRows
	pagination.TotalPages = totalPages
	pagination.Rows = data
	return pagination, nil
}

// GetDataByID implements ProjectService.
func (svc *Project) GetDataByID(id string) (*model.Project, error) {
	panic("unimplemented")
}

func NewProject(projectRepo repositories.ProjectRepository) ProjectService {
	return &Project{
		projectRepository: projectRepo,
	}
}
