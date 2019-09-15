package api

import (
	spec "github.com/typusomega/cqrsorders/pkg/spec/v1"
	"context"
)

func New() *API {
	return &API{}
}

func (it *API) CreateOrder(ctx context.Context, order *spec.Order) (*spec.Order, error) {
	return nil, nil
}

type API struct {
}