package resourceadapter

import (
	"github.com/danielnapierski/go-alt/services/horizon/internal/db2/core"
	. "github.com/danielnapierski/go-alt/protocols/horizon"
)

func PopulateAccountFlags(dest *AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
