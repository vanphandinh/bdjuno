package parse

import (
	parseblocks "github.com/forbole/juno/v3/cmd/parse/blocks"
	parsegenesis "github.com/forbole/juno/v3/cmd/parse/genesis"
	parsetransaction "github.com/forbole/juno/v3/cmd/parse/transactions"
	parse "github.com/forbole/juno/v3/cmd/parse/types"
	"github.com/spf13/cobra"

	parseauth "github.com/forbole/bdjuno/v3/chains/base/parse/auth"
	parsefeegrant "github.com/forbole/bdjuno/v3/chains/base/parse/feegrant"
	parsegov "github.com/forbole/bdjuno/v3/chains/base/parse/gov"
	parsestaking "github.com/forbole/bdjuno/v3/chains/base/parse/staking"
)

// RegisterParseCmd registers the parsing commands
func RegisterParseCmd(cmd *cobra.Command, parseCfg *parse.Config) {
	cmd.AddCommand(
		parseauth.NewAuthCmd(parseCfg),
		parseblocks.NewBlocksCmd(parseCfg),
		parsefeegrant.NewFeegrantCmd(parseCfg),
		parsegenesis.NewGenesisCmd(parseCfg),
		parsegov.NewGovCmd(parseCfg),
		parsestaking.NewStakingCmd(parseCfg),
		parsetransaction.NewTransactionsCmd(parseCfg),
	)
}
