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

// MapStaticXService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapStaticXService] method instead.
type MapStaticXService struct {
	Options []option.RequestOption
}

// NewMapStaticXService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMapStaticXService(opts ...option.RequestOption) (r MapStaticXService) {
	r = MapStaticXService{}
	r.Options = opts
	return
}

// Generates a raster image based on the given bounds (bounds based)
func (r *MapStaticXService) GetBoundsImage(ctx context.Context, format string, params MapStaticXGetBoundsImageParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.Scale == "" {
		err = errors.New("missing required scale parameter")
		return
	}
	if format == "" {
		err = errors.New("missing required format parameter")
		return
	}
	path := fmt.Sprintf("maps/%v/static/miny,\u0000minx,\u0000maxy,\u0000maxx/widthx\u0000height\u0000scale.%s", params.MapID, format)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

// Generates a raster image based on the specified center and zoom level (center
// based)
func (r *MapStaticXService) GetCenteredImage(ctx context.Context, format MapStaticXGetCenteredImageParamsFormat, params MapStaticXGetCenteredImageParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.Scale == "" {
		err = errors.New("missing required scale parameter")
		return
	}
	path := fmt.Sprintf("maps/%v/static/lat,\u0000lon,\u0000zoom/widthx\u0000height\u0000scale.%v", params.MapID, format)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

type MapStaticXGetBoundsImageParams struct {
	// Any of "`hybrid`", "`light`", "`dark`".
	MapID  MapStaticXGetBoundsImageParamsMapID `path:"mapId,omitzero,required" json:"-"`
	Miny   float64                             `path:"miny,required" json:"-"`
	Minx   float64                             `path:"minx,required" json:"-"`
	Maxy   float64                             `path:"maxy,required" json:"-"`
	Maxx   float64                             `path:"maxx,required" json:"-"`
	Width  int64                               `path:"width,required" json:"-"`
	Height int64                               `path:"height,required" json:"-"`
	Scale  string                              `path:"scale,required" json:"-"`
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
	Attribution MapStaticXGetBoundsImageParamsAttribution `query:"attribution,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MapStaticXGetBoundsImageParams]'s query parameters as
// `url.Values`.
func (r MapStaticXGetBoundsImageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MapStaticXGetBoundsImageParamsMapID string

const (
	MapStaticXGetBoundsImageParamsMapIDHybrid MapStaticXGetBoundsImageParamsMapID = "`hybrid`"
	MapStaticXGetBoundsImageParamsMapIDLight  MapStaticXGetBoundsImageParamsMapID = "`light`"
	MapStaticXGetBoundsImageParamsMapIDDark   MapStaticXGetBoundsImageParamsMapID = "`dark`"
)

// Changes the position of map attribution. If you disable the attribution make
// sure to display it in your application yourself (visibly).
type MapStaticXGetBoundsImageParamsAttribution string

const (
	MapStaticXGetBoundsImageParamsAttributionBottomright MapStaticXGetBoundsImageParamsAttribution = "`bottomright`"
	MapStaticXGetBoundsImageParamsAttributionBottomleft  MapStaticXGetBoundsImageParamsAttribution = "`bottomleft`"
	MapStaticXGetBoundsImageParamsAttributionTopleft     MapStaticXGetBoundsImageParamsAttribution = "`topleft`"
	MapStaticXGetBoundsImageParamsAttributionTopright    MapStaticXGetBoundsImageParamsAttribution = "`topright`"
	MapStaticXGetBoundsImageParamsAttributionFalse       MapStaticXGetBoundsImageParamsAttribution = "`false`"
)

type MapStaticXGetCenteredImageParams struct {
	// Any of "`hybrid`", "`light`", "`dark`".
	MapID  MapStaticXGetCenteredImageParamsMapID `path:"mapId,omitzero,required" json:"-"`
	Lat    float64                               `path:"lat,required" json:"-"`
	Lon    float64                               `path:"lon,required" json:"-"`
	Zoom   float64                               `path:"zoom,required" json:"-"`
	Width  int64                                 `path:"width,required" json:"-"`
	Height int64                                 `path:"height,required" json:"-"`
	Scale  string                                `path:"scale,required" json:"-"`
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
	Attribution MapStaticXGetCenteredImageParamsAttribution `query:"attribution,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MapStaticXGetCenteredImageParams]'s query parameters as
// `url.Values`.
func (r MapStaticXGetCenteredImageParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MapStaticXGetCenteredImageParamsMapID string

const (
	MapStaticXGetCenteredImageParamsMapIDHybrid MapStaticXGetCenteredImageParamsMapID = "`hybrid`"
	MapStaticXGetCenteredImageParamsMapIDLight  MapStaticXGetCenteredImageParamsMapID = "`light`"
	MapStaticXGetCenteredImageParamsMapIDDark   MapStaticXGetCenteredImageParamsMapID = "`dark`"
)

type MapStaticXGetCenteredImageParamsFormat string

const (
	MapStaticXGetCenteredImageParamsFormatPng  MapStaticXGetCenteredImageParamsFormat = "`png`"
	MapStaticXGetCenteredImageParamsFormatWebp MapStaticXGetCenteredImageParamsFormat = "`webp`"
	MapStaticXGetCenteredImageParamsFormatJpg  MapStaticXGetCenteredImageParamsFormat = "`jpg`"
)

// Changes the position of map attribution. If you disable the attribution make
// sure to display it in your application yourself (visibly).
type MapStaticXGetCenteredImageParamsAttribution string

const (
	MapStaticXGetCenteredImageParamsAttributionBottomright MapStaticXGetCenteredImageParamsAttribution = "`bottomright`"
	MapStaticXGetCenteredImageParamsAttributionBottomleft  MapStaticXGetCenteredImageParamsAttribution = "`bottomleft`"
	MapStaticXGetCenteredImageParamsAttributionTopleft     MapStaticXGetCenteredImageParamsAttribution = "`topleft`"
	MapStaticXGetCenteredImageParamsAttributionTopright    MapStaticXGetCenteredImageParamsAttribution = "`topright`"
	MapStaticXGetCenteredImageParamsAttributionFalse       MapStaticXGetCenteredImageParamsAttribution = "`false`"
)
