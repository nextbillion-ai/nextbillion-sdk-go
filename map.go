// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
)

// MapService contains methods and other services that help with interacting with
// the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapService] method instead.
type MapService struct {
	Options []option.RequestOption
	Static  MapStaticService
}

// NewMapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMapService(opts ...option.RequestOption) (r MapService) {
	r = MapService{}
	r.Options = opts
	r.Static = NewMapStaticService(opts...)
	return
}

// Road Segments
func (r *MapService) NewSegment(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "map/segments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// GET raster tiles
func (r *MapService) GetTile(ctx context.Context, format MapGetTileParamsFormat, params MapGetTileParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("maps/%v/%v/%v/%v/y\u0000scale.%v", params.MapID, params.TileSize, params.Z, params.X, format)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type MapGetTileParams struct {
	// Any of "`hybrid`", "`dark`", "`default`".
	MapID MapGetTileParamsMapID `path:"mapId,omitzero,required" json:"-"`
	// Any of .
	TileSize int64 `path:"tileSize,omitzero,required" json:"-"`
	Z        int64 `path:"z,required" json:"-"`
	X        int64 `path:"x,required" json:"-"`
	Y        int64 `path:"y,required" json:"-"`
	// Any of "`@2x`".
	Scale MapGetTileParamsScale `path:"scale,omitzero,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to an
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [MapGetTileParams]'s query parameters as `url.Values`.
func (r MapGetTileParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MapGetTileParamsMapID string

const (
	MapGetTileParamsMapIDHybrid  MapGetTileParamsMapID = "`hybrid`"
	MapGetTileParamsMapIDDark    MapGetTileParamsMapID = "`dark`"
	MapGetTileParamsMapIDDefault MapGetTileParamsMapID = "`default`"
)

type MapGetTileParamsScale string

const (
	MapGetTileParamsScale2x MapGetTileParamsScale = "`@2x`"
)

type MapGetTileParamsFormat string

const (
	MapGetTileParamsFormatJpg MapGetTileParamsFormat = "`jpg`"
)
