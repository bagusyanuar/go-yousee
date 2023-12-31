package common

const (
	PageLimit = 10
)

type Pagination struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalRows  int64 `json:"total_rows"`
	TotalPages int   `json:"total_pages"`
	Rows       any   `json:"rows"`
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = PageLimit
	}
	return p.Limit
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}
