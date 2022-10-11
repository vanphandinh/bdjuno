package modules

import (
	"github.com/forbole/juno/v3/modules/messages"
	"github.com/forbole/juno/v3/modules/registrar"

	basemodules "github.com/forbole/bdjuno/v3/chains/base/modules"
)

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser) registrar.Registrar {
	return basemodules.NewRegistrar(parser)
}
