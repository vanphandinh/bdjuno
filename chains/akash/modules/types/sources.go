package types

import (
	"fmt"
	"os"

	providertypes "github.com/ovrclk/akash/x/provider/types/v1beta2"

	providersource "github.com/forbole/bdjuno/v3/chains/custom/modules/provider/source"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/forbole/juno/v3/node/remote"

	"github.com/forbole/juno/v3/node/local"
	markettypes "github.com/ovrclk/akash/x/market/types/v1beta2"

	nodeconfig "github.com/forbole/juno/v3/node/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
	escrowkeeper "github.com/ovrclk/akash/x/escrow/keeper"
	akashmarket "github.com/ovrclk/akash/x/market"
	akashprovider "github.com/ovrclk/akash/x/provider"

	marketsource "github.com/forbole/bdjuno/v3/chains/custom/modules/market/source"
	localmarketsource "github.com/forbole/bdjuno/v3/chains/custom/modules/market/source/local"
	remotemarketsource "github.com/forbole/bdjuno/v3/chains/custom/modules/market/source/remote"
	localprovidersource "github.com/forbole/bdjuno/v3/chains/custom/modules/provider/source/local"
	remoteprovidersource "github.com/forbole/bdjuno/v3/chains/custom/modules/provider/source/remote"

	akashapp "github.com/ovrclk/akash/app"
)

type Sources struct {
	MarketSource   marketsource.Source
	ProviderSource providersource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, simapp.MakeTestEncodingConfig(), simapp.EmptyAppOptions{},
	)

	// For MarketSource & ProviderSource
	akashApp := akashapp.NewApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, 0, map[int64]bool{},
		cfg.Home, simapp.EmptyAppOptions{},
	)
	escrowKeeper := escrowkeeper.NewKeeper(encodingConfig.Marshaler, sdk.NewKVStoreKey(akashprovider.StoreKey), app.BankKeeper)

	sources := &Sources{
		MarketSource:   localmarketsource.NewSource(source, akashmarket.NewKeeper(encodingConfig.Marshaler, sdk.NewKVStoreKey(akashprovider.StoreKey), akashApp.GetSubspace(markettypes.ModuleName), escrowKeeper).NewQuerier()),
		ProviderSource: localprovidersource.NewSource(source, akashprovider.NewKeeper(encodingConfig.Marshaler, sdk.NewKVStoreKey(akashprovider.StoreKey)).NewQuerier()),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		MarketSource:   remotemarketsource.NewSource(source, markettypes.NewQueryClient(source.GrpcConn)),
		ProviderSource: remoteprovidersource.NewSource(source, providertypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
