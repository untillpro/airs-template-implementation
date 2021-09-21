/*
 * Copyright (c) <year>-present unTill Pro, Ltd.
 */

package itemplate

import (
	"github.com/untillpro/godif"
	"github.com/untillpro/godif/services"
)

// Params s.e.
type Params struct {
	Host string
	Port int
}

// Declare s.e.
func Declare(params Params) {
	var svc service
	svc.params = params
	godif.ProvideSliceElement(&services.Services, &svc)

	//godif.Provide(&???.???, implFunc)
}
