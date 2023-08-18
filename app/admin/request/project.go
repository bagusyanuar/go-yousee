package request

type ProjectRequest struct {
	Name         string `json:"name" validate:"required"`
	ClientName   string `json:"client_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Duration     uint   `json:"duration" validate:"required"`
	DurationUnit string `json:"duration_unit" validate:"required"`
}
