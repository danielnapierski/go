package handlers

import (
	"github.com/danielnapierski/go-alt/clients/federation"
	"github.com/danielnapierski/go-alt/clients/horizon"
	"github.com/danielnapierski/go-alt/clients/stellartoml"
	"github.com/danielnapierski/go-alt/services/bridge/internal/config"
	"github.com/danielnapierski/go-alt/services/bridge/internal/db"
	"github.com/danielnapierski/go-alt/services/bridge/internal/listener"
	"github.com/danielnapierski/go-alt/services/bridge/internal/submitter"
	"github.com/danielnapierski/go-alt/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              horizon.ClientInterface                 `inject:""`
	Database             db.Database                             `inject:""`
	StellarTomlResolver  stellartoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}

func (rh *RequestHandler) isAssetAllowed(code string, issuer string) bool {
	for _, asset := range rh.Config.Assets {
		if asset.Code == code && asset.Issuer == issuer {
			return true
		}
	}
	return false
}
