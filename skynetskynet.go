// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
)

// SkynetSkynetService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetSkynetService] method instead.
type SkynetSkynetService struct {
	Options []option.RequestOption
	Asset   SkynetSkynetAssetService
}

// NewSkynetSkynetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSkynetSkynetService(opts ...option.RequestOption) (r SkynetSkynetService) {
	r = SkynetSkynetService{}
	r.Options = opts
	r.Asset = NewSkynetSkynetAssetService(opts...)
	return
}
