// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	indexer_manager "github.com/dydxprotocol/v4-chain/protocol/indexer/indexer_manager"
	clobtypes "github.com/dydxprotocol/v4-chain/protocol/x/clob/types"

	log "cosmossdk.io/log"

	mock "github.com/stretchr/testify/mock"

	subaccountstypes "github.com/dydxprotocol/v4-chain/protocol/x/subaccounts/types"

	time "time"

	types "github.com/cosmos/cosmos-sdk/types"
)

// ClobKeeper is an autogenerated mock type for the ClobKeeper type
type ClobKeeper struct {
	mock.Mock
}

// AddOrderToOrderbookCollatCheck provides a mock function with given fields: ctx, clobPairId, subaccountOpenOrders
func (_m *ClobKeeper) AddOrderToOrderbookCollatCheck(ctx types.Context, clobPairId clobtypes.ClobPairId, subaccountOpenOrders map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) (bool, map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult) {
	ret := _m.Called(ctx, clobPairId, subaccountOpenOrders)

	if len(ret) == 0 {
		panic("no return value specified for AddOrderToOrderbookCollatCheck")
	}

	var r0 bool
	var r1 map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) (bool, map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult)); ok {
		return rf(ctx, clobPairId, subaccountOpenOrders)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) bool); ok {
		r0 = rf(ctx, clobPairId, subaccountOpenOrders)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPairId, map[subaccountstypes.SubaccountId][]clobtypes.PendingOpenOrder) map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, clobPairId, subaccountOpenOrders)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[subaccountstypes.SubaccountId]subaccountstypes.UpdateResult)
		}
	}

	return r0, r1
}

// BatchCancelShortTermOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) BatchCancelShortTermOrder(ctx types.Context, msg *clobtypes.MsgBatchCancel) ([]uint32, []uint32, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for BatchCancelShortTermOrder")
	}

	var r0 []uint32
	var r1 []uint32
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgBatchCancel) ([]uint32, []uint32, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgBatchCancel) []uint32); ok {
		r0 = rf(ctx, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MsgBatchCancel) []uint32); ok {
		r1 = rf(ctx, msg)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]uint32)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MsgBatchCancel) error); ok {
		r2 = rf(ctx, msg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CancelShortTermOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) CancelShortTermOrder(ctx types.Context, msg *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for CancelShortTermOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CancelStatefulOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) CancelStatefulOrder(ctx types.Context, msg *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for CancelStatefulOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConvertFillablePriceToSubticks provides a mock function with given fields: ctx, fillablePrice, isLiquidatingLong, clobPair
func (_m *ClobKeeper) ConvertFillablePriceToSubticks(ctx types.Context, fillablePrice *big.Rat, isLiquidatingLong bool, clobPair clobtypes.ClobPair) clobtypes.Subticks {
	ret := _m.Called(ctx, fillablePrice, isLiquidatingLong, clobPair)

	if len(ret) == 0 {
		panic("no return value specified for ConvertFillablePriceToSubticks")
	}

	var r0 clobtypes.Subticks
	if rf, ok := ret.Get(0).(func(types.Context, *big.Rat, bool, clobtypes.ClobPair) clobtypes.Subticks); ok {
		r0 = rf(ctx, fillablePrice, isLiquidatingLong, clobPair)
	} else {
		r0 = ret.Get(0).(clobtypes.Subticks)
	}

	return r0
}

// CreatePerpetualClobPair provides a mock function with given fields: ctx, clobPairId, perpetualId, stepSizeInBaseQuantums, quantumConversionExponent, subticksPerTick, status
func (_m *ClobKeeper) CreatePerpetualClobPair(ctx types.Context, clobPairId uint32, perpetualId uint32, stepSizeInBaseQuantums subaccountstypes.BaseQuantums, quantumConversionExponent int32, subticksPerTick uint32, status clobtypes.ClobPair_Status) (clobtypes.ClobPair, error) {
	ret := _m.Called(ctx, clobPairId, perpetualId, stepSizeInBaseQuantums, quantumConversionExponent, subticksPerTick, status)

	if len(ret) == 0 {
		panic("no return value specified for CreatePerpetualClobPair")
	}

	var r0 clobtypes.ClobPair
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, uint32, uint32, subaccountstypes.BaseQuantums, int32, uint32, clobtypes.ClobPair_Status) (clobtypes.ClobPair, error)); ok {
		return rf(ctx, clobPairId, perpetualId, stepSizeInBaseQuantums, quantumConversionExponent, subticksPerTick, status)
	}
	if rf, ok := ret.Get(0).(func(types.Context, uint32, uint32, subaccountstypes.BaseQuantums, int32, uint32, clobtypes.ClobPair_Status) clobtypes.ClobPair); ok {
		r0 = rf(ctx, clobPairId, perpetualId, stepSizeInBaseQuantums, quantumConversionExponent, subticksPerTick, status)
	} else {
		r0 = ret.Get(0).(clobtypes.ClobPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, uint32, uint32, subaccountstypes.BaseQuantums, int32, uint32, clobtypes.ClobPair_Status) error); ok {
		r1 = rf(ctx, clobPairId, perpetualId, stepSizeInBaseQuantums, quantumConversionExponent, subticksPerTick, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteLongTermOrderPlacement provides a mock function with given fields: ctx, orderId
func (_m *ClobKeeper) DeleteLongTermOrderPlacement(ctx types.Context, orderId clobtypes.OrderId) {
	_m.Called(ctx, orderId)
}

// GetAllClobPairs provides a mock function with given fields: ctx
func (_m *ClobKeeper) GetAllClobPairs(ctx types.Context) []clobtypes.ClobPair {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllClobPairs")
	}

	var r0 []clobtypes.ClobPair
	if rf, ok := ret.Get(0).(func(types.Context) []clobtypes.ClobPair); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.ClobPair)
		}
	}

	return r0
}

