package handlers

import (
	"strconv"
	"time"

	"github.com/danielnapierski/go-alt/clients/federation"
	"github.com/danielnapierski/go-alt/clients/stellartoml"
	"github.com/danielnapierski/go-alt/services/compliance/internal/config"
	"github.com/danielnapierski/go-alt/services/compliance/internal/crypto"
	"github.com/danielnapierski/go-alt/services/compliance/internal/db"
	"github.com/danielnapierski/go-alt/support/http"
)

// RequestHandler implements compliance server request handlers
type RequestHandler struct {
	Config                  *config.Config                 `inject:""`
	Client                  http.SimpleHTTPClientInterface `inject:""`
	Database                db.Database                    `inject:""`
	SignatureSignerVerifier crypto.SignerVerifierInterface `inject:""`
	StellarTomlResolver     stellartoml.ClientInterface    `inject:""`
	FederationResolver      federation.ClientInterface     `inject:""`
	NonceGenerator          NonceGeneratorInterface        `inject:""`
}

type NonceGeneratorInterface interface {
	Generate() string
}

type NonceGenerator struct{}

func (n *NonceGenerator) Generate() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

type TestNonceGenerator struct{}

func (n *TestNonceGenerator) Generate() string {
	return "nonce"
}
