package main

import (
	"context"
	"log"
	"time"

	"github.com/moeru-ai/inventory/internal/configs"
	"github.com/moeru-ai/inventory/internal/cron"
	grpcservers "github.com/moeru-ai/inventory/internal/grpc/servers"
	"github.com/moeru-ai/inventory/internal/grpc/servers/inventoryapi/apiserver"
	grpcservices "github.com/moeru-ai/inventory/internal/grpc/services"
	"github.com/moeru-ai/inventory/internal/libs"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var (
	configFilePath string
	envFilePath    string
)

func main() {
	root := &cobra.Command{
		Use: "api-server",
		RunE: func(cmd *cobra.Command, args []string) error {
			app := fx.New(
				fx.Provide(configs.NewConfig("moeru-ai", "inventory", configFilePath, envFilePath)),
				fx.Options(libs.Modules()),
				fx.Options(grpcservers.Modules()),
				fx.Options(grpcservices.Modules()),
				fx.Options(cron.Modules()),
				fx.Invoke(apiserver.RunGRPCServer()),
				fx.Invoke(apiserver.RunGatewayServer()),
				fx.Invoke(cron.RunCron()),
			)

			app.Run()

			stopCtx, stopCtxCancel := context.WithTimeout(context.Background(), time.Minute*5)
			defer stopCtxCancel()

			if err := app.Stop(stopCtx); err != nil {
				return err
			}

			return nil
		},
	}

	root.Flags().StringVarP(&configFilePath, "config", "c", "", "config file path")
	root.Flags().StringVarP(&envFilePath, "env", "e", "", "env file path")

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
