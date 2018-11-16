package resourceadapter

import (
	"context"

	"github.com/danielnapierski/go-alt/amount"
	. "github.com/danielnapierski/go-alt/protocols/horizon"
	"github.com/danielnapierski/go-alt/services/horizon/internal/db2/assets"
	"github.com/danielnapierski/go-alt/support/errors"
	"github.com/danielnapierski/go-alt/support/render/hal"
	"github.com/danielnapierski/go-alt/xdr"
)

// PopulateAssetStat fills out the details
//func PopulateAssetStat(
func PopulateAssetStat(
	ctx context.Context,
	res *AssetStat,
	row assets.AssetStatsR,
) (err error) {
	res.Asset.Type = row.Type
	res.Asset.Code = row.Code
	res.Asset.Issuer = row.Issuer
	res.Amount, err = amount.IntStringToAmount(row.Amount)
	if err != nil {
		return errors.Wrap(err, "Invalid amount in PopulateAssetStat")
	}
	res.NumAccounts = row.NumAccounts
	res.Flags = AccountFlags{
		(row.Flags & int8(xdr.AccountFlagsAuthRequiredFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthRevocableFlag)) != 0,
		(row.Flags & int8(xdr.AccountFlagsAuthImmutableFlag)) != 0,
	}
	res.PT = row.SortKey

	res.Links.Toml = hal.NewLink(row.Toml)
	return
}