// GetBankruptcyPriceInQuoteQuantums provides a mock function with given fields: ctx, subaccountId, perpetualId, deltaQuantums
func (_m *ClobKeeper) GetBankruptcyPriceInQuoteQuantums(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32, deltaQuantums *big.Int) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId, deltaQuantums)

	if len(ret) == 0 {
		panic("no return value specified for GetBankruptcyPriceInQuoteQuantums")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId, deltaQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId, deltaQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId, deltaQuantums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetClobPair provides a mock function with given fields: ctx, id
func (_m *ClobKeeper) GetClobPair(ctx types.Context, id clobtypes.ClobPairId) (clobtypes.ClobPair, bool) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetClobPair")
	}

	var r0 clobtypes.ClobPair
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId) (clobtypes.ClobPair, bool)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPairId) clobtypes.ClobPair); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(clobtypes.ClobPair)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.ClobPairId) bool); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetFillablePrice provides a mock function with given fields: ctx, subaccountId, perpetualId, deltaQuantums
func (_m *ClobKeeper) GetFillablePrice(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32, deltaQuantums *big.Int) (*big.Rat, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId, deltaQuantums)

	if len(ret) == 0 {
		panic("no return value specified for GetFillablePrice")
	}

	var r0 *big.Rat
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) (*big.Rat, error)); ok {
		return rf(ctx, subaccountId, perpetualId, deltaQuantums)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) *big.Rat); ok {
		r0 = rf(ctx, subaccountId, perpetualId, deltaQuantums)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Rat)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, *big.Int) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId, deltaQuantums)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIndexerEventManager provides a mock function with given fields:
func (_m *ClobKeeper) GetIndexerEventManager() indexer_manager.IndexerEventManager {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetIndexerEventManager")
	}

	var r0 indexer_manager.IndexerEventManager
	if rf, ok := ret.Get(0).(func() indexer_manager.IndexerEventManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(indexer_manager.IndexerEventManager)
		}
	}

	return r0
}

