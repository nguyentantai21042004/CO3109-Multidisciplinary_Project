package usecase

import (
	"time"

	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/oauth"
	"gitlab.com/tantai-smap/authenticate-api/internal/auth"
	"gitlab.com/tantai-smap/authenticate-api/internal/auth/delivery/rabbitmq/producer"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/smtp"
	"gitlab.com/tantai-smap/authenticate-api/internal/role"
	"gitlab.com/tantai-smap/authenticate-api/internal/session"
	"gitlab.com/tantai-smap/authenticate-api/internal/user"
	"gitlab.com/tantai-smap/authenticate-api/pkg/encrypter"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/redis"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
)

type implUsecase struct {
	l         log.Logger
	prod      producer.Producer
	encrypt   encrypter.Encrypter
	redis     redis.Client
	oauth     oauth.OauthConfig
	scope     scope.Manager
	smtp      smtp.UseCase
	userUC    user.UseCase
	roleUC    role.UseCase
	sessionUC session.UseCase
	clock     func() time.Time
}

var _ auth.UseCase = &implUsecase{}

func New(l log.Logger,
	prod producer.Producer,
	encrypt encrypter.Encrypter,
	redis redis.Client,
	oauth oauth.OauthConfig,
	scope scope.Manager,
	smtp smtp.UseCase,
	userUC user.UseCase,
	roleUC role.UseCase,
	sessionUC session.UseCase,
) auth.UseCase {
	return &implUsecase{
		l:         l,
		prod:      prod,
		encrypt:   encrypt,
		redis:     redis,
		oauth:     oauth,
		scope:     scope,
		smtp:      smtp,
		userUC:    userUC,
		roleUC:    roleUC,
		sessionUC: sessionUC,
		clock:     time.Now,
	}
}
