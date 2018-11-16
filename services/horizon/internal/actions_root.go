package horizon

import (
	"github.com/danielnapierski/go-alt/protocols/horizon"
	"github.com/danielnapierski/go-alt/services/horizon/internal/ledger"
	"github.com/danielnapierski/go-alt/services/horizon/internal/resourceadapter"
	"github.com/danielnapierski/go-alt/support/render/hal"
)

// RootAction provides a summary of the horizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() {
	var res horizon.Root
	resourceadapter.PopulateRoot(
		action.R.Context(),
		&res,
		ledger.CurrentState(),
		action.App.horizonVersion,
		action.App.coreVersion,
		action.App.networkPassphrase,
		action.App.protocolVersion,
		action.App.config.FriendbotURL,
	)

	hal.Render(action.W, res)
}
