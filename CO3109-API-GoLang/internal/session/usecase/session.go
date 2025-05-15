package usecase

import (
	"context"
	"database/sql"

	"gitlab.com/tantai-smap/authenticate-api/internal/session"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (uc implUsecase) Create(ctx context.Context, sc scope.Scope, ip session.CreateSessionInput) (session.CreateSessionOutput, error) {
	u, err := uc.userUC.Detail(ctx, sc, ip.UserID)
	if err != nil {
		uc.l.Errorf(ctx, "session.usecase.Create.userRepo.Detail: %v", err)
		return session.CreateSessionOutput{}, err
	}

	if u.User.IsVerified.IsZero() || u.User.IsVerified.Bool == false {
		uc.l.Warnf(ctx, "session.usecase.Create.userRepo.Detail: %v", err)
		return session.CreateSessionOutput{}, session.ErrUserNotVerified
	}

	r, err := uc.repo.Create(ctx, sc, session.CreateSessionOptions(ip))
	if err != nil {
		if err == sql.ErrNoRows {
			uc.l.Warnf(ctx, "session.usecase.Create.repo.Create: %v", err)
			return session.CreateSessionOutput{}, session.ErrSessionNotFound
		}
		uc.l.Errorf(ctx, "session.usecase.Create.repo.Create: %v", err)
		return session.CreateSessionOutput{}, err
	}

	return session.CreateSessionOutput{
		Session: r,
	}, nil
}
