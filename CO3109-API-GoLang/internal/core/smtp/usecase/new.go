package smtp

import (
	"gitlab.com/tantai-smap/authenticate-api/config"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/smtp"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
)

type implService struct {
	l   log.Logger
	cfg config.SMTPConfig
}

func New(l log.Logger, cfg config.SMTPConfig) smtp.UseCase {
	return implService{
		l:   l,
		cfg: cfg,
	}
}
