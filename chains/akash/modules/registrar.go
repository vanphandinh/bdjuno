package modules

import (
	jmodules "github.com/forbole/juno/v3/modules"
	"github.com/forbole/juno/v3/modules/messages"
	"github.com/forbole/juno/v3/modules/registrar"

	basemodules "github.com/forbole/bdjuno/v3/chains/base/modules"
	"github.com/forbole/bdjuno/v3/chains/custom/database"
	"github.com/forbole/bdjuno/v3/chains/custom/modules/deployment"
	"github.com/forbole/bdjuno/v3/chains/custom/modules/market"
	"github.com/forbole/bdjuno/v3/chains/custom/modules/types"
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: basemodules.UniqueAddressesParser(parser),
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	// Build the default modules
	baseRegistrar := basemodules.NewRegistrar(r.parser)
	baseModules := baseRegistrar.BuildModules(ctx)

	cdc := ctx.EncodingConfig.Marshaler
	db := database.Cast(ctx.Database)

	// Build the sources
	sources, err := types.BuildSources(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(err)
	}

	// Build custom modules
	deploymentModule := deployment.NewModule(cdc, db)
	marketModule := market.NewModule(sources.MarketSource, cdc, db)
	// providerModule := provider.NewModule(sources.ProviderSource, cdc, db)

	return append(baseModules, []jmodules.Module{
		deploymentModule,
		marketModule,
	}...)
}
