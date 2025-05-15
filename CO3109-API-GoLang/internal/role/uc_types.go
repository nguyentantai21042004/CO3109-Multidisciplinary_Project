package role

import (
	"gitlab.com/tantai-smap/authenticate-api/internal/models"
	pag "gitlab.com/tantai-smap/authenticate-api/pkg/paginator"
)

type GetOneInput struct {
	Filter Filter
}

type GetInput struct {
	Filter   Filter
	PagQuery pag.PaginateQuery
}

type ListInput struct {
	Filter Filter
}

type DetailOutput struct {
	Role models.Role
}

type GetOneOutput struct {
	Role models.Role
}

type GetOutput struct {
	Roles     []models.Role
	Paginator pag.Paginator
}

type ListOutput struct {
	Roles []models.Role
}
