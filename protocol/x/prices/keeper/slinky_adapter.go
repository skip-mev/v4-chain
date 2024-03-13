package keeper

import (
	"cosmossdk.io/math"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	slinkytypes "github.com/skip-mev/slinky/pkg/types"
	oracletypes "github.com/skip-mev/slinky/x/oracle/types"
	"strings"
)

/*
 * This package implements the OracleKeeper interface from Slinky's currencypair package
 *
 * It is required in order to convert between x/prices types and the data which slinky stores on chain
 * via Vote Extensions. Using this compatibility layer, we can now use the x/prices keeper as the backing
 * store for converting to and from Slinky's on chain price data.
 */

func (k Keeper) GetCurrencyPairFromID(ctx sdk.Context, id uint64) (cp slinkytypes.CurrencyPair, found bool) {
	// check in the keeper's cache first
	pair, found := k.currencyPairIDCache.GetCurrencyPairFromID(id)
	if found {
		cp, err := slinkytypes.CurrencyPairFromString(pair)
		if err != nil {
			k.Logger(ctx).Error("CurrencyPairFromString", "error", err)
			return cp, false
		}
		return cp, true
	}

	mp, found := k.GetMarketParam(ctx, uint32(id))
	if !found {
		return cp, false
	}
	pair = mp.Pair

	cp, err := MarketPairToCurrencyPair(pair)
	if err != nil {
		k.Logger(ctx).Error("CurrencyPairFromString", "error", err)
		return cp, false
	}
	// cache the result
	k.currencyPairIDCache.AddCurrencyPair(id, cp.String())
	
	return cp, true
}

func (k Keeper) GetIDForCurrencyPair(ctx sdk.Context, cp slinkytypes.CurrencyPair) (uint64, bool) {
	// check in the keeper's cache first
	id, found := k.currencyPairIDCache.GetIDForCurrencyPair(cp.String())
	if found {
		return id, true
	}

	// if not found, iterate through all market params and find the id
	mps := k.GetAllMarketParams(ctx)
	for _, mp := range mps {
		mpCp, err := MarketPairToCurrencyPair(mp.Pair)
		if err != nil {
			k.Logger(ctx).Error("market param pair invalid format", "pair", mp.Pair)
			continue
		}
		if strings.EqualFold(mpCp.String(), cp.String()) {
			// cache the result
			k.currencyPairIDCache.AddCurrencyPair(uint64(mp.Id), cp.String())
			return uint64(mp.Id), true
		}
	}

	return 0, false
}

func (k Keeper) GetPriceForCurrencyPair(ctx sdk.Context, cp slinkytypes.CurrencyPair) (oracletypes.QuotePrice, error) {
	id, found := k.GetIDForCurrencyPair(ctx, cp)
	if !found {
		return oracletypes.QuotePrice{}, fmt.Errorf("currency pair %s not found", cp.String())
	}
	mp, err := k.GetMarketPrice(ctx, uint32(id))
	if err != nil {
		return oracletypes.QuotePrice{}, fmt.Errorf("currency pair %s not found", cp.String())
	}
	return oracletypes.QuotePrice{
		Price: math.NewIntFromUint64(mp.Price),
	}, nil
}

func MarketPairToCurrencyPair(marketPair string) (slinkytypes.CurrencyPair, error) {
	split := strings.Split(marketPair, "-")
	if len(split) != 2 {
		return slinkytypes.CurrencyPair{}, fmt.Errorf("incorrectly formatted CurrencyPair: %s", marketPair)
	}
	cp := slinkytypes.CurrencyPair{
		Base:  strings.ToUpper(split[0]),
		Quote: strings.ToUpper(split[1]),
	}

	return cp, cp.ValidateBasic()
}
