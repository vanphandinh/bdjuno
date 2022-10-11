package parse

import (
	parse "github.com/forbole/juno/v3/cmd/parse/types"
	"github.com/spf13/cobra"

	baseparse "github.com/forbole/bdjuno/v3/chains/base/parse"
)

func RegisterParseCmd(cmd *cobra.Command, parseCfg *parse.Config) {
	baseparse.RegisterParseCmd(cmd, parseCfg)
}
