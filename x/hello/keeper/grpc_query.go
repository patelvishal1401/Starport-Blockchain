package keeper

import (
	"github.com/cosmonaut/hello/x/hello/types"
)

var _ types.QueryServer = Keeper{}
