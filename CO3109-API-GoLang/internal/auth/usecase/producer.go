package usecase

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/internal/auth"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

func (u implUsecase) PubSendEmailMsg(ctx context.Context, sc scope.Scope, ip auth.PubSendEmailMsgInput) error {
	err := u.prod.PubSendEmail(ctx, u.toSendEmailMsg(ip))
	if err != nil {
		u.l.Error(ctx, "auth.usecase.producer.PubSendEmailMsg: %v", err)
		return err
	}
	return nil
}
