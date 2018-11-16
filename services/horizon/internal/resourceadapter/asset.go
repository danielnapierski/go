package resourceadapter

import (
	"context"

	"github.com/danielnapierski/go-alt/xdr"
	. "github.com/danielnapierski/go-alt/protocols/horizon"

)

func PopulateAsset(ctx context.Context, dest *Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
