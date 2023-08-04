package request

// len used for character length for string
type ProvinceRequest struct {
	Name string `json:"name" validate:"required"`
	Code string `json:"code" validate:"required,numeric,len=4"`
}
