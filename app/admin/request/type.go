package request

type TypeRequest struct {
	Name string `json:"name" validate:"required"`
}
