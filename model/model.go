package model

import (
	"context"
	"math"
)

const (
	RequestId = "REQUEST_ID"
)

type ReqIDContextKey string

// func set context value
func NewContext(ctx context.Context, key interface{}, p interface{}) context.Context {
	return context.WithValue(ctx, key, p)
}

// func parse context get value
func ParseContext(ctx context.Context, key interface{}) interface{} {
	if v := ctx.Value(key); v != nil {
		return v
	}
	return nil
}

type Response struct {
	LogId   string      `json:"log_id"`
	Status  int         `json:"status"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

type TableAttributes struct {
	CreatedAt string `json:"created_at,omitempty" db:"created_at"`
	CreatedBy string `json:"created_by,omitempty" db:"created_by" fieldtag:"insert"`
	UpdatedAt string `json:"updated_at,omitempty" db:"updated_at"`
	UpdatedBy string `json:"updated_by,omitempty" db:"updated_by" fieldtag:"insert,update"`
}

type Pagination struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalRows   int `json:"total_rows"`
}

type PaginationRequest struct {
	Page  int `query:"page" json:"page"`
	Limit int `query:"limit" json:"limit"`
}

type (
	QueryReq struct {
		Search Search `json:"search,omitempty"`
		Filter Filter `json:"filter,omitempty"`
		Sort   Sort   `json:"sort,omitempty"`
		PaginationRequest
	}

	Filter struct {
		Author string `json:"autor,omitempty" query:"filter.author"`
		Title  string `json:"title,omitempty" query:"filter.title"`
	}

	Search struct {
		Key   string `json:"key,omitempty" query:"search.key"`
		Value string `json:"value,omitempty" query:"search.value"`
	}

	Sort struct {
		Key   string `json:"key,omitempty" query:"sort.key"`
		Value string `json:"value,omitempty" query:"sort.value"`
	}
)

func (p *Pagination) Offset() int {
	return (p.CurrentPage - 1) * p.Limit
}

func (p *Pagination) Paginate(total int) *Pagination {
	p.TotalPage = int(math.Ceil(float64(total) / float64(p.Limit)))
	p.TotalRows = total
	return p
}
