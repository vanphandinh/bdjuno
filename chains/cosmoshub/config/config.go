package config

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/forbole/juno/v3/modules/messages"

	gaiaapp "github.com/cosmos/gaia/v7/app"

	baseconfig "github.com/forbole/bdjuno/v3/chains/base/types/config"
)

// GetBasicManagers returns the various basic managers that are used to register the encoding to
// support custom messages.
// This should be edited by custom implementations if needed.
func GetBasicManagers() []module.BasicManager {
	return []module.BasicManager{
		gaiaapp.ModuleBasics,
	}
}

// GetAddressesParser returns the messages parser that should be used to get the users involved in
// a specific message.
// This should be edited by custom implementations if needed.
func GetAddressesParser() messages.MessageAddressesParser {
	return baseconfig.GetAddressesParser()
}