// GetInsuranceFundBalance provides a mock function with given fields: ctx, perpetualId
func (_m *ClobKeeper) GetInsuranceFundBalance(ctx types.Context, perpetualId uint32) *big.Int {
	ret := _m.Called(ctx, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for GetInsuranceFundBalance")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, uint32) *big.Int); ok {
		r0 = rf(ctx, perpetualId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetLiquidationInsuranceFundDelta provides a mock function with given fields: ctx, subaccountId, perpetualId, isBuy, fillAmount, subticks
func (_m *ClobKeeper) GetLiquidationInsuranceFundDelta(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32, isBuy bool, fillAmount uint64, subticks clobtypes.Subticks) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId, isBuy, fillAmount, subticks)

	if len(ret) == 0 {
		panic("no return value specified for GetLiquidationInsuranceFundDelta")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, bool, uint64, clobtypes.Subticks) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId, isBuy, fillAmount, subticks)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32, bool, uint64, clobtypes.Subticks) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId, isBuy, fillAmount, subticks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32, bool, uint64, clobtypes.Subticks) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId, isBuy, fillAmount, subticks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLiquidationsConfig provides a mock function with given fields: ctx
func (_m *ClobKeeper) GetLiquidationsConfig(ctx types.Context) clobtypes.LiquidationsConfig {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLiquidationsConfig")
	}

	var r0 clobtypes.LiquidationsConfig
	if rf, ok := ret.Get(0).(func(types.Context) clobtypes.LiquidationsConfig); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(clobtypes.LiquidationsConfig)
	}

	return r0
}

// GetLongTermOrderPlacement provides a mock function with given fields: ctx, orderId
func (_m *ClobKeeper) GetLongTermOrderPlacement(ctx types.Context, orderId clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool) {
	ret := _m.Called(ctx, orderId)

	if len(ret) == 0 {
		panic("no return value specified for GetLongTermOrderPlacement")
	}

	var r0 clobtypes.LongTermOrderPlacement
	var r1 bool
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) (clobtypes.LongTermOrderPlacement, bool)); ok {
		return rf(ctx, orderId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.OrderId) clobtypes.LongTermOrderPlacement); ok {
		r0 = rf(ctx, orderId)
	} else {
		r0 = ret.Get(0).(clobtypes.LongTermOrderPlacement)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.OrderId) bool); ok {
		r1 = rf(ctx, orderId)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetMaxAndMinPositionNotionalLiquidatable provides a mock function with given fields: ctx, positionToLiquidate
func (_m *ClobKeeper) GetMaxAndMinPositionNotionalLiquidatable(ctx types.Context, positionToLiquidate *subaccountstypes.PerpetualPosition) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx, positionToLiquidate)

	if len(ret) == 0 {
		panic("no return value specified for GetMaxAndMinPositionNotionalLiquidatable")
	}

	var r0 *big.Int
	var r1 *big.Int
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, *subaccountstypes.PerpetualPosition) (*big.Int, *big.Int, error)); ok {
		return rf(ctx, positionToLiquidate)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *subaccountstypes.PerpetualPosition) *big.Int); ok {
		r0 = rf(ctx, positionToLiquidate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, *subaccountstypes.PerpetualPosition) *big.Int); ok {
		r1 = rf(ctx, positionToLiquidate)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	if rf, ok := ret.Get(2).(func(types.Context, *subaccountstypes.PerpetualPosition) error); ok {
		r2 = rf(ctx, positionToLiquidate)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPerpetualPositionToLiquidate provides a mock function with given fields: ctx, subaccountId
func (_m *ClobKeeper) GetPerpetualPositionToLiquidate(ctx types.Context, subaccountId subaccountstypes.SubaccountId) (uint32, error) {
	ret := _m.Called(ctx, subaccountId)

	if len(ret) == 0 {
		panic("no return value specified for GetPerpetualPositionToLiquidate")
	}

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) (uint32, error)); ok {
		return rf(ctx, subaccountId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) uint32); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId) error); ok {
		r1 = rf(ctx, subaccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProcessProposerMatchesEvents provides a mock function with given fields: ctx
func (_m *ClobKeeper) GetProcessProposerMatchesEvents(ctx types.Context) clobtypes.ProcessProposerMatchesEvents {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetProcessProposerMatchesEvents")
	}

	var r0 clobtypes.ProcessProposerMatchesEvents
	if rf, ok := ret.Get(0).(func(types.Context) clobtypes.ProcessProposerMatchesEvents); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(clobtypes.ProcessProposerMatchesEvents)
	}

	return r0
}

