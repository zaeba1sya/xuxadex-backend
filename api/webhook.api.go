package api

import (
	"context"

	"github.com/xuxadex/backend-mvp-main/pkg/logger"
	"github.com/xuxadex/backend-mvp-main/pkg/web/middlewares"
)

type webhookApi struct {
	ctx context.Context
	log logger.Logger
}

func NewWebhookApi(ctx context.Context, log logger.Logger) Controller {
	return &webhookApi{
		ctx: ctx,
		log: log,
	}
}

func (c *webhookApi) GetGroup() string {
	return "/"
}

func (c *webhookApi) GetMiddlewares() []middlewares.Middleware {
	return []middlewares.Middleware{}
}

func (c *webhookApi) GetHandlers() []ControllerHandler {
	return []ControllerHandler{}
}
