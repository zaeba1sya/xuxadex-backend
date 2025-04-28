package events

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain/contracts"
)

type EventListener interface {
	Listen(contract *contracts.Contracts, opts *bind.WatchOpts)
	Release()
}
