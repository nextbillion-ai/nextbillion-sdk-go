// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
)

// MapStaticService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapStaticService] method instead.
type MapStaticService struct {
	Options []option.RequestOption
	X       MapStaticXService
	Auto    MapStaticAutoService
}

// NewMapStaticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMapStaticService(opts ...option.RequestOption) (r MapStaticService) {
	r = MapStaticService{}
	r.Options = opts
	r.X = NewMapStaticXService(opts...)
	r.Auto = NewMapStaticAutoService(opts...)
	return
}
