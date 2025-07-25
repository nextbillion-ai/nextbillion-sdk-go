// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
)

// MapStaticAutoService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapStaticAutoService] method instead.
type MapStaticAutoService struct {
	Options []option.RequestOption
}

// NewMapStaticAutoService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMapStaticAutoService(opts ...option.RequestOption) (r MapStaticAutoService) {
	r = MapStaticAutoService{}
	r.Options = opts
	return
}
