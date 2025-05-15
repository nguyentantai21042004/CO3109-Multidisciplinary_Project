package main

import (
	"context"

	_ "github.com/lib/pq"
	"gitlab.com/tantai-smap/authenticate-api/config"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/oauth"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/postgres"
	"gitlab.com/tantai-smap/authenticate-api/internal/consumer"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/cloudinary"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/ggdrive"
	pkgCrt "gitlab.com/tantai-smap/authenticate-api/pkg/encrypter"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
	pkgRedis "gitlab.com/tantai-smap/authenticate-api/pkg/redis"
	"golang.org/x/oauth2"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// Encrypter
	crp := pkgCrt.NewEncrypter(cfg.Encrypter.Key)

	// Postgres
	database, err := postgres.Connect(context.Background(), cfg.Postgres)
	if err != nil {
		panic(err)
	}
	defer postgres.Disconnect(context.Background(), database)

	// RabbitMQ
	conn, err := rabbitmq.Dial(cfg.RabbitMQConfig.URL, true)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Redis
	redisClient, err := pkgRedis.Connect(pkgRedis.NewClientOptions().SetOptions(cfg.Redis))
	if err != nil {
		panic(err)
	}
	defer redisClient.Disconnect()

	// Oauth
	oauthConfig := oauth.NewOauthConfig(cfg.Oauth)

	// Google Drive
	ggDriveConfig := &oauth2.Config{
		ClientID:     cfg.GoogleDrive.ClientID,
		ClientSecret: cfg.GoogleDrive.ClientSecret,
		RedirectURL:  cfg.GoogleDrive.RedirectURL,
	}
	ggdrive, err := ggdrive.New(context.Background(), ggDriveConfig)
	if err != nil {
		panic(err)
	}

	// Cloudinary
	cloudinary, err := cloudinary.New(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		panic(err)
	}

	// Logger
	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	// Consumer
	srv, err := consumer.New(l, consumer.ConsumerConfig{
		Encrypter:    crp,
		JwtSecretKey: cfg.JWT.SecretKey,
		Telegram: consumer.TeleCredentials{
			BotKey: cfg.Telegram.BotKey,
			ChatIDs: consumer.ChatIDs{
				ReportBug: cfg.Telegram.ChatIDs.ReportBug,
			},
		},
		InternalKey: cfg.InternalConfig.InternalKey,
		Database:    database,
		SMTPConfig:  cfg.SMTP,
		AMQPConn:    conn,
		RedisClient: &redisClient,
		OauthConfig: oauthConfig,
		GoogleDrive: ggdrive,
		Cloudinary:  cloudinary,
	})
	if err != nil {
		panic(err)
	}

	// Run
	err = srv.Run()
	if err != nil {
		panic(err)
	}
}
