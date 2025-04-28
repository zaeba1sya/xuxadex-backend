package api

import (
	"context"

	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/web/middlewares"
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
