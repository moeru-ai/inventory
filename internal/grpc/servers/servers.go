package servers

import (
	"github.com/moeru-ai/inventory/internal/grpc/servers/inventoryapi/apiserver"
	"go.uber.org/fx"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(apiserver.NewGRPCServer()),
		fx.Provide(apiserver.NewGatewayServer()),
	)
}
