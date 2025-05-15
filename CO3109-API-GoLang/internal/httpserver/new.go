package httpserver

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	"gitlab.com/tantai-smap/authenticate-api/config"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/oauth"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/cloudinary"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/ggdrive"
	pkgCrt "gitlab.com/tantai-smap/authenticate-api/pkg/encrypter"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
	pkgRabbitMQ "gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
	"gitlab.com/tantai-smap/authenticate-api/pkg/redis"
)

type HTTPServer struct {
	gin          *gin.Engine
	l            pkgLog.Logger
	jwtSecretKey string
	mode         string
	encrypter    pkgCrt.Encrypter
	host         string
	port         int
	telegram     TeleCredentials
	internalKey  string
	database     *sql.DB
	amqpConn     *pkgRabbitMQ.Connection
	smtpConfig   config.SMTPConfig
	redisClient  *redis.Client
	oauthConfig  oauth.OauthConfig
	googleDrive  ggdrive.Usecase
	cloudinary   cloudinary.Usecase
}

type Config struct {
	JwtSecretKey string
	Mode         string
	Encrypter    pkgCrt.Encrypter
	Host         string
	Port         int
	Telegram     TeleCredentials
	InternalKey  string
	Database     *sql.DB
	AMQPConn     *pkgRabbitMQ.Connection
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

func New(l pkgLog.Logger, cfg Config) (*HTTPServer, error) {
	if cfg.Mode == productionMode {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	h := &HTTPServer{
		l:            l,
		gin:          gin.Default(),
		jwtSecretKey: cfg.JwtSecretKey,
		mode:         cfg.Mode,
		encrypter:    cfg.Encrypter,
		host:         cfg.Host,
		port:         cfg.Port,
		telegram:     cfg.Telegram,
		internalKey:  cfg.InternalKey,
		database:     cfg.Database,
		amqpConn:     cfg.AMQPConn,
		smtpConfig:   cfg.SMTPConfig,
		redisClient:  cfg.RedisClient,
		oauthConfig:  cfg.OauthConfig,
		googleDrive:  cfg.GoogleDrive,
		cloudinary:   cfg.Cloudinary,
	}

	if err := h.validate(); err != nil {
		return nil, err
	}

	return h, nil
}

func (s HTTPServer) validate() error {
	requiredDeps := []struct {
		dep interface{}
		msg string
	}{
		{s.l, "logger is required"},
		{s.mode, "mode is required"},
		{s.jwtSecretKey, "jwtSecretKey is required"},
		{s.encrypter, "encrypter is required"},
		{s.telegram, "telegram is required"},
		{s.internalKey, "internalKey is required"},
		{s.database, "database is required"},
		{s.smtpConfig, "smtpConfig is required"},
		{s.amqpConn, "amqpConn is required"},
		{s.redisClient, "redisClient is required"},
		{s.cloudinary, "cloudinary is required"},
	}

	for _, dep := range requiredDeps {
		if dep.dep == nil {
			return errors.New(dep.msg)
		}
	}

	return nil
}
