package services

import (
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"

	grpcpkg "github.com/moeru-ai/inventory/internal/pkg/grpc"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(NewRegister()),
	)
}

type NewRegisterParams struct {
	fx.In
}

func NewRegister() func(params NewRegisterParams) *grpcpkg.Register {
	return func(params NewRegisterParams) *grpcpkg.Register {
		register := grpcpkg.NewRegister()

		register.RegisterGrpcService(func(s reflection.GRPCServer) {
			reflection.Register(s)
		})

		return register
	}
}
