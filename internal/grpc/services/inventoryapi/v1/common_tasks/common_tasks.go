package common_tasksv1

import (
	"context"

	v1 "github.com/moeru-ai/inventory/apis/inventoryapi/v1"
	"github.com/moeru-ai/inventory/internal/pkg/apierrors"
	"github.com/nekomeowww/xo/logger"
	"go.uber.org/fx"
)

type NewCommonTasksServiceParams struct {
	fx.In

	Logger *logger.Logger
}

type CommonTasksService struct {
	v1.UnimplementedCommonTasksServiceServer

	logger *logger.Logger
}

func NewCommonTasksService() func(NewCommonTasksServiceParams) *CommonTasksService {
	return func(params NewCommonTasksServiceParams) *CommonTasksService {
		return &CommonTasksService{
			logger: params.Logger,
		}
	}
}

func (s *CommonTasksService) GetModels(ctx context.Context, in *v1.GetModelsRequest) (*v1.GetModelsResponse, error) {
	return nil, apierrors.NewErrNotFound().WithDetail("not implemented").AsStatus()
}
