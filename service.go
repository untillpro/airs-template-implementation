/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"context"

	gc "github.com/untillpro/gochips"
)

// Service s.e.
type Service struct {
	Host string
	Port int
}

type contextKeyType string

const contextKey = contextKeyType("contextKey")

func getService(ctx context.Context) *Service {
	return ctx.Value(contextKey).(*Service)
}

// Start s.e.
func (s *Service) Start(ctx context.Context) (context.Context, error) {
	gc.Info(*s)
	return context.WithValue(ctx, contextKey, s), nil
}

// Stop s.e.
func (s *Service) Stop(ctx context.Context) {

}
