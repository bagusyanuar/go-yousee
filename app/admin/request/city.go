package request

type CityRequest struct {
	ProvinceID string `json:"province_id" validate:"required,uuid4"`
	Name       string `json:"name" validate:"required"`
	Code       string `json:"code" validate:"required,numeric,len=6"`
}