// GetStatePosition provides a mock function with given fields: ctx, subaccountId, clobPairId
func (_m *ClobKeeper) GetStatePosition(ctx types.Context, subaccountId subaccountstypes.SubaccountId, clobPairId clobtypes.ClobPairId) *big.Int {
	ret := _m.Called(ctx, subaccountId, clobPairId)

	if len(ret) == 0 {
		panic("no return value specified for GetStatePosition")
	}

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, clobtypes.ClobPairId) *big.Int); ok {
		r0 = rf(ctx, subaccountId, clobPairId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetStatefulOrdersTimeSlice provides a mock function with given fields: ctx, goodTilBlockTime
func (_m *ClobKeeper) GetStatefulOrdersTimeSlice(ctx types.Context, goodTilBlockTime time.Time) []clobtypes.OrderId {
	ret := _m.Called(ctx, goodTilBlockTime)

	if len(ret) == 0 {
		panic("no return value specified for GetStatefulOrdersTimeSlice")
	}

	var r0 []clobtypes.OrderId
	if rf, ok := ret.Get(0).(func(types.Context, time.Time) []clobtypes.OrderId); ok {
		r0 = rf(ctx, goodTilBlockTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.OrderId)
		}
	}

	return r0
}

// GetSubaccountLiquidationInfo provides a mock function with given fields: ctx, subaccountId
func (_m *ClobKeeper) GetSubaccountLiquidationInfo(ctx types.Context, subaccountId subaccountstypes.SubaccountId) clobtypes.SubaccountLiquidationInfo {
	ret := _m.Called(ctx, subaccountId)

	if len(ret) == 0 {
		panic("no return value specified for GetSubaccountLiquidationInfo")
	}

	var r0 clobtypes.SubaccountLiquidationInfo
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) clobtypes.SubaccountLiquidationInfo); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(clobtypes.SubaccountLiquidationInfo)
	}

	return r0
}

