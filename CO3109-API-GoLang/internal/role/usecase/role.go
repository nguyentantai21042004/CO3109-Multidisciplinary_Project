package usecase

import (
	"context"
	"database/sql"

	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (uc implUsecase) Detail(ctx context.Context, sc scope.Scope, ID string) (role.DetailOutput, error) {
	r, err := uc.repo.Detail(ctx, sc, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "role.usecase.Detail.repo.Detail: %v", err)
			return role.DetailOutput{}, role.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "role.usecase.Detail.repo.Detail: %v", err)
		return role.DetailOutput{}, err
	}

	return role.DetailOutput{
		Role: r,
	}, nil
}

func (uc implUsecase) GetOne(ctx context.Context, sc scope.Scope, ip role.GetOneInput) (role.GetOneOutput, error) {
	r, err := uc.repo.GetOne(ctx, sc, role.GetOneOptions(ip))
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "role.usecase.GetOne.repo.GetOne: %v", err)
			return role.GetOneOutput{}, role.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "role.usecase.GetOne.repo.GetOne: %v", err)
		return role.GetOneOutput{}, err
	}

	return role.GetOneOutput{
		Role: r,
	}, nil
}

func (uc implUsecase) Get(ctx context.Context, sc scope.Scope, ip role.GetInput) (role.GetOutput, error) {
	rs, pag, err := uc.repo.Get(ctx, sc, role.GetOptions(ip))
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "role.usecase.Get.repo.Get: %v", err)
			return role.GetOutput{}, role.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "role.usecase.Get.repo.Get: %v", err)
		return role.GetOutput{}, err
	}

	return role.GetOutput{
		Roles:     rs,
		Paginator: pag,
	}, nil
}

func (uc implUsecase) List(ctx context.Context, sc scope.Scope, ip role.ListInput) (role.ListOutput, error) {
	rs, err := uc.repo.List(ctx, sc, role.ListOptions(ip))
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "role.usecase.List.repo.List: %v", err)
			return role.ListOutput{}, role.ErrRoleNotFound
		}
		uc.l.Errorf(ctx, "role.usecase.List.repo.List: %v", err)
		return role.ListOutput{}, err
	}

	return role.ListOutput{
		Roles: rs,
	}, nil
}
