// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
)

// SkynetSkynetAssetService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetSkynetAssetService] method instead.
type SkynetSkynetAssetService struct {
	Options []option.RequestOption
}

// NewSkynetSkynetAssetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSkynetSkynetAssetService(opts ...option.RequestOption) (r SkynetSkynetAssetService) {
	r = SkynetSkynetAssetService{}
	r.Options = opts
	return
}

// Bind asset to device
func (r *SkynetSkynetAssetService) Bind(ctx context.Context, id string, params SkynetSkynetAssetBindParams, opts ...option.RequestOption) (res *SimpleResp, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("skynet/skynet/asset/%s/bind", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SkynetSkynetAssetBindParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Device ID to be linked to the `asset` identified by `id`.
	//
	// Please note that the device needs to be linked to an `asset` before using it in
	// the _Upload locations of an Asset_ method for sending GPS information about the
	// `asset`.
	DeviceID string `json:"device_id,required"`
	paramObj
}

func (r SkynetSkynetAssetBindParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSkynetAssetBindParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSkynetAssetBindParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetSkynetAssetBindParams]'s query parameters as
// `url.Values`.
func (r SkynetSkynetAssetBindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
