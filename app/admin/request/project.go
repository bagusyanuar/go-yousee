package request

type ProjectRequest struct {
	Name         string `json:"name" validate:"required"`
	ClientName   string `json:"client_name" validate:"required"`
	Description  string `json:"description" validate:"required"`
	Duration     uint   `json:"duration" validate:"required"`
	Qty          uint   `json:"qty" validate:"required,numeric"`
	IsLightOn    bool   `json:"is_light_on" validate:"required,boolean"`
	DurationUnit string `json:"duration_unit" validate:"required"`
}