// GetSubaccountMaxInsuranceLost provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *ClobKeeper) GetSubaccountMaxInsuranceLost(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for GetSubaccountMaxInsuranceLost")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSubaccountMaxNotionalLiquidatable provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *ClobKeeper) GetSubaccountMaxNotionalLiquidatable(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for GetSubaccountMaxNotionalLiquidatable")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleMsgCancelOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) HandleMsgCancelOrder(ctx types.Context, msg *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for HandleMsgCancelOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HandleMsgPlaceOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) HandleMsgPlaceOrder(ctx types.Context, msg *clobtypes.MsgPlaceOrder) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for HandleMsgPlaceOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasAuthority provides a mock function with given fields: authority
func (_m *ClobKeeper) HasAuthority(authority string) bool {
	ret := _m.Called(authority)

	if len(ret) == 0 {
		panic("no return value specified for HasAuthority")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(authority)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// InitializeBlockRateLimit provides a mock function with given fields: ctx, config
func (_m *ClobKeeper) InitializeBlockRateLimit(ctx types.Context, config clobtypes.BlockRateLimitConfiguration) error {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for InitializeBlockRateLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.BlockRateLimitConfiguration) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitializeEquityTierLimit provides a mock function with given fields: ctx, config
func (_m *ClobKeeper) InitializeEquityTierLimit(ctx types.Context, config clobtypes.EquityTierLimitConfiguration) error {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for InitializeEquityTierLimit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.EquityTierLimitConfiguration) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitializeNewGrpcStreams provides a mock function with given fields: ctx
func (_m *ClobKeeper) InitializeNewGrpcStreams(ctx types.Context) {
	_m.Called(ctx)
}

// IsLiquidatable provides a mock function with given fields: ctx, subaccountId
func (_m *ClobKeeper) IsLiquidatable(ctx types.Context, subaccountId subaccountstypes.SubaccountId) (bool, error) {
	ret := _m.Called(ctx, subaccountId)

	if len(ret) == 0 {
		panic("no return value specified for IsLiquidatable")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) (bool, error)); ok {
		return rf(ctx, subaccountId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) bool); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId) error); ok {
		r1 = rf(ctx, subaccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Logger provides a mock function with given fields: ctx
func (_m *ClobKeeper) Logger(ctx types.Context) log.Logger {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Logger")
	}

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func(types.Context) log.Logger); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// MaybeDeleverageSubaccount provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *ClobKeeper) MaybeDeleverageSubaccount(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) (*big.Int, error) {
	ret := _m.Called(ctx, subaccountId, perpetualId)

	if len(ret) == 0 {
		panic("no return value specified for MaybeDeleverageSubaccount")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) (*big.Int, error)); ok {
		return rf(ctx, subaccountId, perpetualId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId, uint32) *big.Int); ok {
		r0 = rf(ctx, subaccountId, perpetualId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId, uint32) error); ok {
		r1 = rf(ctx, subaccountId, perpetualId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaybeGetLiquidationOrder provides a mock function with given fields: ctx, subaccountId
func (_m *ClobKeeper) MaybeGetLiquidationOrder(ctx types.Context, subaccountId subaccountstypes.SubaccountId) (*clobtypes.LiquidationOrder, error) {
	ret := _m.Called(ctx, subaccountId)

	if len(ret) == 0 {
		panic("no return value specified for MaybeGetLiquidationOrder")
	}

	var r0 *clobtypes.LiquidationOrder
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) (*clobtypes.LiquidationOrder, error)); ok {
		return rf(ctx, subaccountId)
	}
	if rf, ok := ret.Get(0).(func(types.Context, subaccountstypes.SubaccountId) *clobtypes.LiquidationOrder); ok {
		r0 = rf(ctx, subaccountId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*clobtypes.LiquidationOrder)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Context, subaccountstypes.SubaccountId) error); ok {
		r1 = rf(ctx, subaccountId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MustAddOrderToStatefulOrdersTimeSlice provides a mock function with given fields: ctx, goodTilBlockTime, orderId
func (_m *ClobKeeper) MustAddOrderToStatefulOrdersTimeSlice(ctx types.Context, goodTilBlockTime time.Time, orderId clobtypes.OrderId) {
	_m.Called(ctx, goodTilBlockTime, orderId)
}

// MustRemoveStatefulOrder provides a mock function with given fields: ctx, orderId
func (_m *ClobKeeper) MustRemoveStatefulOrder(ctx types.Context, orderId clobtypes.OrderId) {
	_m.Called(ctx, orderId)
}

// MustSetProcessProposerMatchesEvents provides a mock function with given fields: ctx, processProposerMatchesEvents
func (_m *ClobKeeper) MustSetProcessProposerMatchesEvents(ctx types.Context, processProposerMatchesEvents clobtypes.ProcessProposerMatchesEvents) {
	_m.Called(ctx, processProposerMatchesEvents)
}

// MustUpdateSubaccountPerpetualLiquidated provides a mock function with given fields: ctx, subaccountId, perpetualId
func (_m *ClobKeeper) MustUpdateSubaccountPerpetualLiquidated(ctx types.Context, subaccountId subaccountstypes.SubaccountId, perpetualId uint32) {
	_m.Called(ctx, subaccountId, perpetualId)
}

// PerformOrderCancellationStatefulValidation provides a mock function with given fields: ctx, msgCancelOrder, blockHeight
func (_m *ClobKeeper) PerformOrderCancellationStatefulValidation(ctx types.Context, msgCancelOrder *clobtypes.MsgCancelOrder, blockHeight uint32) error {
	ret := _m.Called(ctx, msgCancelOrder, blockHeight)

	if len(ret) == 0 {
		panic("no return value specified for PerformOrderCancellationStatefulValidation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder, uint32) error); ok {
		r0 = rf(ctx, msgCancelOrder, blockHeight)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PerformStatefulOrderValidation provides a mock function with given fields: ctx, order, blockHeight, isPreexistingStatefulOrder
func (_m *ClobKeeper) PerformStatefulOrderValidation(ctx types.Context, order *clobtypes.Order, blockHeight uint32, isPreexistingStatefulOrder bool) error {
	ret := _m.Called(ctx, order, blockHeight, isPreexistingStatefulOrder)

	if len(ret) == 0 {
		panic("no return value specified for PerformStatefulOrderValidation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.Order, uint32, bool) error); ok {
		r0 = rf(ctx, order, blockHeight, isPreexistingStatefulOrder)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PlacePerpetualLiquidation provides a mock function with given fields: ctx, liquidationOrder
func (_m *ClobKeeper) PlacePerpetualLiquidation(ctx types.Context, liquidationOrder clobtypes.LiquidationOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, error) {
	ret := _m.Called(ctx, liquidationOrder)

	if len(ret) == 0 {
		panic("no return value specified for PlacePerpetualLiquidation")
	}

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.LiquidationOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, error)); ok {
		return rf(ctx, liquidationOrder)
	}
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.LiquidationOrder) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, liquidationOrder)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, clobtypes.LiquidationOrder) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, liquidationOrder)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, clobtypes.LiquidationOrder) error); ok {
		r2 = rf(ctx, liquidationOrder)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PlaceShortTermOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) PlaceShortTermOrder(ctx types.Context, msg *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for PlaceShortTermOrder")
	}

	var r0 subaccountstypes.BaseQuantums
	var r1 clobtypes.OrderStatus
	var r2 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) (subaccountstypes.BaseQuantums, clobtypes.OrderStatus, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) subaccountstypes.BaseQuantums); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(subaccountstypes.BaseQuantums)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MsgPlaceOrder) clobtypes.OrderStatus); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Get(1).(clobtypes.OrderStatus)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r2 = rf(ctx, msg)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PlaceStatefulOrder provides a mock function with given fields: ctx, msg
