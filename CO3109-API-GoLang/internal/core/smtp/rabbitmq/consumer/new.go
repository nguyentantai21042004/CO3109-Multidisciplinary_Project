package consumer

import (
	"gitlab.com/tantai-smap/authenticate-api/internal/core/smtp"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
	rabbitmqPkg "gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
)

type Consumer struct {
	l    log.Logger
	conn *rabbitmqPkg.Connection
	uc   smtp.UseCase
}

func NewConsumer(l log.Logger, conn *rabbitmqPkg.Connection, uc smtp.UseCase) Consumer {
	return Consumer{
		l:    l,
		conn: conn,
		uc:   uc,
	}
}
