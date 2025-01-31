package v24

import (
	"github.com/osmosis-labs/osmosis/v23/app/upgrades"

	store "github.com/cosmos/cosmos-sdk/store/types"
)

// UpgradeName defines the on-chain upgrade name for the Osmosis v24 upgrade.
const UpgradeName = "v24"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
