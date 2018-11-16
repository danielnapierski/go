package resourceadapter

import (
	"context"

	"github.com/danielnapierski/go-alt/services/horizon/internal/httpx"
	"github.com/danielnapierski/go-alt/services/horizon/internal/txsub"
	. "github.com/danielnapierski/go-alt/protocols/horizon"
	"github.com/danielnapierski/go-alt/support/render/hal"
)

// Populate fills out the details
func PopulateTransactionSuccess(ctx context.Context, dest *TransactionSuccess, result txsub.Result) {
	dest.Hash = result.Hash
	dest.Ledger = result.LedgerSequence
	dest.Env = result.EnvelopeXDR
	dest.Result = result.ResultXDR
	dest.Meta = result.ResultMetaXDR

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Transaction = lb.Link("/transactions", result.Hash)
	return
}
