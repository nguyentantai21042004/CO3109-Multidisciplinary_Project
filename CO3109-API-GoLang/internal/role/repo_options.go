package role

import pag "gitlab.com/tantai-smap/authenticate-api/pkg/paginator"

type Filter struct {
	IDs   []string
	Alias []string
	Code  []string
}

type GetOneOptions struct {
	Filter Filter
}

type GetOptions struct {
	Filter   Filter
	PagQuery pag.PaginateQuery
}

type ListOptions struct {
	Filter Filter
}