func (_m *ClobKeeper) PlaceStatefulOrder(ctx types.Context, msg *clobtypes.MsgPlaceOrder) error {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for PlaceStatefulOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessProposerOperations provides a mock function with given fields: ctx, operations
func (_m *ClobKeeper) ProcessProposerOperations(ctx types.Context, operations []clobtypes.OperationRaw) error {
	ret := _m.Called(ctx, operations)

	if len(ret) == 0 {
		panic("no return value specified for ProcessProposerOperations")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, []clobtypes.OperationRaw) error); ok {
		r0 = rf(ctx, operations)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProcessSingleMatch provides a mock function with given fields: ctx, matchWithOrders
func (_m *ClobKeeper) ProcessSingleMatch(ctx types.Context, matchWithOrders *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, *clobtypes.OffchainUpdates, error) {
	ret := _m.Called(ctx, matchWithOrders)

	if len(ret) == 0 {
		panic("no return value specified for ProcessSingleMatch")
	}

	var r0 bool
	var r1 subaccountstypes.UpdateResult
	var r2 subaccountstypes.UpdateResult
	var r3 *clobtypes.OffchainUpdates
	var r4 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) (bool, subaccountstypes.UpdateResult, subaccountstypes.UpdateResult, *clobtypes.OffchainUpdates, error)); ok {
		return rf(ctx, matchWithOrders)
	}
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MatchWithOrders) bool); ok {
		r0 = rf(ctx, matchWithOrders)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r1 = rf(ctx, matchWithOrders)
	} else {
		r1 = ret.Get(1).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(2).(func(types.Context, *clobtypes.MatchWithOrders) subaccountstypes.UpdateResult); ok {
		r2 = rf(ctx, matchWithOrders)
	} else {
		r2 = ret.Get(2).(subaccountstypes.UpdateResult)
	}

	if rf, ok := ret.Get(3).(func(types.Context, *clobtypes.MatchWithOrders) *clobtypes.OffchainUpdates); ok {
		r3 = rf(ctx, matchWithOrders)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*clobtypes.OffchainUpdates)
		}
	}

	if rf, ok := ret.Get(4).(func(types.Context, *clobtypes.MatchWithOrders) error); ok {
		r4 = rf(ctx, matchWithOrders)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// PruneStateFillAmountsForShortTermOrders provides a mock function with given fields: ctx
func (_m *ClobKeeper) PruneStateFillAmountsForShortTermOrders(ctx types.Context) {
	_m.Called(ctx)
}

// RateLimitCancelOrder provides a mock function with given fields: ctx, order
func (_m *ClobKeeper) RateLimitCancelOrder(ctx types.Context, order *clobtypes.MsgCancelOrder) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for RateLimitCancelOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgCancelOrder) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RateLimitPlaceOrder provides a mock function with given fields: ctx, order
func (_m *ClobKeeper) RateLimitPlaceOrder(ctx types.Context, order *clobtypes.MsgPlaceOrder) error {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for RateLimitPlaceOrder")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, *clobtypes.MsgPlaceOrder) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveClobPair provides a mock function with given fields: ctx, id
func (_m *ClobKeeper) RemoveClobPair(ctx types.Context, id clobtypes.ClobPairId) {
	_m.Called(ctx, id)
}

// RemoveExpiredStatefulOrdersTimeSlices provides a mock function with given fields: ctx, blockTime
func (_m *ClobKeeper) RemoveExpiredStatefulOrdersTimeSlices(ctx types.Context, blockTime time.Time) []clobtypes.OrderId {
	ret := _m.Called(ctx, blockTime)

	if len(ret) == 0 {
		panic("no return value specified for RemoveExpiredStatefulOrdersTimeSlices")
	}

	var r0 []clobtypes.OrderId
	if rf, ok := ret.Get(0).(func(types.Context, time.Time) []clobtypes.OrderId); ok {
		r0 = rf(ctx, blockTime)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]clobtypes.OrderId)
		}
	}

	return r0
}

