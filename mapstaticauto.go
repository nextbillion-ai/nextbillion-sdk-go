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
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
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

// Generates a raster image based on the given features, auto fitted
func (r *MapStaticAutoService) GetAutoImage(ctx context.Context, format MapStaticAutoGetAutoImageParamsFormat, params MapStaticAutoGetAutoImageParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.Scale == "" {
		err = errors.New("missing required scale parameter")
		return
	}
	path := fmt.Sprintf("maps/%v/static/auto/widthx\u0000height\u0000scale.%v", params.MapID, format)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type MapStaticAutoGetAutoImageParams struct {
	// Any of "`hybrid`", "`light`", "`dark`".
	MapID  MapStaticAutoGetAutoImageParamsMapID `path:"mapId,omitzero,required" json:"-"`
	Width  int64                                `path:"width,required" json:"-"`
	Height int64                                `path:"height,required" json:"-"`
	Scale  string                               `path:"scale,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Define marker(s) to be drawn on top of the map. Can be used multiple times. Each
	// marker is a comma-separated [longitude,latitude] pair. Individual markers are
	// separated with a pipe symbol (`|`). Each marker can also have a different
	// color - specified as the 3rd parameter. Additionally, the following special
	// "commands" can be used (in the form of `key:value` appended before the `marker`
	// coordinates):
	//
	// `icon` - URL to a remote image (URL-encoded). You can use a URL shortener
	// service to create more compact queries. The icon has to be at most 64 kB and
	// 4096 pixels (e.g., 64x64 image). If used, the specified colors of the markers
	// are ignored.
	//
	// `anchor` - The anchor point of the custom icon. Allowed values are `top`,
	// `left`, `bottom`, `right`, `center`, `topleft`, `bottomleft`, `topright`,
	// `bottomright` with `bottom` being the default value.
	//
	// `scale` - the scale of the image (useful if you want to provide a HiDPI image
	// that scales down correctly), the default value is 1.
	Markers param.Opt[string] `query:"markers,omitzero" json:"-"`
	// Ensures the autofitted bounds or features are comfortably visible in the
	// resulting area. E.g. use 0.1 to add 10% margin (at least) of the size to each
	// side.
	Padding param.Opt[float64] `query:"padding,omitzero" json:"-"`
	// Define path(s) to be drawn on top of the map. The polyline co-ordinates should
	// be in (latitude,longitude) order. Use pipe (`|`) to separate the coordinates and
	// other properties of the path. Available properties for the `path` parameter are:
	// `fill` - Color to use as the fill (e.g. `red`, `rgba(255,255,255,0.5)`,
	// `#0000ff`, `none`, `false`) `stroke` - The color of the `path` stroke `width` -
	// Width of the `stroke` (in pixels) `enc` - Encoded path geometry in the
	// [Google Encoded Polyline Format](https://developers.google.com/maps/documentation/utilities/polylinealgorithm).
	// If used, the rest of the path parameter is considered to be part of the encoded
	// polyline string. Consequently, users need not specify the coordinate pairs.
	//
	// This parameter can be be used multiple times. You can also use `fill`, `stroke`
	// and `width` as separate query parameters to specify default styling for all the
	// paths.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Changes the position of map attribution. If you disable the attribution make
	// sure to display it in your application yourself (visibly).
	//
	// Any of "`bottomright`", "`bottomleft`", "`topleft`", "`topright`", "`false`".
	Attribution MapStaticAutoGetAutoImageParamsAttribution `query:"attribution,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MapStaticAutoGetAutoImageParams]'s query parameters as
// `url.Values`.
func (r MapStaticAutoGetAutoImageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MapStaticAutoGetAutoImageParamsMapID string

const (
	MapStaticAutoGetAutoImageParamsMapIDHybrid MapStaticAutoGetAutoImageParamsMapID = "`hybrid`"
	MapStaticAutoGetAutoImageParamsMapIDLight  MapStaticAutoGetAutoImageParamsMapID = "`light`"
	MapStaticAutoGetAutoImageParamsMapIDDark   MapStaticAutoGetAutoImageParamsMapID = "`dark`"
)

type MapStaticAutoGetAutoImageParamsFormat string

const (
	MapStaticAutoGetAutoImageParamsFormatPng  MapStaticAutoGetAutoImageParamsFormat = "`png`"
	MapStaticAutoGetAutoImageParamsFormatWebp MapStaticAutoGetAutoImageParamsFormat = "`webp`"
	MapStaticAutoGetAutoImageParamsFormatJpg  MapStaticAutoGetAutoImageParamsFormat = "`jpg`"
)

// Changes the position of map attribution. If you disable the attribution make
// sure to display it in your application yourself (visibly).
type MapStaticAutoGetAutoImageParamsAttribution string

const (
	MapStaticAutoGetAutoImageParamsAttributionBottomright MapStaticAutoGetAutoImageParamsAttribution = "`bottomright`"
	MapStaticAutoGetAutoImageParamsAttributionBottomleft  MapStaticAutoGetAutoImageParamsAttribution = "`bottomleft`"
	MapStaticAutoGetAutoImageParamsAttributionTopleft     MapStaticAutoGetAutoImageParamsAttribution = "`topleft`"
	MapStaticAutoGetAutoImageParamsAttributionTopright    MapStaticAutoGetAutoImageParamsAttribution = "`topright`"
	MapStaticAutoGetAutoImageParamsAttributionFalse       MapStaticAutoGetAutoImageParamsAttribution = "`false`"
)
