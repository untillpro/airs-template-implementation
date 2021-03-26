/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 */

package iconfigcon

import (
	"testing"

	"github.com/stretchr/testify/require"
	// intf "github.com/untillpro/airs-istorage"
)

func Test_Impl(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	// Your tests here
	//intf.TestImpl(ctx, t)
}
