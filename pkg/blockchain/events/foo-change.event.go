package events

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain/contracts"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
)

type fooChangeEvent struct {
	ctx      context.Context
	log      logger.Logger
	consumer chan *contracts.ContractsFooChangeEvent
}

func NewFooChangeEvent(ctx context.Context, log logger.Logger) *fooChangeEvent {
	return &fooChangeEvent{
		ctx:      ctx,
		log:      log,
		consumer: nil,
	}
}

func (e *fooChangeEvent) Listen(contract *contracts.Contracts, opts *bind.WatchOpts) {
	e.consumer = make(chan *contracts.ContractsFooChangeEvent)

	sub, err := contract.WatchFooChangeEvent(opts, e.consumer)
	if err != nil {
		e.log.Errorf("Failed to watch foo change event: %v", err)
	}
	defer sub.Unsubscribe()

	for {
		select {
		case event, ok := <-e.consumer:
			if !ok {
				return
			}
			e.log.Infof("Received FooChangeEvent(%s, %s, %d, %s)", event.Sender.String(), event.Foo1.String(), event.Foo2.Uint64(), event.Foo3)
		case <-e.ctx.Done():
			return
		}
	}
}

func (e *fooChangeEvent) Release() {
	if e.consumer != nil {
		close(e.consumer)
	}
}
