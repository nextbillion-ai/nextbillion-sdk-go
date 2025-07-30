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

// MdmService contains methods and other services that help with interacting with
// the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMdmService] method instead.
type MdmService struct {
	Options []option.RequestOption
}

// NewMdmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMdmService(opts ...option.RequestOption) (r MdmService) {
	r = MdmService{}
	r.Options = opts
	return
}

// Create a massive distance matrix task
func (r *MdmService) NewDistanceMatrix(ctx context.Context, params MdmNewDistanceMatrixParams, opts ...option.RequestOption) (res *MdmNewDistanceMatrixResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "mdm/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get massive distance matrix task status
func (r *MdmService) GetDistanceMatrixStatus(ctx context.Context, query MdmGetDistanceMatrixStatusParams, opts ...option.RequestOption) (res *MdmGetDistanceMatrixStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "mdm/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MdmNewDistanceMatrixResponse struct {
	// A string indicating the state of the response. On successful responses, the
	// value will be Ok. Indicative error messages/codes are returned in case of
	// errors. See the [API Error Codes](#api-error-codes) section below for more
	// information.
	Code string `json:"code"`
	// Returns the error message in case a request fails. This field will not be
	// present in the response, if a request is successfully submitted.
	Message string `json:"message"`
	// A unique ID which can be used in the Asynchronous Distance Matrix GET method to
	// retrieve the final result.
	TaskID string `json:"task_id"`
	// Display the warnings, if any, for the given input parameters and values. In case
	// there are no warnings then this field would not be present in the response.
	Warning []string `json:"warning"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Message     respjson.Field
		TaskID      respjson.Field
		Warning     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MdmNewDistanceMatrixResponse) RawJSON() string { return r.JSON.raw }
func (r *MdmNewDistanceMatrixResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MdmGetDistanceMatrixStatusResponse struct {
	// A code representing the status of the request.
	//
	// Any of "Ok", "Processing", "Failed".
	Code MdmGetDistanceMatrixStatusResponseCode `json:"code"`
	// Returns the GCS result of a successful task. Please note that this is an
	// internal field.
	//
	// _internal field, the gcs result of specific task if task is success._
	OutputAddr string `json:"output_addr"`
	// Returns the link for the result file (csv format) once the task is completed
	// successfully.
	ResultLink string `json:"result_link"`
	// Returns the status detail of the result. Indicative error messages/codes are
	// returned in case of errors. See the [API Error Codes](#api-error-codes) section
	// below for more information.
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		OutputAddr  respjson.Field
		ResultLink  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MdmGetDistanceMatrixStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *MdmGetDistanceMatrixStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A code representing the status of the request.
type MdmGetDistanceMatrixStatusResponseCode string

const (
	MdmGetDistanceMatrixStatusResponseCodeOk         MdmGetDistanceMatrixStatusResponseCode = "Ok"
	MdmGetDistanceMatrixStatusResponseCodeProcessing MdmGetDistanceMatrixStatusResponseCode = "Processing"
	MdmGetDistanceMatrixStatusResponseCodeFailed     MdmGetDistanceMatrixStatusResponseCode = "Failed"
)

type MdmNewDistanceMatrixParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Use this option to switch to truck-specific routing or time based routing or if
	// you want to choose between the fastest and shortest route types.
	//
	// Any of "flexible".
	Option MdmNewDistanceMatrixParamsOption `query:"option,omitzero,required" json:"-"`
	// origins are the starting point of your route. Ensure that origins are routable
	// land locations. Multiple origins should be separated by a pipe symbol (|).
	//
	// **Format:** latitude_1,longitude_1|latitude_2,longitude_2|…
	Origins string `json:"origins,required"`
	// Specify if crossing an international border is expected for operations near
	// border areas. When set to false, the API will prohibit routes going back & forth
	// between countries. Consequently, routes within the same country will be
	// preferred if they are feasible for the given set of destination or waypoints .
	// When set to true, the routes will be allowed to go back & forth between
	// countries as needed.
	//
	// This feature is available in North America region only. Please get in touch with
	// [support@nextbillion.ai](mailto:support@nextbillion.ai) to enquire/enable other
	// areas.
	CrossBorder param.Opt[bool] `json:"cross_border,omitzero"`
	// This is a number in UNIX epoch timestamp in seconds format that can be used to
	// provide the departure time. The response will return the distance and duration
	// of the route based on typical traffic for at the given start time.If no input is
	// provided for this parameter then the traffic conditions at the time of making
	// the request are considered.
	//
	// Please note that when route_type is set to shortest then the departure_time will
	// be ineffective as the service will return the result for the shortest path
	// possible irrespective of the traffic conditions.
	DepartureTime param.Opt[int64] `json:"departure_time,omitzero"`
	// destinations are the ending coordinates of your route. Ensure that destinations
	// are routable land locations. Multiple destinations should be separated by a pipe
	// symbol (|).
	//
	// In case destinations are not provided or if it is left empty, then the input
	// value of origins will be copied to destinations to create the OD matrix pairs.
	//
	// **Format:** latitude_1,longitude_1|latitude_2,longitude_2|…
	Destinations param.Opt[string] `json:"destinations,omitzero"`
	// Specify the total load per axle (including the weight of trailers and shipped
	// goods) of the truck, in tonnes. When used, the service will return routes which
	// are legally allowed to carry the load specified per axle.
	//
	// Please note this parameter is effective only when mode=truck.
	TruckAxleLoad param.Opt[float64] `json:"truck_axle_load,omitzero"`
	// This defines the dimensions of a truck in centimeters (cm) in the format of
	// "height,width,length". This parameter is effective only when mode=truck and
	// option=flexible. Maximum dimensions are as follows:
	//
	// Height = 1000 cm Width = 5000 cm Length = 5000 cm
	TruckSize param.Opt[string] `json:"truck_size,omitzero"`
	// This parameter defines the weight of the truck including trailers and shipped
	// goods in kilograms (kg). This parameter is effective only when mode=truck and
	// option=flexible.
	TruckWeight param.Opt[int64] `json:"truck_weight,omitzero"`
	// Specify a spliter to split the matrix by. It accepts 2 values:
	//
	// - od_number_spliter:
	//
	// - straight_distance_spliter:
	//
	// Please note it is an internal, debug field only.
	//
	// _debug field. choose specific spliter to split matrix._
	//
	// Any of "od_number_spliter", "straight_distance_spliter", "location_spliter".
	Spliter MdmNewDistanceMatrixParamsSpliter `query:"spliter,omitzero" json:"-"`
	// Provide the country that the coordinates belong to.
	//
	// _the input coordinates area._
	//
	// Any of "singapore", "usa", "india".
	Area MdmNewDistanceMatrixParamsArea `json:"area,omitzero"`
	// Setting this will ensure the route avoids the object(s) specified as input.
	// Multiple values should be separated by a pipe (|). If none is provided along
	// with other values, an error is returned as a valid route is not feasible.
	//
	// - **Note:**
	//
	//   - This parameter is effective only when route_type=fastest.
	//
	//   - When this parameter is not provided in the input, ferries are set to be
	//     avoided by default. When avoid input is provided, only the mentioned objects
	//     are avoided.
	//
	//   - When using avoid=bbox users also need to specify the boundaries of the
	//     bounding box to be avoid. Multiple bounding boxes can be specified
	//     simultaneously. Please note that bounding box is a hard filter and if it
	//     blocks all possible routes between given locations, a 4xx error is returned.
	//
	//   - **Format:** bbox: min_latitude,min_longtitude,max_latitude,max_longitude.
	//
	//   - **Example:** avoid=bbox: 34.0635,-118.2547, 34.0679,-118.2478 | bbox:
	//     34.0521,-118.2342, 34.0478,-118.2437
	//
	//   - When using avoid=sharp_turn, default range of permissible turn angles is
	//     \[120,240\].
	//
	// Any of "toll", "ferry", "highway", "sharp_turn", "service_road", "bbox",
	// "left_turn", "right_turn", "none".
	Avoid MdmNewDistanceMatrixParamsAvoid `json:"avoid,omitzero"`
	// Specify the side of the road from which to approach destinations points. Please
	// note that the given approach will be applied to all the destinations.
	//
	// Any of "unrestricted", "curb".
	DestinationsApproach MdmNewDistanceMatrixParamsDestinationsApproach `json:"destinations_approach,omitzero"`
	// Specify the type of hazardous material being carried and the service will avoid
	// roads which are not suitable for the type of goods specified. Multiple values
	// can be separated using a pipe operator | .
	//
	// Please note that this parameter is effective only when mode=truck.
	//
	// Any of "general", "circumstantial", "explosive", "harmful_to_water".
	HazmatType MdmNewDistanceMatrixParamsHazmatType `json:"hazmat_type,omitzero"`
	// Set which driving mode the service should use to determine a route.
	//
	// For example, if you use car, the API will return a route that a car can take.
	// Using truck will return a route a truck can use, taking into account appropriate
	// truck routing restrictions.
	//
	// Any of "car", "truck".
	Mode MdmNewDistanceMatrixParamsMode `json:"mode,omitzero"`
	// Specify the side of the road from which to approach origins points. Please note
	// that the given approach will be applied to all the points provided as origins.
	//
	// Any of "unrestricted", "curb".
	OriginsApproach MdmNewDistanceMatrixParamsOriginsApproach `json:"origins_approach,omitzero"`
	// Set the route type that needs to be returned. Please note that route_type is
	// effective only when option=flexible.
	//
	// Any of "fastest", "shortest".
	RouteType MdmNewDistanceMatrixParamsRouteType `json:"route_type,omitzero"`
	paramObj
}

func (r MdmNewDistanceMatrixParams) MarshalJSON() (data []byte, err error) {
	type shadow MdmNewDistanceMatrixParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MdmNewDistanceMatrixParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MdmNewDistanceMatrixParams]'s query parameters as
// `url.Values`.
func (r MdmNewDistanceMatrixParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Use this option to switch to truck-specific routing or time based routing or if
// you want to choose between the fastest and shortest route types.
type MdmNewDistanceMatrixParamsOption string

const (
	MdmNewDistanceMatrixParamsOptionFlexible MdmNewDistanceMatrixParamsOption = "flexible"
)

// Specify a spliter to split the matrix by. It accepts 2 values:
//
// - od_number_spliter:
//
// - straight_distance_spliter:
//
// Please note it is an internal, debug field only.
//
// _debug field. choose specific spliter to split matrix._
type MdmNewDistanceMatrixParamsSpliter string

const (
	MdmNewDistanceMatrixParamsSpliterOdNumberSpliter         MdmNewDistanceMatrixParamsSpliter = "od_number_spliter"
	MdmNewDistanceMatrixParamsSpliterStraightDistanceSpliter MdmNewDistanceMatrixParamsSpliter = "straight_distance_spliter"
	MdmNewDistanceMatrixParamsSpliterLocationSpliter         MdmNewDistanceMatrixParamsSpliter = "location_spliter"
)

// Provide the country that the coordinates belong to.
//
// _the input coordinates area._
type MdmNewDistanceMatrixParamsArea string

const (
	MdmNewDistanceMatrixParamsAreaSingapore MdmNewDistanceMatrixParamsArea = "singapore"
	MdmNewDistanceMatrixParamsAreaUsa       MdmNewDistanceMatrixParamsArea = "usa"
	MdmNewDistanceMatrixParamsAreaIndia     MdmNewDistanceMatrixParamsArea = "india"
)

// Setting this will ensure the route avoids the object(s) specified as input.
// Multiple values should be separated by a pipe (|). If none is provided along
// with other values, an error is returned as a valid route is not feasible.
//
// - **Note:**
//
//   - This parameter is effective only when route_type=fastest.
//
//   - When this parameter is not provided in the input, ferries are set to be
//     avoided by default. When avoid input is provided, only the mentioned objects
//     are avoided.
//
//   - When using avoid=bbox users also need to specify the boundaries of the
//     bounding box to be avoid. Multiple bounding boxes can be specified
//     simultaneously. Please note that bounding box is a hard filter and if it
//     blocks all possible routes between given locations, a 4xx error is returned.
//
//   - **Format:** bbox: min_latitude,min_longtitude,max_latitude,max_longitude.
//
//   - **Example:** avoid=bbox: 34.0635,-118.2547, 34.0679,-118.2478 | bbox:
//     34.0521,-118.2342, 34.0478,-118.2437
//
//   - When using avoid=sharp_turn, default range of permissible turn angles is
//     \[120,240\].
type MdmNewDistanceMatrixParamsAvoid string

const (
	MdmNewDistanceMatrixParamsAvoidToll        MdmNewDistanceMatrixParamsAvoid = "toll"
	MdmNewDistanceMatrixParamsAvoidFerry       MdmNewDistanceMatrixParamsAvoid = "ferry"
	MdmNewDistanceMatrixParamsAvoidHighway     MdmNewDistanceMatrixParamsAvoid = "highway"
	MdmNewDistanceMatrixParamsAvoidSharpTurn   MdmNewDistanceMatrixParamsAvoid = "sharp_turn"
	MdmNewDistanceMatrixParamsAvoidServiceRoad MdmNewDistanceMatrixParamsAvoid = "service_road"
	MdmNewDistanceMatrixParamsAvoidBbox        MdmNewDistanceMatrixParamsAvoid = "bbox"
	MdmNewDistanceMatrixParamsAvoidLeftTurn    MdmNewDistanceMatrixParamsAvoid = "left_turn"
	MdmNewDistanceMatrixParamsAvoidRightTurn   MdmNewDistanceMatrixParamsAvoid = "right_turn"
	MdmNewDistanceMatrixParamsAvoidNone        MdmNewDistanceMatrixParamsAvoid = "none"
)

// Specify the side of the road from which to approach destinations points. Please
// note that the given approach will be applied to all the destinations.
type MdmNewDistanceMatrixParamsDestinationsApproach string

const (
	MdmNewDistanceMatrixParamsDestinationsApproachUnrestricted MdmNewDistanceMatrixParamsDestinationsApproach = "unrestricted"
	MdmNewDistanceMatrixParamsDestinationsApproachCurb         MdmNewDistanceMatrixParamsDestinationsApproach = "curb"
)

// Specify the type of hazardous material being carried and the service will avoid
// roads which are not suitable for the type of goods specified. Multiple values
// can be separated using a pipe operator | .
//
// Please note that this parameter is effective only when mode=truck.
type MdmNewDistanceMatrixParamsHazmatType string

const (
	MdmNewDistanceMatrixParamsHazmatTypeGeneral        MdmNewDistanceMatrixParamsHazmatType = "general"
	MdmNewDistanceMatrixParamsHazmatTypeCircumstantial MdmNewDistanceMatrixParamsHazmatType = "circumstantial"
	MdmNewDistanceMatrixParamsHazmatTypeExplosive      MdmNewDistanceMatrixParamsHazmatType = "explosive"
	MdmNewDistanceMatrixParamsHazmatTypeHarmfulToWater MdmNewDistanceMatrixParamsHazmatType = "harmful_to_water"
)

// Set which driving mode the service should use to determine a route.
//
// For example, if you use car, the API will return a route that a car can take.
// Using truck will return a route a truck can use, taking into account appropriate
// truck routing restrictions.
type MdmNewDistanceMatrixParamsMode string

const (
	MdmNewDistanceMatrixParamsModeCar   MdmNewDistanceMatrixParamsMode = "car"
	MdmNewDistanceMatrixParamsModeTruck MdmNewDistanceMatrixParamsMode = "truck"
)

// Specify the side of the road from which to approach origins points. Please note
// that the given approach will be applied to all the points provided as origins.
type MdmNewDistanceMatrixParamsOriginsApproach string

const (
	MdmNewDistanceMatrixParamsOriginsApproachUnrestricted MdmNewDistanceMatrixParamsOriginsApproach = "unrestricted"
	MdmNewDistanceMatrixParamsOriginsApproachCurb         MdmNewDistanceMatrixParamsOriginsApproach = "curb"
)

// Set the route type that needs to be returned. Please note that route_type is
// effective only when option=flexible.
type MdmNewDistanceMatrixParamsRouteType string

const (
	MdmNewDistanceMatrixParamsRouteTypeFastest  MdmNewDistanceMatrixParamsRouteType = "fastest"
	MdmNewDistanceMatrixParamsRouteTypeShortest MdmNewDistanceMatrixParamsRouteType = "shortest"
)

type MdmGetDistanceMatrixStatusParams struct {
	// Provide the unique ID that was returned on successful submission of the
	// Asynchronous Distance Matrix POST request.
	ID string `query:"id,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	paramObj
}

// URLQuery serializes [MdmGetDistanceMatrixStatusParams]'s query parameters as
// `url.Values`.
func (r MdmGetDistanceMatrixStatusParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
