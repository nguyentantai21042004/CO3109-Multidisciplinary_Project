package producer

import (
	"context"

	rabb "gitlab.com/tantai-smap/authenticate-api/internal/auth/delivery/rabbitmq"
	"gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
)

//go:generate mockery --name=Producer
type Producer interface {
	PubSendEmail(ctx context.Context, msg rabb.SendEmailMsg) error
	// Run runs the producer
	Run() error
	// Close closes the producer
	Close()
}

type implProducer struct {
	l               log.Logger
	conn            *rabbitmq.Connection
	sendEmailWriter *rabbitmq.Channel
}

var _ Producer = &implProducer{}

func NewProducer(l log.Logger, conn *rabbitmq.Connection) Producer {
	return &implProducer{
		l:    l,
		conn: conn,
	}
}
