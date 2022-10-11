package main

import (
	junocmd "github.com/forbole/juno/v3/cmd"
	junoinitcmd "github.com/forbole/juno/v3/cmd/init"
	junoparsetypes "github.com/forbole/juno/v3/cmd/parse/types"
	junostartcmd "github.com/forbole/juno/v3/cmd/start"

	parsecmd "github.com/forbole/bdjuno/v3/cmd"

	basemigratecmd "github.com/forbole/bdjuno/v3/chains/base/migrate"
	baseconfig "github.com/forbole/bdjuno/v3/chains/base/types/config"

	customconfig "github.com/forbole/bdjuno/v3/chains/custom/config"
	customdb "github.com/forbole/bdjuno/v3/chains/custom/database"
	custommodules "github.com/forbole/bdjuno/v3/chains/custom/modules"
)

func main() {
	initCfg := junoinitcmd.NewConfig().
		WithConfigCreator(baseconfig.Creator)

	parseCfg := junoparsetypes.NewConfig().
		WithDBBuilder(customdb.Builder).
		WithEncodingConfigBuilder(baseconfig.MakeEncodingConfig(customconfig.GetBasicManagers())).
		WithRegistrar(custommodules.NewRegistrar(customconfig.GetAddressesParser()))

	cfg := junocmd.NewConfig("bdjuno").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	// Run the command
	rootCmd := junocmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		junocmd.VersionCmd(),
		junoinitcmd.NewInitCmd(cfg.GetInitConfig()),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		basemigratecmd.NewMigrateCmd(cfg.GetName(), cfg.GetParseConfig()),
		junostartcmd.NewStartCmd(cfg.GetParseConfig()),
	)

	executor := junocmd.PrepareRootCmd(cfg.GetName(), rootCmd)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
