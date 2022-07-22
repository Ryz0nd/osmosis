package wasmbinding

import (
	"encoding/json"
	"fmt"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/osmosis-labs/osmosis/v10/wasmbinding/bindings"
)

func StarGateQuerier(queryRouter baseapp.GRPCQueryRouter) func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error) {
	return func(ctx sdk.Context, request *wasmvmtypes.StargateQuery) ([]byte, error) {
		binding, whitelisted := StargateLayerBindings.Load(request.Path)
		if !whitelisted {
			return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("'%s' path is not allowed from the contract", request.Path)}
		}

		route := queryRouter.Route(request.Path)
		if route == nil {
			return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("No route to query '%s'", request.Path)}
		}

		res, err := route(ctx, abci.RequestQuery{
			Data: request.Data,
			Path: request.Path,
		})

		if err != nil {
			return nil, err
		}

		// normalize response to ensure backward compatibility
		bz, err := NormalizeReponses(binding, res.Value)
		if err != nil {
			return nil, err
		}

		return bz, nil
	}
}

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.OsmosisQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "osmosis query")
		}

		switch {
		case contractQuery.FullDenom != nil:
			creator := contractQuery.FullDenom.CreatorAddr
			subdenom := contractQuery.FullDenom.Subdenom

			fullDenom, err := GetFullDenom(creator, subdenom)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo full denom query")
			}

			res := bindings.FullDenomResponse{
				Denom: fullDenom,
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo full denom query response")
			}

			return bz, nil

		case contractQuery.DenomAdmin != nil:
			res, err := qp.GetDenomAdmin(ctx, contractQuery.DenomAdmin.Subdenom)
			if err != nil {
				return nil, err
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, fmt.Errorf("failed to JSON marshal DenomAdminResponse response: %w", err)
			}

			return bz, nil

		case contractQuery.PoolState != nil:
			poolId := contractQuery.PoolState.PoolId

			state, err := qp.GetPoolState(ctx, poolId)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo pool state query")
			}

			assets := ConvertSdkCoinsToWasmCoins(state.Assets)
			shares := ConvertSdkCoinToWasmCoin(state.Shares)

			res := bindings.PoolStateResponse{
				Assets: assets,
				Shares: shares,
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo pool state query response")
			}

			return bz, nil

		case contractQuery.SpotPrice != nil:
			spotPrice, err := qp.GetSpotPrice(ctx, contractQuery.SpotPrice)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo spot price query")
			}

			res := bindings.SpotPriceResponse{Price: spotPrice.String()}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo spot price query response")
			}

			return bz, nil

		case contractQuery.EstimateSwap != nil:
			swapAmount, err := qp.EstimateSwap(ctx, contractQuery.EstimateSwap)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo estimate swap query")
			}

			res := bindings.EstimatePriceResponse{Amount: *swapAmount}
			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "osmo estimate swap query response")
			}

			return bz, nil

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown osmosis query variant"}
		}
	}
}

func NormalizeReponses(binding interface{}, bz []byte) ([]byte, error) {
	// all values are proto message
	message, ok := binding.(proto.Message)
	if !ok {
		return nil, wasmvmtypes.Unknown{}
	}

	// unmarshal binary into stargate response data structure
	err := proto.Unmarshal(bz, message)
	if err != nil {
		return nil, wasmvmtypes.Unknown{}
	}

	// build new deterministic response
	bz, err = proto.Marshal(message)
	if err != nil {
		return nil, wasmvmtypes.Unknown{}
	}

	// clear proto message
	message.Reset()

	return bz, nil
}

// ConvertSdkCoinsToWasmCoins converts sdk type coins to wasm vm type coins
func ConvertSdkCoinsToWasmCoins(coins []sdk.Coin) wasmvmtypes.Coins {
	var toSend wasmvmtypes.Coins
	for _, coin := range coins {
		c := ConvertSdkCoinToWasmCoin(coin)
		toSend = append(toSend, c)
	}
	return toSend
}

// ConvertSdkCoinToWasmCoin converts a sdk type coin to a wasm vm type coin
func ConvertSdkCoinToWasmCoin(coin sdk.Coin) wasmvmtypes.Coin {
	return wasmvmtypes.Coin{
		Denom: coin.Denom,
		// Note: gamm tokens have 18 decimal places, so 10^22 is common, no longer in u64 range
		Amount: coin.Amount.String(),
	}
}
