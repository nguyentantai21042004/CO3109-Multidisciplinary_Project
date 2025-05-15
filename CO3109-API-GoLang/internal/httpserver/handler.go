package httpserver

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gitlab.com/tantai-smap/authenticate-api/internal/middleware"
	"gitlab.com/tantai-smap/authenticate-api/pkg/i18n"
	"gitlab.com/tantai-smap/authenticate-api/pkg/scope"
	"gitlab.com/tantai-smap/authenticate-api/pkg/telegram"

	roleRepoPostgres "gitlab.com/tantai-smap/authenticate-api/internal/role/repository/postgres"
	roleUsecase "gitlab.com/tantai-smap/authenticate-api/internal/role/usecase"

	authHTTP "gitlab.com/tantai-smap/authenticate-api/internal/auth/delivery/http"
	authProd "gitlab.com/tantai-smap/authenticate-api/internal/auth/delivery/rabbitmq/producer"
	authUC "gitlab.com/tantai-smap/authenticate-api/internal/auth/usecase"

	smtpUC "gitlab.com/tantai-smap/authenticate-api/internal/core/smtp/usecase"

	sessionRepoPostgres "gitlab.com/tantai-smap/authenticate-api/internal/session/repository/postgres"
	sessionUsecase "gitlab.com/tantai-smap/authenticate-api/internal/session/usecase"

	userHTTP "gitlab.com/tantai-smap/authenticate-api/internal/user/delivery/http"
	userRepoPostgres "gitlab.com/tantai-smap/authenticate-api/internal/user/repository/postgres"
	userUsecase "gitlab.com/tantai-smap/authenticate-api/internal/user/usecase"

	uploadHTTP "gitlab.com/tantai-smap/authenticate-api/internal/upload/delivery/http"
	uploadRepoPostgres "gitlab.com/tantai-smap/authenticate-api/internal/upload/repository/postgres"
	uploadUC "gitlab.com/tantai-smap/authenticate-api/internal/upload/usecase"

	// Import this to execute the init function in docs.go which setups the Swagger docs.
	_ "gitlab.com/tantai-smap/authenticate-api/docs"
)

const (
	Api         = "/api/v1"
	InternalApi = "internal/api/v1"
)

func (srv HTTPServer) mapHandlers() error {
	teleBot := telegram.NewManager(srv.telegram.BotKey)
	srv.gin.Use(middleware.Recovery(teleBot, srv.telegram.ChatIDs.ReportBug))

	//swagger api
	srv.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	scopeManager := scope.NewManager(srv.jwtSecretKey)
	// internalKey, err := srv.encrypter.Encrypt(srv.internalKey)
	// if err != nil {
	// 	srv.l.Fatal(context.Background(), err)
	// 	return err
	// }

	i18n.Init()
	smtpUC := smtpUC.New(srv.l, srv.smtpConfig)

	// Middleware
	mw := middleware.New(srv.l, scopeManager)

	roleRepoPostgres := roleRepoPostgres.New(srv.l, srv.database)
	roleUC := roleUsecase.New(srv.l, roleRepoPostgres)

	userRepoPostgres := userRepoPostgres.New(srv.l, srv.database)
	userUC := userUsecase.New(srv.l, userRepoPostgres, roleUC)
	userH := userHTTP.New(srv.l, userUC)

	sessionRepoPostgres := sessionRepoPostgres.New(srv.l, srv.database)
	sessionUC := sessionUsecase.New(srv.l, sessionRepoPostgres, nil)

	authProd := authProd.NewProducer(srv.l, srv.amqpConn)
	if err := authProd.Run(); err != nil {
		return err
	}
	authUC := authUC.New(srv.l, authProd, srv.encrypter, *srv.redisClient, srv.oauthConfig, scopeManager, smtpUC, userUC, roleUC, sessionUC)
	authH := authHTTP.New(srv.l, authUC)
	sessionUC.SetUserUseCase(userUC)

	uploadRepoPostgres := uploadRepoPostgres.New(srv.l, srv.database)
	uploadUC := uploadUC.New(srv.l, uploadRepoPostgres, srv.cloudinary, userUC)
	uploadH := uploadHTTP.New(srv.l, uploadUC)

	// // Apply locale middleware
	srv.gin.Use(mw.Locale()).Use(mw.Cors())
	api := srv.gin.Group(Api)

	authHTTP.MapAuthRoutes(api.Group("/auth"), authH, mw)
	uploadHTTP.MapUploadRoutes(api.Group("/upload"), uploadH, mw)
	userHTTP.MapUserRoutes(api.Group("/user"), userH, mw)

	return nil
}
