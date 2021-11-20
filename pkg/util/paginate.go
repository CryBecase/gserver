package util

type Paginate struct {
	Page  int `json:"page"  form:"page"`
	Size  int `json:"size"  form:"size"`
	Total int `json:"total" form:"total"`
}

func (p *Paginate) Fix() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Size <= 0 {
		p.Size = 10
	}
}

func (p *Paginate) OffSize() int {
	return (p.Page - 1) * p.Size
}
