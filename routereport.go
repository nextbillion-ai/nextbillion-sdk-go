// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// RouteReportService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteReportService] method instead.
type RouteReportService struct {
	Options []option.RequestOption
}

// NewRouteReportService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRouteReportService(opts ...option.RequestOption) (r RouteReportService) {
	r = RouteReportService{}
	r.Options = opts
	return
}

// Route Report
func (r *RouteReportService) New(ctx context.Context, params RouteReportNewParams, opts ...option.RequestOption) (res *RouteReportNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "route_report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type RouteReportNewResponse struct {
	// An array of objects returning encoded geometry of the routes. Each object
	// represents an individual route in the input.
	Geometry []string `json:"geometry"`
	// Returns the details of route segments in each state or country that the route
	// passes through. Each object represents an individual route in the input request.
	Mileage []RouteReportNewResponseMileage `json:"mileage"`
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// An array of objects returning a summary of the route with information about
	// tolls, bridges, tunnels, segments, maximum speeds and more. Each array
	// represents an individual route in the input request.
	RoadSummary []RouteReportNewResponseRoadSummary `json:"road_summary"`
	// A string indicating the state of the response. On normal responses, the value
	// will be Ok. Indicative HTTP error codes are returned for different errors. See
	// the
	// [**API Errors Codes**](https://app.reapi.com/ws/hmx8aL45B5jjrJa8/p/vNNilNksLVz675pI/s/ealJmVGjTQv4x5Wi/edit/path/VYzo7gOlRsQQZo0U#api-error-codes)
	// section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Mileage     respjson.Field
		Msg         respjson.Field
		RoadSummary respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseMileage struct {
	// Returns the details of road segments that the route covers in different states
	// and countries.
	Segment RouteReportNewResponseMileageSegment `json:"segment"`
	// Returns a summary of distances that the route covers in different states and
	// countries.
	Summary RouteReportNewResponseMileageSummary `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segment     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseMileage) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseMileage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the details of road segments that the route covers in different states
// and countries.
type RouteReportNewResponseMileageSegment struct {
	// An array of objects containing country-wise break up of the route segments. Each
	// object returns the segment details of a different country.
	Country []RouteReportNewResponseMileageSegmentCountry `json:"country"`
	// An array of objects containing state-wise break up of the route segments. Each
	// object returns the segment details of a different state.
	State []RouteReportNewResponseMileageSegmentState `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseMileageSegment) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseMileageSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseMileageSegmentCountry struct {
	// Represents the total distance of this segment, in meters.
	Distance float64 `json:"distance"`
	// Represents a sequence of ‘n’ consecutive vertices in the route geometry starting
	// from the offset, forming a continuous section of route with a distance indicated
	// in distancefield.
	Length int64 `json:"length"`
	// Represents the index value of the vertex of current segment's starting point in
	// route geometry. First vertex in the route geometry has an offset of 0.
	Offset int64 `json:"offset"`
	// Returns the name of the country in which the segment lies.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Length      respjson.Field
		Offset      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseMileageSegmentCountry) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseMileageSegmentCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseMileageSegmentState struct {
	// Represents the real distance of this segment, in meters.
	Distance float64 `json:"distance"`
	// Represents a sequence of ‘n’ consecutive vertices in the route geometry starting
	// from the offset, forming a continuous section of route with a distance indicated
	// in distancefield.
	Length int64 `json:"length"`
	// Represents the index value of the vertex of current segment's starting point in
	// route geometry. First vertex in the route geometry has an offset of 0.
	Offset int64 `json:"offset"`
	// Returns the name of the state in which the segment lies.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Length      respjson.Field
		Offset      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseMileageSegmentState) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseMileageSegmentState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns a summary of distances that the route covers in different states and
// countries.
type RouteReportNewResponseMileageSummary struct {
	// A break up of country-wise distances that the route covers in key:value pair
	// format.
	Country any `json:"country"`
	// A break up of state-wise distances that the route covers specified in key:value
	// pair format.
	State any `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseMileageSummary) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseMileageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseRoadSummary struct {
	// Returns the segment-wise road class and max speed information of the route.
	Segment RouteReportNewResponseRoadSummarySegment `json:"segment"`
	// Returns an overview of the route with information about trip distance, duration
	// and road class details among others.
	Summary RouteReportNewResponseRoadSummarySummary `json:"summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Segment     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseRoadSummary) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseRoadSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the segment-wise road class and max speed information of the route.
type RouteReportNewResponseRoadSummarySegment struct {
	// An array of objects returning the maximum speed of different segments that the
	// route goes through.
	MaxSpeed []RouteReportNewResponseRoadSummarySegmentMaxSpeed `json:"max_speed"`
	// An array of objects returning the details of road segments belonging to
	// different road classes that the route goes through. Each object refers to a
	// unique road class.
	RoadClass []RouteReportNewResponseRoadSummarySegmentRoadClass `json:"road_class"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxSpeed    respjson.Field
		RoadClass   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseRoadSummarySegment) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseRoadSummarySegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseRoadSummarySegmentMaxSpeed struct {
	// Returns the total distance of this segment, in meters.
	Distance int64 `json:"distance"`
	// Represents a sequence of ‘n’ consecutive vertices in the route geometry starting
	// from the offset, forming a continuous section of route where the maximum speed
	// is same and is indicated in value.
	Length int64 `json:"length"`
	// Represents the index value of the vertex of current segment's starting point in
	// route geometry. First vertex in the route geometry has an offset of 0.
	Offset int64 `json:"offset"`
	// Denotes the maximum speed of this segment, in kilometers per hour. - A value of
	// “-1” indicates that the speed is unlimited for this road segment. - A value of
	// “0” indicates that there is no information about the maximum speed for this road
	// segment.
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Length      respjson.Field
		Offset      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseRoadSummarySegmentMaxSpeed) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseRoadSummarySegmentMaxSpeed) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewResponseRoadSummarySegmentRoadClass struct {
	// Returns the total distance of this segment, in meters.
	Distance int64 `json:"distance"`
	// Represents a sequence of ‘n’ consecutive vertices in the route geometry starting
	// from the offset, forming a continuous section of route with a distance indicated
	// in distancefield.
	Length int64 `json:"length"`
	// Represents the index value of the vertex of current segment's starting point in
	// route geometry. First vertex in the route geometry has an offset of 0.
	Offset int64 `json:"offset"`
	// Returns the road class name to which the segment belongs.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Length      respjson.Field
		Offset      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseRoadSummarySegmentRoadClass) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseRoadSummarySegmentRoadClass) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns an overview of the route with information about trip distance, duration
// and road class details among others.
type RouteReportNewResponseRoadSummarySummary struct {
	// Returns the total distance of the route , in meters.
	Distance float64 `json:"distance"`
	// Returns the total duration of the route, in seconds.
	Duration float64 `json:"duration"`
	// A boolean value indicating if there are any bridges in the given route.
	HasBridge bool `json:"has_bridge"`
	// A boolean value indicating if there are any roundabouts in the given route.
	HasRoundabout bool `json:"has_roundabout"`
	// A boolean value indicating if there are any tolls in the given route.
	HasToll bool `json:"has_toll"`
	// A boolean value indicating if there are any tunnels in the given route.
	HasTunnel bool `json:"has_tunnel"`
	// An object with details about the different types of road classes that the route
	// goes through. Distance traversed on a given road class is also returned. The
	// contents of this object follow the key:value pair format.
	RoadClass any `json:"road_class"`
	// Returns the total distance travelled on toll roads. This field is present in the
	// response only when the has_toll property is true.
	TollDistance float64 `json:"toll_distance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance      respjson.Field
		Duration      respjson.Field
		HasBridge     respjson.Field
		HasRoundabout respjson.Field
		HasToll       respjson.Field
		HasTunnel     respjson.Field
		RoadClass     respjson.Field
		TollDistance  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteReportNewResponseRoadSummarySummary) RawJSON() string { return r.JSON.raw }
func (r *RouteReportNewResponseRoadSummarySummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteReportNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key" api:"required" format:"32 character alphanumeric string" json:"-"`
	// Takes a route geometry as input and returns the route details. Accepts polyline
	// and polyline6 encoded geometry as input.
	//
	// **Note**: Route geometries generated from sources other than
	// [NextBillion.ai](http://NextBillion.ai) services, are not supported in this
	// version.
	OriginalShape string `json:"original_shape" api:"required"`
	// Specify the encoding type of route geometry provided in original_shape input.
	// Please note that an error is returned when this parameter is not specified while
	// an input is added to original_shape parameter.
	//
	// Any of "polyline", "polyline6".
	OriginalShapeType RouteReportNewParamsOriginalShapeType `json:"original_shape_type,omitzero" api:"required"`
	paramObj
}

func (r RouteReportNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RouteReportNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RouteReportNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [RouteReportNewParams]'s query parameters as `url.Values`.
func (r RouteReportNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the encoding type of route geometry provided in original_shape input.
// Please note that an error is returned when this parameter is not specified while
// an input is added to original_shape parameter.
type RouteReportNewParamsOriginalShapeType string

const (
	RouteReportNewParamsOriginalShapeTypePolyline  RouteReportNewParamsOriginalShapeType = "polyline"
	RouteReportNewParamsOriginalShapeTypePolyline6 RouteReportNewParamsOriginalShapeType = "polyline6"
)
