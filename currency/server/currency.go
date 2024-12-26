package server

import (
	"context"
	protos "currency/protos/currency/protos"
	"github.com/hashicorp/go-hclog"
)

type Currency struct {
	protos.UnimplementedCurrencyServer // Embed the unimplemented server
	log                                hclog.Logger
}

func (c Currency) GetRate(ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle request for GetRate", "base", rr.GetBase(), "dest", rr.GetDestination())
	return &protos.RateResponse{Rate: 0.5}, nil
}

func NewCurrency(log hclog.Logger) *Currency {
	return &Currency{log: log}
}
