package types

import (
	"context"

	"github.com/dydxprotocol/v4-chain/protocol/x/revshare/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected account keeper used for simulations.
type AccountKeeper interface {
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
}

// BankKeeper defines the expected bank keeper used for simulations.
type BankKeeper interface {
	SpendableCoins(ctx context.Context, addr sdk.AccAddress) sdk.Coins
}

// RevShareKeeper defines the expected revshare keeper used for simulations.
type RevShareKeeper interface {
	CreateNewMarketRevShare(ctx sdk.Context, marketId uint32)
	SetMarketMapperRevShareDetails(
		ctx sdk.Context,
		marketId uint32,
		params types.MarketMapperRevShareDetails,
	)
}
