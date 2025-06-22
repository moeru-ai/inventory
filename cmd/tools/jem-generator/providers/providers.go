package providers

import (
	"slices"

	"github.com/moeru-ai/inventory/cmd/tools/jem-generator/types"
)

var Providers = []types.Provider{
	providerOpenai,
	providerMinimax,
}

var Models = slices.Concat(
	modelsMinimax,
	modelsOpenai,
)
