package keeper_test

import (
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v21/x/authenticator/authenticator"
)

func (s *KeeperTestSuite) TestKeeper_AddAuthenticatorWithId() {
	ctx := s.Ctx

	// Ensure the SigVerificationAuthenticator type is registered
	s.Require().True(s.am.IsAuthenticatorTypeRegistered(authenticator.SignatureVerificationAuthenticator{}.Type()))

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	err := s.App.AuthenticatorKeeper.AddAuthenticatorWithId(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
		0,
	)
	s.Require().NoError(err)

	err = s.App.AuthenticatorKeeper.AddAuthenticatorWithId(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		priv.PubKey().Bytes(),
		1,
	)
	s.Require().NoError(err)

	_, authenticators, err := s.App.AuthenticatorKeeper.GetAuthenticatorsForAccount(ctx, accAddress)
	s.Require().NoError(err)
	s.Require().Equal(len(authenticators), 2, "Getting authenticators returning incorrect data")

	err = s.App.AuthenticatorKeeper.AddAuthenticatorWithId(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		[]byte("BrokenBytes"),
		2,
	)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "invalid secp256k1 public key size")

	s.App.AuthenticatorManager.ResetAuthenticators()
	err = s.App.AuthenticatorKeeper.AddAuthenticatorWithId(
		ctx,
		accAddress,
		"SignatureVerificationAuthenticator",
		[]byte("BrokenBytes"),
		2,
	)
	s.Require().Error(err)
	s.Require().ErrorContains(err, "authenticator type")
}

func (s *KeeperTestSuite) TestKeeper_GetAllAuthenticatorDataGenesis() {
	ctx := s.Ctx

	// Ensure the SigVerificationAuthenticator type is registered
	s.Require().True(s.am.IsAuthenticatorTypeRegistered(authenticator.SignatureVerificationAuthenticator{}.Type()))

	// Set up account
	key := "6cf5103c60c939a5f38e383b52239c5296c968579eec1c68a47d70fbf1d19159"
	bz, _ := hex.DecodeString(key)
	priv := &secp256k1.PrivKey{Key: bz}
	accAddress := sdk.AccAddress(priv.PubKey().Address())

	for i := 0; i < 5; i++ {
		err := s.App.AuthenticatorKeeper.AddAuthenticator(
			ctx,
			accAddress,
			"SignatureVerificationAuthenticator",
			priv.PubKey().Bytes(),
		)
		s.Require().NoError(err)
	}

	authenticators, err := s.App.AuthenticatorKeeper.GetAllAuthenticatorData(ctx)
	s.Require().NoError(err, "Error getting authenticator data")
	s.Require().Equal(5, len(authenticators[0].Authenticators), "Getting authenticators returning incorrect data")
	s.Require().Equal(accAddress.String(), authenticators[0].Address, "Authenticator Address is incorrect")
}
