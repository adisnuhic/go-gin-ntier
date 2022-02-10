package viewmodels

import "github.com/adisnuhic/go-gin-ntier/pkg/paging"

type PaginatedModel struct {
	Results   interface{}       `json:"results"`
	Paginator *paging.Paginator `json:"paginator"`
}
