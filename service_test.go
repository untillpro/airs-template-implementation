/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 */

/*

	Test service start/stop here

*/

package iconfigcon

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif/services"

//	intf "github.com/untillpro/airs-istorage"
)

func Test_StartStop(t *testing.T) {
	ctx, err := setUp(t)
	defer tearDown(ctx, t)
	require.Nil(t, err, err)

	/*
		Your tests here
	*/
}

func setUp(t *testing.T) (context.Context, error) {

	// Declare own service
	Declare(Params{Host: "localhost", Port: 800})

	// intf.DeclareForTest()

	return services.ResolveAndStart()
}

func tearDown(ctx context.Context, t *testing.T) {
	services.StopAndReset(ctx)
}
