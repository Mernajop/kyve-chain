package keeper

import (
	"cosmossdk.io/errors"
	"cosmossdk.io/math"
	"fmt"

	"github.com/KYVENetwork/chain/x/funders/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorsTypes "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k Keeper) CreateFundingState(ctx sdk.Context, poolId uint64) {
	fundingState := types.FundingState{
		PoolId:                poolId,
		ActiveFunderAddresses: []string{},
	}
	k.SetFundingState(ctx, &fundingState)
}

func (k Keeper) GetTotalActiveFunding(ctx sdk.Context, poolId uint64) (amounts sdk.Coins) {
	state, found := k.GetFundingState(ctx, poolId)
	if !found {
		return sdk.NewCoins()
	}
	for _, address := range state.ActiveFunderAddresses {
		funding, _ := k.GetFunding(ctx, address, poolId)
		amounts = amounts.Add(funding.Amounts...)
	}
	return
}

// ChargeFundersOfPool charges all funders of a pool with their amount_per_bundle
// If the amount is lower than the amount_per_bundle,
// the max amount is charged and the funder is removed from the active funders list.
// The amount is transferred from the funders to the pool module account where it can be paid out.
// If there are no more active funders, an event is emitted.
func (k Keeper) ChargeFundersOfPool(ctx sdk.Context, poolId uint64) (payouts sdk.Coins, err error) {
	// Get funding state for pool
	fundingState, found := k.GetFundingState(ctx, poolId)
	if !found {
		return sdk.NewCoins(), errors.Wrapf(errorsTypes.ErrNotFound, types.ErrFundingStateDoesNotExist.Error(), poolId)
	}

	// If there are no active fundings we immediately return
	activeFundings := k.GetActiveFundings(ctx, fundingState)
	if len(activeFundings) == 0 {
		return sdk.NewCoins(), nil
	}

	// This is the amount every funding will be charged
	for _, funding := range activeFundings {
		payouts = payouts.Add(funding.ChargeOneBundle()...)
		if funding.Amounts.IsZero() {
			fundingState.SetInactive(&funding)
		}
		k.SetFunding(ctx, &funding)
	}

	// Save funding state
	k.SetFundingState(ctx, &fundingState)

	// Emit a pool out of funds event if there are no more active funders
	if len(fundingState.ActiveFunderAddresses) == 0 {
		_ = ctx.EventManager().EmitTypedEvent(&types.EventPoolOutOfFunds{
			PoolId: poolId,
		})
	}

	// Move funds to pool module account
	if !payouts.IsZero() {
		if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, pooltypes.ModuleName, payouts); err != nil {
			return sdk.NewCoins(), err
		}
	}

	return
}

// GetLowestFunding returns the funding with the lowest amount
// Precondition: len(fundings) > 0
func (k Keeper) GetLowestFunding(fundings []types.Funding, whitelist []*types.WhitelistCoinEntry) (lowestFunding *types.Funding, err error) {
	if len(fundings) == 0 {
		return nil, fmt.Errorf("no active fundings")
	}

	lowestFundingIndex := 0
	for i := range fundings {
		if fundings[i].GetScore(whitelist) < fundings[lowestFundingIndex].GetScore(whitelist) {
			lowestFundingIndex = i
		}
	}
	return &fundings[lowestFundingIndex], nil
}

// ensureParamsCompatibility checks compatibility of the provided funding with the pool params.
// i.e.
// - coin is in whitelist
// - minimum funding per bundle
// - minimum funding amount
// - minimum funding multiple
func (k Keeper) ensureParamsCompatibility(ctx sdk.Context, funding types.Funding) error {
	params := k.GetParams(ctx)

	minFundingAmounts := sdk.NewCoins()
	minFundingAmountsPerBundle := sdk.NewCoins()

	for _, entry := range params.CoinWhitelist {
		minFundingAmounts = minFundingAmounts.Add(sdk.NewInt64Coin(entry.CoinDenom, int64(entry.MinFundingAmount)))
		minFundingAmountsPerBundle = minFundingAmountsPerBundle.Add(sdk.NewInt64Coin(entry.CoinDenom, int64(entry.MinFundingAmountPerBundle)))
	}

	// throw error if a coin in amounts is not in the whitelist
	if !funding.Amounts.DenomsSubsetOf(minFundingAmounts) {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrCoinNotWhitelisted.Error())
	}

	// throw error if a coin in amounts per bundle is not in the whitelist
	if !funding.AmountsPerBundle.DenomsSubsetOf(minFundingAmountsPerBundle) {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrCoinNotWhitelisted.Error())
	}

	// throw error if a coin is less than the minimum funding amount
	if minFundingAmounts.IsAnyGT(funding.Amounts) {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrMinFundingAmount.Error())
	}

	// throw error if a coin is less than the minimum funding amount per bundle
	if minFundingAmountsPerBundle.IsAnyGT(funding.AmountsPerBundle) {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrMinAmountPerBundle.Error())
	}

	// throw error if a coin can not fulfill the funding multiple threshold
	if funding.AmountsPerBundle.MulInt(math.NewInt(int64(params.MinFundingMultiple))).IsAnyGT(funding.Amounts) {
		return errors.Wrapf(errorsTypes.ErrInvalidRequest, types.ErrMinFundingMultiple.Error())
	}

	return nil
}

// ensureFreeSlot makes sure that a funder can add funding to a given pool.
// If this is not possible an appropriate error is returned.
// A pool has a fixed amount of funding-slots. If there are still free slots
// a funder can just join (even with the smallest funding possible).
// If all slots are taken, it checks if the new funding has more funds
// than the current lowest funding in that pool.
// If so, the lowest funding gets removed from the pool, so that the
// new funding can be added.
// CONTRACT: no KV Writing on newFunding and fundingState
func (k Keeper) ensureFreeSlot(ctx sdk.Context, newFunding *types.Funding, fundingState *types.FundingState) error {
	activeFundings := k.GetActiveFundings(ctx, *fundingState)
	// check if slots are still available
	if len(activeFundings) < types.MaxFunders {
		return nil
	}

	params := k.GetParams(ctx)

	lowestFunding, err := k.GetLowestFunding(activeFundings, params.CoinWhitelist)
	if err != nil {
		return err
	}

	if lowestFunding.FunderAddress == newFunding.FunderAddress {
		// Funder already has a funding slot
		return nil
	}

	// Check if lowest funding is lower than new funding based on amount (amount per bundle is ignored)
	if newFunding.GetScore(params.CoinWhitelist) < lowestFunding.GetScore(params.CoinWhitelist) {
		return errors.Wrapf(errorsTypes.ErrLogic, types.ErrFundsTooLow.Error(), lowestFunding.GetScore(params.CoinWhitelist))
	}

	// Defund lowest funder
	recipient := sdk.MustAccAddressFromBech32(lowestFunding.FunderAddress)
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, recipient, lowestFunding.Amounts); err != nil {
		return err
	}

	lowestFunding.Amounts = lowestFunding.Amounts.Sub(lowestFunding.Amounts...)
	fundingState.SetInactive(lowestFunding)
	k.SetFunding(ctx, lowestFunding)

	// Emit a defund event.
	_ = ctx.EventManager().EmitTypedEvent(&types.EventDefundPool{
		PoolId:  fundingState.PoolId,
		Address: lowestFunding.FunderAddress,
		Amounts: lowestFunding.Amounts,
	})

	return nil
}
