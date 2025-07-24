// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
)

// TileService contains methods and other services that help with interacting with
// the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTileService] method instead.
type TileService struct {
	Options []option.RequestOption
}

// NewTileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTileService(opts ...option.RequestOption) (r TileService) {
	r = TileService{}
	r.Options = opts
	return
}

// Get vector tiles
func (r *TileService) Get(ctx context.Context, key string, params TileGetParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.TilesID == "" {
		err = errors.New("missing required tilesId parameter")
		return
	}
	if params.Format == "" {
		err = errors.New("missing required format parameter")
		return
	}
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("tiles/%s/%v/%v/%v.%s?%s", params.TilesID, params.Z, params.X, params.Y, params.Format, key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type TileGetParams struct {
	TilesID string `path:"tilesId,required" json:"-"`
	Z       int64  `path:"z,required" json:"-"`
	X       int64  `path:"x,required" json:"-"`
	Y       int64  `path:"y,required" json:"-"`
	Format  string `path:"format,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [TileGetParams]'s query parameters as `url.Values`.
func (r TileGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
