/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

/*

	Test service start/stop here

*/

package iconfigcon

import (
	"context"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/untillpro/godif"
	gc "github.com/untillpro/gochips"
	"github.com/untillpro/godif/iservices"
	"github.com/untillpro/godif/services"
)


func Test_StartStop(t *testing.T){
	ctx, err := start(t)
	gc.FatalIfError(t, err, "start failed")


	/*

		Your tests here

	*/


	defer stop(ctx, t)
}

func start(t *testing.T) (context.Context, error) {

	// Provide iservices interface

	godif.Require(&iservices.Start)
	godif.Require(&iservices.Stop)

	services.Declare()

	// Declare own service
	Declare()

	errs := godif.ResolveAll()
	require.True(t, len(errs) == 0, "Resolve problem", errs)

	ctx := context.Background()
	ctx, err := iservices.Start(ctx)
	return ctx, err

}

func stop(ctx context.Context, t *testing.T) {
	iservices.Stop(ctx)
	godif.Reset()
}
