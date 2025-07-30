// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// IsochroneService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIsochroneService] method instead.
type IsochroneService struct {
	Options []option.RequestOption
}

// NewIsochroneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIsochroneService(opts ...option.RequestOption) (r IsochroneService) {
	r = IsochroneService{}
	r.Options = opts
	return
}

// The NextBillion.ai Isochrone API computes areas that are reachable within a
// specified amount of time from a location, and returns the reachable regions as
// contours of polygons or lines that you can display on a map.
func (r *IsochroneService) Compute(ctx context.Context, query IsochroneComputeParams, opts ...option.RequestOption) (res *IsochroneComputeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "isochrone/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type IsochroneComputeResponse struct {
	// A
	// [GeoJSON FeatureCollection](https://datatracker.ietf.org/doc/html/rfc7946#section-3.3)
	// object with details of the isochrone contours. Each feature object in this
	// collection represents an isochrone.
	Features []IsochroneComputeResponseFeature `json:"features"`
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// A string indicating the state of the response. On normal responses, the value
	// will be Ok. Indicative HTTP error codes are returned for different errors. See
	// the [API Errors Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// Type of the GeoJSON object. As prescribed in
	// [GeoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946#section-1.4),
	// its value is FeatureCollection as the feature property contains a list of
	// geoJSON feature objects.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		Msg         respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsochroneComputeResponse) RawJSON() string { return r.JSON.raw }
func (r *IsochroneComputeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsochroneComputeResponseFeature struct {
	// A [GeoJSON geometry](https://datatracker.ietf.org/doc/html/rfc7946#page-7)
	// object with details of the contour line.
	Geometry IsochroneComputeResponseFeatureGeometry `json:"geometry"`
	// An object with details of how the isochrone contour can be drawn on a map.
	Properties IsochroneComputeResponseFeatureProperties `json:"properties"`
	// Type of the GeoJSON object. Its value is Feature as per the
	// [GeoJSON standard](https://datatracker.ietf.org/doc/html/rfc7946#section-1.4)
	// object.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsochroneComputeResponseFeature) RawJSON() string { return r.JSON.raw }
func (r *IsochroneComputeResponseFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A [GeoJSON geometry](https://datatracker.ietf.org/doc/html/rfc7946#page-7)
// object with details of the contour line.
type IsochroneComputeResponseFeatureGeometry struct {
	// An array of coordinate points, in [longitude,latitude] format representing the
	// isochrone contour line.
	Coordinates []float64 `json:"coordinates"`
	// Type of the geoJSON geometry.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsochroneComputeResponseFeatureGeometry) RawJSON() string { return r.JSON.raw }
func (r *IsochroneComputeResponseFeatureGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of how the isochrone contour can be drawn on a map.
type IsochroneComputeResponseFeatureProperties struct {
	// The hex code of the color of the isochrone contour line
	Color string `json:"color"`
	// The value of the metric used in this contour. See the metric property to
	// determine whether this is a time or distance contour. When the metric is time
	// this value denotes the travel time in minutes and when the metric is distance
	// this value denotes the travel distance in kilometers.
	Contour float64 `json:"contour"`
	// The hex code for the fill color of the isochrone contour line.
	Fill string `json:"fill"`
	// The hex code for the fill color of the isochrone contour line
	FillColor string `json:"fillColor"`
	// The fill opacity for the isochrone contour line. It is a float value starting
	// from 0.0 with a max value of 1.0. Higher number indicates a higher fill opacity.
	FillOpacity float64 `json:"fillOpacity"`
	// The metric that the contour represents - either distance or time
	Metric string `json:"metric"`
	// The opacity of the isochrone contour line. It is a float value starting from 0.0
	// with a max value of 1.0. Higher number indicates a higher line opacity
	Opacity float64 `json:"opacity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color       respjson.Field
		Contour     respjson.Field
		Fill        respjson.Field
		FillColor   respjson.Field
		FillOpacity respjson.Field
		Metric      respjson.Field
		Opacity     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IsochroneComputeResponseFeatureProperties) RawJSON() string { return r.JSON.raw }
func (r *IsochroneComputeResponseFeatureProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IsochroneComputeParams struct {
	// The distances, in meters, to use for each isochrone contour. You can specify up
	// to four contours. Distances must be in increasing order. The maximum distance
	// that can be specified is 60000 meters (60 km).
	ContoursMeters int64 `query:"contours_meters,required" json:"-"`
	// The times, in minutes, to use for each isochrone contour. You can specify up to
	// four contours. Times must be in increasing order. The maximum time that can be
	// specified is 40 minutes.
	ContoursMinutes int64 `query:"contours_minutes,required" json:"-"`
	// The coordinates of the location which acts as the starting point for which the
	// isochrone lines need to be determined.
	Coordinates string `query:"coordinates,required" format:"latitude,longitude" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// The hex code of the color to fill isochrone contour. When requesting multiple
	// contours, it is recommended to provide color codes for each of the requested
	// contours, separated by a ",". If no colors are specified, the Isochrone API will
	// assign a random color scheme to the output.
	ContoursColors param.Opt[string] `query:"contours_colors,omitzero" format:"value_1,value_2,..." json:"-"`
	// A floating point value from 0.0 to 1.0 can be used to remove smaller contours.
	// The default is 1.0. A value of 1.0 will only return the largest contour for a
	// given value. A value of 0.5 drops any contours that are less than half the area
	// of the largest contour in the set of contours for that same value.
	Denoise param.Opt[float64] `query:"denoise,omitzero" json:"-"`
	// Use this parameter to set a departure time, expressed as UNIX epoch timestamp in
	// seconds, for calculating the isochrone contour. The response will consider the
	// typical traffic conditions at the given time and return a contour which can be
	// reached under those traffic conditions. Please note that if no input is provided
	// for this parameter then the traffic conditions at the time of making the request
	// are considered.
	DepartureTime param.Opt[int64] `query:"departure_time,omitzero" json:"-"`
	// A positive floating point value, in meters, used as the tolerance for
	// Douglas-Peucker generalization. There is no upper bound. If no value is
	// specified in the request, the Isochrone API will choose the most optimized
	// generalization to use for the request. Note that the generalization of contours
	// can lead to self-intersections, as well as intersections of adjacent contours.
	Generalize param.Opt[float64] `query:"generalize,omitzero" json:"-"`
	// Specify whether to return the contours as GeoJSON polygons (true) or linestrings
	// (false, default). When polygons=true, any contour that forms a ring is returned
	// as a polygon.
	Polygons param.Opt[bool] `query:"polygons,omitzero" json:"-"`
	// Set which driving mode the service should use to determine the contour. For
	// example, if you use "car", the API will return an isochrone contour that a car
	// can reach within the specified time or after driving the specified distance.
	// Using "truck" will return a contour that a truck can reach after taking into
	// account appropriate truck routing restrictions.
	//
	// Note: Only the "car" profile is enabled by default. Please note that customized
	// profiles (including "truck") might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "car", "truck".
	Mode IsochroneComputeParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IsochroneComputeParams]'s query parameters as `url.Values`.
func (r IsochroneComputeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set which driving mode the service should use to determine the contour. For
// example, if you use "car", the API will return an isochrone contour that a car
// can reach within the specified time or after driving the specified distance.
// Using "truck" will return a contour that a truck can reach after taking into
// account appropriate truck routing restrictions.
//
// Note: Only the "car" profile is enabled by default. Please note that customized
// profiles (including "truck") might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type IsochroneComputeParamsMode string

const (
	IsochroneComputeParamsModeCar   IsochroneComputeParamsMode = "car"
	IsochroneComputeParamsModeTruck IsochroneComputeParamsMode = "truck"
)
