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

// DistanceMatrixJsonService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDistanceMatrixJsonService] method instead.
type DistanceMatrixJsonService struct {
	Options []option.RequestOption
}

// NewDistanceMatrixJsonService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDistanceMatrixJsonService(opts ...option.RequestOption) (r DistanceMatrixJsonService) {
	r = DistanceMatrixJsonService{}
	r.Options = opts
	return
}

// asfd
func (r *DistanceMatrixJsonService) New(ctx context.Context, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "distancematrix/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return
}

// Nextbillion.ai Distance Matrix API computes distances and ETAs between a set of
// origins and destinations — could be for one-to-many or many-to-many scenarios.
// The API call returns a matrix of ETAs and distances for each origin and
// destination pair. For example, If the set is Origins {A,B} and Destinations
// {C,D,E} we can get the following set of results with distance (meters) and time
// (seconds) for each. The GET method can only handle up to 100 locations (1
// location is either 1 origin or 1 destination).
func (r *DistanceMatrixJsonService) Get(ctx context.Context, query DistanceMatrixJsonGetParams, opts ...option.RequestOption) (res *DistanceMatrixJsonGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "distancematrix/json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DistanceMatrixJsonGetResponse struct {
	// Displays the error message in case of a failed request or operation. Please note
	// that this parameter is not returned in the response in case of a successful
	// request.
	Msg string `json:"msg"`
	// Container object for a response with an array of arrays structure.
	Rows []DistanceMatrixJsonGetResponseRow `json:"rows"`
	// A string indicating the state of the response. On normal responses, the value
	// will be Ok. Indicative HTTP error codes are returned for different errors. See
	// the [API Errors Codes](#api-error-codes) section below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Rows        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DistanceMatrixJsonGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DistanceMatrixJsonGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DistanceMatrixJsonGetResponseRow struct {
	// An array of objects. Each elements array corresponds to a single origins
	// coordinate and contains objects with distance and duration values for each of
	// the destinations. The details in the first elements array correspond to the
	// first origins point and the first object corresponds to the first destinations
	// point and so on.
	Elements []DistanceMatrixJsonGetResponseRowElement `json:"elements"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Elements    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DistanceMatrixJsonGetResponseRow) RawJSON() string { return r.JSON.raw }
func (r *DistanceMatrixJsonGetResponseRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DistanceMatrixJsonGetResponseRowElement struct {
	// Distance of the route from an origin to a destination, in meters.
	Distance float64 `json:"distance"`
	// Duration of the trip from an origin to a destination, in seconds.
	Duration float64 `json:"duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DistanceMatrixJsonGetResponseRowElement) RawJSON() string { return r.JSON.raw }
func (r *DistanceMatrixJsonGetResponseRowElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DistanceMatrixJsonGetParams struct {
	// "destinations" are the ending coordinates of your route. Ensure that
	// "destinations" are routable land locations. Multiple "destinations" should be
	// separated by a pipe symbol "|".
	Destinations string `query:"destinations,required" format:"latitude_1, longitue_1|latitude_2, longitude_2|..." json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// "origins" are the starting point of your route. Ensure that "origins" are
	// routable land locations. Multiple "origins" should be separated by a pipe symbol
	// "|".
	Origins string `query:"origins,required" format:"latitude_1,longitude_1|latitude_2,longitude_2|..." json:"-"`
	// Limits the search to segments with given bearing in degrees towards true north
	// in clockwise direction. Each "bearing" should be in the format of
	// "degree,range", where the "degree" should be a value between \[0, 360\] and
	// "range" should be a value between \[0, 180\]. Please note that the number of
	// "bearings" should be equal to the sum of the number of points in "origins" and
	// "destinations". If a route can approach a destination from any direction, the
	// bearing for that point can be specified as "0,180".
	Bearings param.Opt[string] `query:"bearings,omitzero" format:"degree_0,range_0;degree_1,range_1;..." json:"-"`
	// A prompt to modify the response in case no feasible route is available for a
	// given pair of origin and destination. When set to "true", a value of "-1" is
	// returned for those pairs in which:
	//
	// \- Either origin or the destination can not be snapped to a nearest road. Please
	// note that if all the origins and destinations in a request can't be snapped to
	// their nearest roads, a 4xx error is returned instead, as the entire request
	// failed.
	//
	// \- Both origin and destination can be snapped to the nearest road, but the
	// service can't find a valid route between them. However, a value of "0" is
	// returned if both the origin and destination are snapped to the same location.
	//
	// "false" is the default value. In this case, a "0" value is returned for all the
	// above cases. A 4xx error is returned, in this case as well, when all origins and
	// destinations in the request can't be snapped to their nearest road.
	RouteFailedPrompt param.Opt[bool] `query:"route_failed_prompt,omitzero" json:"-"`
	// A semicolon-separated list indicating the side of the road from which the route
	// will approach "destinations". When set to "unrestricted" a route can arrive at a
	// destination from either side of the road. When set to "curb" the route will
	// arrive at a destination on the driving side of the region. Please note the
	// number of values provided must be equal to the number of "destinations".
	// However, you can skip a coordinate and show its position in the list with the
	// ";" separator. The values provided for the "approaches" parameter are effective
	// for the "destinations" value at the same index. Example: "curb;;curb" will apply
	// curbside restriction on the "destinations" points provided at the first and
	// third index.
	//
	// Any of "unrestricted", "curb".
	Approaches DistanceMatrixJsonGetParamsApproaches `query:"approaches,omitzero" json:"-"`
	// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
	// Multiple values should be separated by a pipe (|). If "none" is provided along
	// with other values, an error is returned as a valid route is not feasible. Please
	// note that when this parameter is not provided in the input, ferries are set to
	// be avoided by default. When this parameter is provided, only the mentioned
	// objects are avoided.
	//
	// Any of "toll", "ferry", "highway", "none".
	Avoid DistanceMatrixJsonGetParamsAvoid `query:"avoid,omitzero" json:"-"`
	// Set which driving mode the service should use to determine the "distance" and
	// "duration" values. For example, if you use "car", the API will return the
	// duration and distance of a route that a car can take. Using "truck" will return
	// the same for a route a truck can use, taking into account appropriate truck
	// routing restrictions.
	//
	// When "mode=truck", following are the default dimensions that are used:
	//
	// \- truck_height = 214 centimeters
	//
	// \- truck_width = 183 centimeters
	//
	// \- truck_length = 519 centimeters
	//
	// \- truck_weight = 5000 kg
	//
	// Please use the Distance Matrix Flexible version if you want to use custom truck
	// dimensions.
	//
	// Note: Only the "car" profile is enabled by default. Please note that customized
	// profiles (including "truck") might not be available for all regions. Please
	// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
	// representative or reach out at
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
	// additional profiles.
	//
	// Any of "car", "truck".
	Mode DistanceMatrixJsonGetParamsMode `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DistanceMatrixJsonGetParams]'s query parameters as
// `url.Values`.
func (r DistanceMatrixJsonGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// A semicolon-separated list indicating the side of the road from which the route
// will approach "destinations". When set to "unrestricted" a route can arrive at a
// destination from either side of the road. When set to "curb" the route will
// arrive at a destination on the driving side of the region. Please note the
// number of values provided must be equal to the number of "destinations".
// However, you can skip a coordinate and show its position in the list with the
// ";" separator. The values provided for the "approaches" parameter are effective
// for the "destinations" value at the same index. Example: "curb;;curb" will apply
// curbside restriction on the "destinations" points provided at the first and
// third index.
type DistanceMatrixJsonGetParamsApproaches string

const (
	DistanceMatrixJsonGetParamsApproachesUnrestricted DistanceMatrixJsonGetParamsApproaches = "unrestricted"
	DistanceMatrixJsonGetParamsApproachesCurb         DistanceMatrixJsonGetParamsApproaches = "curb"
)

// Setting this will ensure the route avoids ferries, tolls, highways or nothing.
// Multiple values should be separated by a pipe (|). If "none" is provided along
// with other values, an error is returned as a valid route is not feasible. Please
// note that when this parameter is not provided in the input, ferries are set to
// be avoided by default. When this parameter is provided, only the mentioned
// objects are avoided.
type DistanceMatrixJsonGetParamsAvoid string

const (
	DistanceMatrixJsonGetParamsAvoidToll    DistanceMatrixJsonGetParamsAvoid = "toll"
	DistanceMatrixJsonGetParamsAvoidFerry   DistanceMatrixJsonGetParamsAvoid = "ferry"
	DistanceMatrixJsonGetParamsAvoidHighway DistanceMatrixJsonGetParamsAvoid = "highway"
	DistanceMatrixJsonGetParamsAvoidNone    DistanceMatrixJsonGetParamsAvoid = "none"
)

// Set which driving mode the service should use to determine the "distance" and
// "duration" values. For example, if you use "car", the API will return the
// duration and distance of a route that a car can take. Using "truck" will return
// the same for a route a truck can use, taking into account appropriate truck
// routing restrictions.
//
// When "mode=truck", following are the default dimensions that are used:
//
// \- truck_height = 214 centimeters
//
// \- truck_width = 183 centimeters
//
// \- truck_length = 519 centimeters
//
// \- truck_weight = 5000 kg
//
// Please use the Distance Matrix Flexible version if you want to use custom truck
// dimensions.
//
// Note: Only the "car" profile is enabled by default. Please note that customized
// profiles (including "truck") might not be available for all regions. Please
// contact your [NextBillion.ai](http://NextBillion.ai) account manager, sales
// representative or reach out at
// [support@nextbillion.ai](mailto:support@nextbillion.ai) in case you need
// additional profiles.
type DistanceMatrixJsonGetParamsMode string

const (
	DistanceMatrixJsonGetParamsModeCar   DistanceMatrixJsonGetParamsMode = "car"
	DistanceMatrixJsonGetParamsModeTruck DistanceMatrixJsonGetParamsMode = "truck"
)
