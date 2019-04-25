/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"context"
)

type service struct {
}

type contextKeyType string

const contextKey = contextKeyType("contextKey")

// Start s.e.
func (s *service) Start(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// Stop s.e.
func (s *service) Stop(ctx context.Context) {

}