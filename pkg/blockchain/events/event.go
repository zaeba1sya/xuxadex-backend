package events

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/xuxadex/backend-mvp-main/pkg/blockchain/contracts"
)

type EventListener interface {
	Listen(contract *contracts.Contracts, opts *bind.WatchOpts)
	Release()
}