// RemoveOrderFillAmount provides a mock function with given fields: ctx, orderId
func (_m *ClobKeeper) RemoveOrderFillAmount(ctx types.Context, orderId clobtypes.OrderId) {
	_m.Called(ctx, orderId)
}

// SetLongTermOrderPlacement provides a mock function with given fields: ctx, order, blockHeight
func (_m *ClobKeeper) SetLongTermOrderPlacement(ctx types.Context, order clobtypes.Order, blockHeight uint32) {
	_m.Called(ctx, order, blockHeight)
}

// UpdateClobPair provides a mock function with given fields: ctx, clobPair
func (_m *ClobKeeper) UpdateClobPair(ctx types.Context, clobPair clobtypes.ClobPair) error {
	ret := _m.Called(ctx, clobPair)

	if len(ret) == 0 {
		panic("no return value specified for UpdateClobPair")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.ClobPair) error); ok {
		r0 = rf(ctx, clobPair)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLiquidationsConfig provides a mock function with given fields: ctx, config
func (_m *ClobKeeper) UpdateLiquidationsConfig(ctx types.Context, config clobtypes.LiquidationsConfig) error {
	ret := _m.Called(ctx, config)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLiquidationsConfig")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.Context, clobtypes.LiquidationsConfig) error); ok {
		r0 = rf(ctx, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSubaccountLiquidationInfo provides a mock function with given fields: ctx, subaccountId, notionalLiquidatedQuoteQuantums, insuranceFundDeltaQuoteQuantums
func (_m *ClobKeeper) UpdateSubaccountLiquidationInfo(ctx types.Context, subaccountId subaccountstypes.SubaccountId, notionalLiquidatedQuoteQuantums *big.Int, insuranceFundDeltaQuoteQuantums *big.Int) {
	_m.Called(ctx, subaccountId, notionalLiquidatedQuoteQuantums, insuranceFundDeltaQuoteQuantums)
}

// NewClobKeeper creates a new instance of ClobKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClobKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClobKeeper {
	mock := &ClobKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
