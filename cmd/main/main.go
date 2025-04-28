package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"gitlab.com/xyxa.gg/backend-mvp-main/api"
	"gitlab.com/xyxa.gg/backend-mvp-main/config"
	"gitlab.com/xyxa.gg/backend-mvp-main/db"
	_ "gitlab.com/xyxa.gg/backend-mvp-main/docs"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/activity"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain/events"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/sockets"
)

// @title XuXaDex API
// @version 1.0
// @description Backend for XuXaDex Platform.

// @BasePath /api/v1
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.InitConfig()

	log := logger.NewApiLogger(cfg)
	log.InitLogger()

	dbClient := db.New(cfg)
	dbClient.MustConnect()

	row := dbClient.GetClient().QueryRow("select now() as t")
	if row.Err() != nil {
		log.Info("failed to get current time")
	}
	var t time.Time
	err := row.Scan(&t)
	if err != nil {
		log.Info("failed to get current time")
	}
	log.Info(fmt.Sprintf("current DB time: %s", t.Format("02.01.2006 15:04:05")))

	socketPool := sockets.NewSocketsPool()

	activityMonitor := activity.NewActivityMonitor(socketPool)
	go activityMonitor.StartMonitoring()

	// blockchainServer := blockchain.NewBlochchainServer(cfg, log)
	// if err = blockchainServer.InitConnection(); err != nil {
	// 	panic(err)
	// }
	// prepareListeners(ctx, blockchainServer, log)
	// blockchainServer.Listen(ctx)

	webServer := web.NewWebServer(cfg, log)
	go webServer.Listen()

	prepareRoutes(ctx, webServer, socketPool, dbClient, log)
	prepareMiddlewares(webServer, cfg)

	log.Infof("web server started on port: %d", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	dbClient.Release()
	// blockchainServer.Release(ctx)
	if err := webServer.Release(ctx); err != nil {
		log.Error(err.Error())
	}
}

func prepareRoutes(
	ctx context.Context,
	webServer *web.WebServer,
	socketPool *sockets.SocketsPool,
	db *db.DBClient,
	log logger.Logger,
) {
	webServer.RegisterRoutes([]api.Controller{
		api.NewSwaggerApi(),
		api.NewProfilerApi(ctx),
		api.NewGameApi(ctx, log, db),
		api.NewTournamentApi(ctx, log, db),
		api.NewSystemApi(ctx, log, db),
		api.NewAuthApi(ctx, log, db),
		api.NewMatchApi(ctx, log, db),
		api.NewWebsocketApi(ctx, socketPool, log),
	})
}

func prepareMiddlewares(webServer *web.WebServer, cfg *config.Config) {
	webServer.RegisterMiddlewares([]middlewares.Middleware{
		middlewares.NewCorsMiddleware(),
		middlewares.NewSessionMiddleware(cfg),
		middlewares.NewRecoverMiddleware(),
		middlewares.NewBodyLimitMiddleware("10M"),
		middlewares.NewSecureMiddleware(),
	})
}

func prepareListeners(ctx context.Context, blockchainServer *blockchain.BlockchainServer, log logger.Logger) {
	blockchainServer.RegisterListeners([]events.EventListener{
		events.NewFooChangeEvent(ctx, log),
	})
}
