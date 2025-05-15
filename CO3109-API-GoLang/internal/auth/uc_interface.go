package auth

import (
	"context"

	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

//go:generate mockery --name UseCase
type UseCase interface {
	Producer
	Register(ctx context.Context, sc scope.Scope, ip RegisterInput) (RegisterOutput, error)
	SendOTP(ctx context.Context, sc scope.Scope, ip SendOTPInput) error
	VerifyOTP(ctx context.Context, sc scope.Scope, ip VerifyOTPInput) error
	Login(ctx context.Context, sc scope.Scope, ip LoginInput) (LoginOutput, error)
	SocialLogin(ctx context.Context, sc scope.Scope, ip SocialLoginInput) (SocialLoginOutput, error)
	SocialCallback(ctx context.Context, sc scope.Scope, ip SocialCallbackInput) (SocialCallbackOutput, error)
	DetailMe(ctx context.Context, sc scope.Scope) (DetailMeOutput, error)
}

type Producer interface {
	PubSendEmailMsg(ctx context.Context, sc scope.Scope, ip PubSendEmailMsgInput) error
}
