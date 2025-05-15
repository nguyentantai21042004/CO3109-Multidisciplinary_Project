package consumer

import (
	smtpConsumer "gitlab.com/tantai-smap/authenticate-api/internal/core/smtp/rabbitmq/consumer"
	smtpUC "gitlab.com/tantai-smap/authenticate-api/internal/core/smtp/usecase"
)

func (srv Consumer) mapHandlers() error {
	smtpUC := smtpUC.New(srv.l, srv.smtpConfig)

	// roleRepoPostgres := roleRepoPostgres.New(srv.l, srv.database)
	// roleUC := roleUsecase.New(srv.l, roleRepoPostgres)

	// userRepoPostgres := userRepoPostgres.New(srv.l, srv.database)
	// userUC := userUsecase.New(srv.l, userRepoPostgres, roleUC)

	// authUC := authUC.New(srv.l, srv.encrypter, userUC, smtpUC)

	var forever chan bool
	smtpConsumer.NewConsumer(srv.l, srv.amqpConn, smtpUC).Consume()
	<-forever

	return nil
}
