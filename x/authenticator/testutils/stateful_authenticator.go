package testutils

import (
	"encoding/json"

	"github.com/osmosis-labs/osmosis/v19/x/authenticator/iface"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ iface.Authenticator     = &StatefulAuthenticator{}
	_ iface.AuthenticatorData = &StatefulAuthenticatorData{}
)

type StatefulAuthenticatorData struct {
	Value int
}

// StatefulAuthenticator is an experiment of how to write authenticators that handle state
type StatefulAuthenticator struct {
	KvStoreKey sdk.StoreKey
}

func (s StatefulAuthenticator) Type() string {
	return "Stateful"
}

func (s StatefulAuthenticator) StaticGas() uint64 {
	return 1000
}

func (s StatefulAuthenticator) Initialize(data []byte) (iface.Authenticator, error) {
	return s, nil
}

func (s StatefulAuthenticator) GetAuthenticationData(ctx sdk.Context, tx sdk.Tx, messageIndex int8, simulate bool) (iface.AuthenticatorData, error) {
	return StatefulAuthenticatorData{Value: s.GetValue(ctx)}, nil
}

func (s StatefulAuthenticator) Authenticate(ctx sdk.Context, account sdk.AccAddress, msg sdk.Msg, authenticationData iface.AuthenticatorData) iface.AuthenticationResult {
	statefulData, ok := authenticationData.(StatefulAuthenticatorData)
	if !ok {
		return iface.Rejected("", sdkerrors.Wrap(sdkerrors.ErrInvalidType, "authenticationData is not StatefulAuthenticatorData"))
	}
	if statefulData.Value > 10 {
		return iface.Rejected("value is too high", nil)
	}
	s.SetValue(ctx, statefulData.Value+1)
	return iface.Authenticated()
}

func (s StatefulAuthenticator) Track(ctx sdk.Context, account sdk.AccAddress, msg sdk.Msg) error {
	return nil
}

func (s StatefulAuthenticator) SetValue(ctx sdk.Context, value int) {
	kvStore := prefix.NewStore(ctx.KVStore(s.KvStoreKey), []byte(s.Type()))
	statefulData := StatefulAuthenticatorData{Value: value}
	newBz, _ := json.Marshal(statefulData)
	kvStore.Set([]byte("value"), newBz)
}

func (s StatefulAuthenticator) GetValue(ctx sdk.Context) int {
	kvStore := prefix.NewStore(ctx.KVStore(s.KvStoreKey), []byte(s.Type()))
	bz := kvStore.Get([]byte("value")) // global value. On the real thing we may want the account
	var statefulData StatefulAuthenticatorData
	_ = json.Unmarshal(bz, &statefulData) // if we can't unmarshal, we just assume it's 0
	return statefulData.Value
}

func (s StatefulAuthenticator) ConfirmExecution(ctx sdk.Context, account sdk.AccAddress, msg sdk.Msg, authenticationData iface.AuthenticatorData) iface.ConfirmationResult {
	s.SetValue(ctx, s.GetValue(ctx)+1)
	return iface.Confirm()
}

func (s StatefulAuthenticator) OnAuthenticatorAdded(ctx sdk.Context, account sdk.AccAddress, data []byte) error {
	return nil
}

func (s StatefulAuthenticator) OnAuthenticatorRemoved(ctx sdk.Context, account sdk.AccAddress, data []byte) error {
	return nil
}
