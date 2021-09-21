/*
 * Copyright (c) <year>-present unTill Pro, Ltd.
 */

package itemplate

import (
	"context"
)

type service struct{
	params Params
}

type contextKeyType string

const contextKey = contextKeyType("contextKey")

func getService(ctx context.Context) *service {
	return ctx.Value(contextKey).(*service)
}

// Start s.e.
func (s *service) Start(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, contextKey, s), nil
}

// Stop s.e.
func (s *service) Stop(ctx context.Context) {

}
