package consumer

import (
	"database/sql"
	"errors"

	"gitlab.com/tantai-smap/authenticate-api/config"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/oauth"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/cloudinary"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/ggdrive"
	pkgCrt "gitlab.com/tantai-smap/authenticate-api/pkg/encrypter"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
	"gitlab.com/tantai-smap/authenticate-api/pkg/redis"
)

type Consumer struct {
	l            pkgLog.Logger
	jwtSecretKey string
	amqpConn     *rabbitmq.Connection
	encrypter    pkgCrt.Encrypter
	telegram     TeleCredentials
	internalKey  string
	database     *sql.DB
	smtpConfig   config.SMTPConfig
	redisClient  *redis.Client
	oauthConfig  oauth.OauthConfig
	googleDrive  ggdrive.Usecase
	cloudinary   cloudinary.Usecase
}

type ConsumerConfig struct {
	JwtSecretKey string
	AMQPConn     *rabbitmq.Connection
	Encrypter    pkgCrt.Encrypter
	Telegram     TeleCredentials
	InternalKey  string
	Database     *sql.DB
	SMTPConfig   config.SMTPConfig
	RedisClient  *redis.Client
	OauthConfig  oauth.OauthConfig
	GoogleDrive  ggdrive.Usecase
	Cloudinary   cloudinary.Usecase
}

type TeleCredentials struct {
	BotKey string
	ChatIDs
}

type ChatIDs struct {
	ReportBug int64
}

func New(l pkgLog.Logger, cfg ConsumerConfig) (*Consumer, error) {

	h := &Consumer{
		l:            l,
		amqpConn:     cfg.AMQPConn,
		jwtSecretKey: cfg.JwtSecretKey,
		encrypter:    cfg.Encrypter,
		telegram:     cfg.Telegram,
		internalKey:  cfg.InternalKey,
		database:     cfg.Database,
		smtpConfig:   cfg.SMTPConfig,
		redisClient:  cfg.RedisClient,
		oauthConfig:  cfg.OauthConfig,
		googleDrive:  cfg.GoogleDrive,
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	return h, nil
}

func (s Consumer) validate() error {
	requiredDeps := []struct {
		dep interface{}
		msg string
	}{
		{s.l, "logger is required"},
		{s.amqpConn, "amqpConn is required"},
		{s.jwtSecretKey, "jwtSecretKey is required"},
		{s.encrypter, "encrypter is required"},
		{s.telegram, "telegram is required"},
		{s.internalKey, "internalKey is required"},
		{s.database, "database is required"},
		{s.smtpConfig, "smtpConfig is required"},
		{s.redisClient, "redisClient is required"},
		{s.oauthConfig, "oauthConfig is required"},
		{s.googleDrive, "googleDrive is required"},
	}

	for _, dep := range requiredDeps {
		if dep.dep == nil {
			return errors.New(dep.msg)
		}
	}

	return nil
}
