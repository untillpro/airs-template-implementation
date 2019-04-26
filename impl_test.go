/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Impl(t *testing.T) {
	ctx := start(t)
	defer stop(ctx, t)

	require.True(t, implFunc(ctx))

	//iconfig.TestImpl(ctx, t)

}
