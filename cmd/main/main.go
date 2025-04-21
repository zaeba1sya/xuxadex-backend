package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/xuxadex/backend-mvp-main/api"
	"github.com/xuxadex/backend-mvp-main/config"
	"github.com/xuxadex/backend-mvp-main/db"
	_ "github.com/xuxadex/backend-mvp-main/docs"
	"github.com/xuxadex/backend-mvp-main/pkg/blockchain"
	"github.com/xuxadex/backend-mvp-main/pkg/blockchain/events"
	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/web"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
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
		log.Info("Failed to get current time")
	}
	var t time.Time
	err := row.Scan(&t)
	if err != nil {
		log.Info("Failed to get current time")
	}
	log.Info(fmt.Sprintf("Current DB Time: %s", t.Format("02.01.2006 15:04:05")))

	// blockchainServer := blockchain.NewBlochchainServer(cfg, log)
	// if err = blockchainServer.InitConnection(); err != nil {
	// 	panic(err)
	// }
	// prepareListeners(ctx, blockchainServer, log)
	// blockchainServer.Listen(ctx)

	webServer := web.NewWebServer(cfg, log)
	go webServer.Listen()

	prepareRoutes(ctx, webServer, dbClient, log)
	prepareMiddlewares(webServer)

	log.Infof("Web server started on port: %d", cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	dbClient.Release()
	// blockchainServer.Release(ctx)
	if err := webServer.Release(ctx); err != nil {
		log.Error(err.Error())
	}
}

func prepareRoutes(ctx context.Context, webServer *web.WebServer, db *db.DBClient, log logger.Logger) {
	webServer.RegisterRoutes([]api.Controller{
		api.NewSwaggerApi(),
		api.NewProfilerApi(ctx),
		api.NewGameApi(ctx, log, db),
		api.NewTournamentApi(ctx, log, db),
		api.NewHealthcheckApi(ctx, log, db),
		api.NewAuthApi(ctx, log, db),
		api.NewMatchApi(ctx, log, db),
	})
}

func prepareMiddlewares(webServer *web.WebServer) {
	webServer.RegisterMiddlewares([]middlewares.Middleware{
		middlewares.NewCorsMiddleware(),
		middlewares.NewStaticMiddleware(),
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
