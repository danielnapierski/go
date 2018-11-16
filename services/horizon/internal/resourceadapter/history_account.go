package resourceadapter

import (
	"context"

	"github.com/danielnapierski/go-alt/services/horizon/internal/db2/history"
	. "github.com/danielnapierski/go-alt/protocols/horizon"
)

func PopulateHistoryAccount(ctx context.Context, dest *HistoryAccount, row history.Account) {
	dest.ID = row.Address
	dest.AccountID = row.Address
}
