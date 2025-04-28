package blockchain

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/xyxa.gg/backend-mvp-main/config"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain/contracts"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/blockchain/events"
	"gitlab.com/xyxa.gg/backend-mvp-main/pkg/logger"
)

type BlockchainServer struct {
	cfg       *config.Config
	log       logger.Logger
	client    *ethclient.Client
	listeners []events.EventListener
}

func NewBlochchainServer(cfg *config.Config, log logger.Logger) *BlockchainServer {
	return &BlockchainServer{
		cfg:       cfg,
		log:       log,
		client:    nil,
		listeners: []events.EventListener{},
	}
}

func (s *BlockchainServer) InitConnection() error {
	client, err := ethclient.Dial(s.cfg.Blockchain.NodeURL)
	if err != nil {
		return err
	}

	s.client = client
	return nil
}

func (s *BlockchainServer) RegisterListeners(listeners []events.EventListener) {
	for _, listener := range listeners {
		s.listeners = append(s.listeners, listener)
	}
}

func (s *BlockchainServer) Listen(ctx context.Context) error {
	contract, err := contracts.NewContracts(common.HexToAddress(s.cfg.Blockchain.ContractAddress), s.client)
	if err != nil {
		return errors.New(fmt.Sprintf("Init Error: Contract %s can't be initialized", s.cfg.Blockchain.ContractAddress))
	}

	opts := &bind.WatchOpts{Context: ctx, Start: nil}

	for _, event := range s.listeners {
		go event.Listen(contract, opts)
	}

	return nil
}

func (s *BlockchainServer) Release(ctx context.Context) {
	for _, listener := range s.listeners {
		listener.Release()
	}

	s.client.Close()
}
