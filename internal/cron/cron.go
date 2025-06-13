package cron

import (
	"context"

	"github.com/moeru-ai/inventory/internal/cron/models"
	"github.com/robfig/cron/v3"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(NewCron()),
	)
}

type NewCronParams struct {
	fx.In
}

type Cron struct {
	*cron.Cron
}

func NewCron() func(NewCronParams) (*Cron, error) {
	return func(NewCronParams) (*Cron, error) {
		models.EmbeddingVoyageAI()
		models.EmbeddingOpenAI()

		c := cron.New(cron.WithSeconds())

		lo.Must(c.AddFunc("@daily", models.EmbeddingVoyageAI))
		lo.Must(c.AddFunc("@daily", models.EmbeddingOpenAI))

		return &Cron{
			Cron: c,
		}, nil
	}
}

func RunCron() func(*Cron, fx.Lifecycle) error {
	return func(c *Cron, lc fx.Lifecycle) error {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				c.Stop()
				return nil
			},
		})

		go c.Start()

		return nil
	}
}
