package request

type ProvinceRequest struct {
	Name string `json:"name" validate:"required"`
	Code int    `json:"code" validate:"required,numeric,len=4"`
}
