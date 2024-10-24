package types

import (
	"context"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	SendCoins(ctx context.Context, from, to sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx context.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	IsSendEnabledCoins(ctx context.Context, coins ...sdk.Coin) error
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// WasmKeeper defines the expected interface for the wasm module.
type WasmKeeper interface {
	GetContractInfo(ctx context.Context, contractAddress sdk.AccAddress) *wasmtypes.ContractInfo
	HasContractInfo(ctx context.Context, contractAddress sdk.AccAddress) bool
}

// ContractKeeper defines the expected interface for the wasm module.
type ContractKeeper interface {
	Execute(ctx sdk.Context, contractAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
}
