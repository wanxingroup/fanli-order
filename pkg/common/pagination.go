package common

import (
	"strconv"
)

type Pagination struct {
	PageNumber uint64 `json:"pageNumber,omitempty"`
	PageSize   uint64 `json:"pageSize,omitempty"`
	PageTotal  uint64 `json:"pageTotal,omitempty"`
}

const (
	DefaultPageSize   = 10
	DefaultPageNumber = 1
)

func (p *Pagination) SetPageNumber(pageString string) {
	if len(pageString) == 0 {
		p.PageNumber = DefaultPageNumber
		return
	}
	page, err := strconv.ParseUint(pageString, 10, 64)
	if err != nil {
		p.PageNumber = DefaultPageNumber
		return
	}
	p.PageNumber = page
	return
}

func (p *Pagination) GetPageNumber() uint64 {
	return p.PageNumber
}

func (p *Pagination) SetPageSize(pageSizeString string) {
	if len(pageSizeString) == 0 {
		p.PageSize = DefaultPageSize
		return
	}
	pageSize, err := strconv.ParseUint(pageSizeString, 10, 64)
	if err != nil {
		p.PageSize = DefaultPageSize
		return
	}
	p.PageSize = pageSize
	return
}

func (p *Pagination) GetPageSize() uint64 {
	return p.PageSize
}
