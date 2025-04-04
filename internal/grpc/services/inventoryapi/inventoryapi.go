package inventoryapi

import (
	inventoryapiv1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
	commontasksv1 "github.com/moeru-ai/inventory/internal/grpc/services/inventoryapi/v1/commontasksv1"
	grpcpkg "github.com/moeru-ai/inventory/internal/pkg/grpc"
	"go.uber.org/fx"
	"google.golang.org/grpc/reflection"
)

func Modules() fx.Option {
	return fx.Options(
		fx.Provide(NewInventoryAPI()),
		fx.Provide(commontasksv1.NewCommonTasksService()),
	)
}

type NewInventoryAPIParams struct {
	fx.In

	CommonTasks *commontasksv1.CommonTasksService
}

type InventoryAPI struct {
	params *NewInventoryAPIParams
}

func NewInventoryAPI() func(params NewInventoryAPIParams) *InventoryAPI {
	return func(params NewInventoryAPIParams) *InventoryAPI {
		return &InventoryAPI{params: &params}
	}
}

func (c *InventoryAPI) Register(r *grpcpkg.Register) {
	r.RegisterHTTPHandlers([]grpcpkg.HTTPHandler{
		inventoryapiv1.RegisterCommonTasksServiceHandler,
	})

	r.RegisterGrpcService(func(s reflection.GRPCServer) {
		inventoryapiv1.RegisterCommonTasksServiceServer(s, c.params.CommonTasks)
	})
}
