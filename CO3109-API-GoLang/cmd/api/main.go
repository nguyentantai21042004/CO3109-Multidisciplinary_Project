package main

import (
	"context"

	_ "github.com/lib/pq"
	"gitlab.com/tantai-smap/authenticate-api/config"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/oauth"
	"gitlab.com/tantai-smap/authenticate-api/internal/appconfig/postgres"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/cloudinary"
	"gitlab.com/tantai-smap/authenticate-api/internal/core/ggdrive"
	"gitlab.com/tantai-smap/authenticate-api/internal/httpserver"
	pkgCrt "gitlab.com/tantai-smap/authenticate-api/pkg/encrypter"
	pkgLog "gitlab.com/tantai-smap/authenticate-api/pkg/log"
	"gitlab.com/tantai-smap/authenticate-api/pkg/rabbitmq"
	pkgRedis "gitlab.com/tantai-smap/authenticate-api/pkg/redis"
	"golang.org/x/oauth2"
)

// @title CO3109 Multidisciplinary Project
// @description This is the API documentation for CO3109 Multidisciplinary Project-GoLang.
// @description authenticate
// @description `110001 ("Wrong query"),`
// @description `110002 ("Wrong body"),`
// @description `110003 ("User not found"),`
// @description `110004 ("Email existed"),`
// @description `110005 ("Wrong password"),`
// @description user
// @description `120001 ("Wrong query"),`
// @description `120002 ("Wrong body"),`
// @description `120007 ("Permission denied"),`
// @description `120008 ("User existed"),`
// @description `120009 ("User not found"),`
// @description `120010 ("Role not found"),`
// @description upload
// @description `130001 ("Wrong query"),`
// @description `130002 ("Wrong body"),`
// @description `130003 ("Unauthorized"),`
// @description `130004 ("Invalid URL"),`
// @description `130005 ("Invalid Path"),`
// @description `130006 ("Upload not found"),`
// @version 1
// @host localhost:8085
// @schemes http
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
	cloudinaryClient, err := cloudinary.New(cfg.Cloudinary.CloudName, cfg.Cloudinary.APIKey, cfg.Cloudinary.APISecret)
	if err != nil {
		panic(err)
	}

	// Logger
	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	// HTTPServer
	srv, err := httpserver.New(l, httpserver.Config{
		Host:         cfg.HTTPServer.Host,
		Port:         cfg.HTTPServer.Port,
		Encrypter:    crp,
		JwtSecretKey: cfg.JWT.SecretKey,
		Mode:         cfg.HTTPServer.Mode,
		Telegram: httpserver.TeleCredentials{
			BotKey: cfg.Telegram.BotKey,
			ChatIDs: httpserver.ChatIDs{
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
		Cloudinary:  cloudinaryClient,
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
