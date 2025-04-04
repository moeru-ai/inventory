package services

import (
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"

	"github.com/moeru-ai/inventory/internal/grpc/services/inventoryapi"
	grpcpkg "github.com/moeru-ai/inventory/internal/pkg/grpc"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(NewRegister()),
		fx.Options(inventoryapi.Modules()),
	)
}

type NewRegisterParams struct {
	fx.In

	InventoryAPI *inventoryapi.InventoryAPI
}

func NewRegister() func(params NewRegisterParams) *grpcpkg.Register {
	return func(params NewRegisterParams) *grpcpkg.Register {
		register := grpcpkg.NewRegister()

		params.InventoryAPI.Register(register)

		register.RegisterGrpcService(func(s reflection.GRPCServer) {
			reflection.Register(s)
		})

		return register
	}
}
