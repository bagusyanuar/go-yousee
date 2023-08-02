package common

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    any    `json:"meta"`
	Data    any    `json:"data"`
}

type PaginationMeta struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
}
