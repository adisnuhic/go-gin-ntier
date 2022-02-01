package viewmodels

import "github.com/adisnuhic/hearken/pkg/paging"

type PaginatedModel struct {
	Results   interface{}       `json:"results"`
	Paginator *paging.Paginator `json:"paginator"`
}
