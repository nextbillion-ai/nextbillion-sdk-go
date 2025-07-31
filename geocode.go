// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apiquery"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// GeocodeService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeocodeService] method instead.
type GeocodeService struct {
	Options []option.RequestOption
}

// NewGeocodeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGeocodeService(opts ...option.RequestOption) (r GeocodeService) {
	r = GeocodeService{}
	r.Options = opts
	return
}

// Geocode
func (r *GeocodeService) Get(ctx context.Context, query GeocodeGetParams, opts ...option.RequestOption) (res *GeocodeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Batch Geocode
func (r *GeocodeService) BatchNew(ctx context.Context, params GeocodeBatchNewParams, opts ...option.RequestOption) (res *GeocodeBatchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geocode/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Structured Geocode
func (r *GeocodeService) StructuredGet(ctx context.Context, query GeocodeStructuredGetParams, opts ...option.RequestOption) (res *GeocodeStructuredGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "geocode/structured"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// An array returning the location coordinates of all the access points of the
// search result.
type Access struct {
	// The latitude of the access point of the search result.
	Lat float64 `json:"lat"`
	// The longitude of the access point of the search result.
	Lng float64 `json:"lng"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Access) RawJSON() string { return r.JSON.raw }
func (r *Access) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Postal address of the result item.
type Address struct {
	// The name of the primary locality of the place.
	City string `json:"city"`
	// A three-letter country code.
	CountryCode string `json:"countryCode"`
	// The localised country name.
	CountryName string `json:"countryName"`
	// A division of a state; typically, a secondary-level administrative division of a
	// country or equivalent.
	County string `json:"county"`
	// A division of city; typically an administrative unit within a larger city or a
	// customary name of a city's neighborhood.
	District string `json:"district"`
	// House number of the returned place, if available.
	HouseNumber string `json:"houseNumber"`
	// Assembled address value built out of the address components according to the
	// regional postal rules. These are the same rules for all endpoints. It may not
	// include all the input terms.
	Label string `json:"label"`
	// An alphanumeric string included in a postal address to facilitate mail sorting,
	// such as post code, postcode, or ZIP code.
	PostalCode string `json:"postalCode"`
	// The state division of a country.
	State string `json:"state"`
	// A country specific state code or state name abbreviation. For example, in the
	// United States it is the two letter state abbreviation: "CA" for California.
	StateCode string `json:"stateCode"`
	// Name of street of the returned place, if available.
	Street string `json:"street"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		City        respjson.Field
		CountryCode respjson.Field
		CountryName respjson.Field
		County      respjson.Field
		District    respjson.Field
		HouseNumber respjson.Field
		Label       respjson.Field
		PostalCode  respjson.Field
		State       respjson.Field
		StateCode   respjson.Field
		Street      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Address) RawJSON() string { return r.JSON.raw }
func (r *Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Categories struct {
	// Identifier number for an associated category.
	ID string `json:"id"`
	// Name of the place category in the result item language.
	Name string `json:"name"`
	// Whether or not it is a primary category. This field is visible only when the
	// value is 'true'.
	Primary string `json:"primary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Primary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Categories) RawJSON() string { return r.JSON.raw }
func (r *Categories) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactObject struct {
	// The list of place categories this contact refers to.
	Categories []ContactObjectCategory `json:"categories"`
	// Optional label for the contact string, such as "Customer Service" or "Pharmacy
	// Fax".
	Label string `json:"label"`
	// Contact information, as specified by the contact type.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories  respjson.Field
		Label       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactObject) RawJSON() string { return r.JSON.raw }
func (r *ContactObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContactObjectCategory struct {
	// Identifier number for an associated category. For example: "900-9300-0000"
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContactObjectCategory) RawJSON() string { return r.JSON.raw }
func (r *ContactObjectCategory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Contacts struct {
	Email    []ContactObject `json:"email"`
	Fax      []ContactObject `json:"fax"`
	Mobile   []ContactObject `json:"mobile"`
	Phone    []ContactObject `json:"phone"`
	TollFree []ContactObject `json:"tollFree"`
	Www      []ContactObject `json:"www"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Fax         respjson.Field
		Mobile      respjson.Field
		Phone       respjson.Field
		TollFree    respjson.Field
		Www         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Contacts) RawJSON() string { return r.JSON.raw }
func (r *Contacts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The bounding box enclosing the geometric shape (area or line) that an individual
// result covers. place typed results have no mapView.
type MapView struct {
	// Longitude of the eastern-side of the box.
	East string `json:"east"`
	// Longitude of the northern-side of the box.
	North string `json:"north"`
	// Longitude of the southern-side of the box.
	South string `json:"south"`
	// Longitude of the western-side of the box.
	West string `json:"west"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		East        respjson.Field
		North       respjson.Field
		South       respjson.Field
		West        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MapView) RawJSON() string { return r.JSON.raw }
func (r *MapView) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the location coordinates of the result.
type Position struct {
	// The latitude of the searched place.
	Lat string `json:"lat"`
	// The longitude of the searched place.
	Lng string `json:"lng"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Position) RawJSON() string { return r.JSON.raw }
func (r *Position) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeGetResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []GeocodeGetResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeGetResponseItem struct {
	// The unique identifier for the result item.
	ID string `json:"id"`
	// An array returning the location coordinates of all the access points of the
	// search result.
	Access Access `json:"access"`
	// Postal address of the result item.
	Address Address `json:"address"`
	// The list of categories assigned to this place.
	Categories []Categories `json:"categories"`
	// Contact information like phone, email or website.
	Contacts []Contacts `json:"contacts"`
	// The distance "as the crow flies" from the search center to this result item in
	// meters.
	Distance int64 `json:"distance"`
	// The bounding box enclosing the geometric shape (area or line) that an individual
	// result covers. place typed results have no mapView.
	MapView MapView `json:"mapView"`
	// Returns the operating hours of the place, if available.
	OpeningHours GeocodeGetResponseItemOpeningHours `json:"openingHours"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// Score of the result. A higher score indicates a closer match.
	Scoring GeocodeGetResponseItemScoring `json:"scoring"`
	// The localized display name of this result item.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Access       respjson.Field
		Address      respjson.Field
		Categories   respjson.Field
		Contacts     respjson.Field
		Distance     respjson.Field
		MapView      respjson.Field
		OpeningHours respjson.Field
		Position     respjson.Field
		Scoring      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the operating hours of the place, if available.
type GeocodeGetResponseItemOpeningHours struct {
	// A collection of attributes with details about the opening and closing hours for
	// each day of the week.
	TimeRanges []GeocodeGetResponseItemOpeningHoursTimeRange `json:"timeRanges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeRanges  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItemOpeningHours) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItemOpeningHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeGetResponseItemOpeningHoursTimeRange struct {
	// Returns the closing time details.
	EndTime GeocodeGetResponseItemOpeningHoursTimeRangeEndTime `json:"endTime"`
	// Returns the open time details.
	StartTime GeocodeGetResponseItemOpeningHoursTimeRangeStartTime `json:"startTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItemOpeningHoursTimeRange) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItemOpeningHoursTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the closing time details.
type GeocodeGetResponseItemOpeningHoursTimeRangeEndTime struct {
	// The date to which the subsequent closing time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place closes.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place closes.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItemOpeningHoursTimeRangeEndTime) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItemOpeningHoursTimeRangeEndTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the open time details.
type GeocodeGetResponseItemOpeningHoursTimeRangeStartTime struct {
	// The date to which the subsequent open time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place opens.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place opens.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItemOpeningHoursTimeRangeStartTime) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItemOpeningHoursTimeRangeStartTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Score of the result. A higher score indicates a closer match.
type GeocodeGetResponseItemScoring struct {
	// A breakdown of how closely individual field of the result matched with the
	// provided query q.
	FieldScore any `json:"fieldScore"`
	// A score, out of 1, indicating how closely the result matches with the provided
	// query q .
	QueryScore float64 `json:"queryScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldScore  respjson.Field
		QueryScore  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeGetResponseItemScoring) RawJSON() string { return r.JSON.raw }
func (r *GeocodeGetResponseItemScoring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeBatchNewResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []GeocodeBatchNewResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeBatchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *GeocodeBatchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeBatchNewResponseItem struct {
	// The unique identifier for the result item.
	ID string `json:"id"`
	// An array returning the location coordinates of all the access points of the
	// search result.
	Access Access `json:"access"`
	// Postal address of the result item.
	Address Address `json:"address"`
	// The list of categories assigned to this place.
	Categories []Categories `json:"categories"`
	// Contact information like phone, email or website.
	Contacts []Contacts `json:"contacts"`
	// The distance "as the crow flies" from the search center to this result item in
	// meters.
	Distance int64 `json:"distance"`
	// The bounding box enclosing the geometric shape (area or line) that an individual
	// result covers. place typed results have no mapView.
	MapView MapView `json:"mapView"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// Score of the result. A higher score indicates a closer match.
	Scoring GeocodeBatchNewResponseItemScoring `json:"scoring"`
	// The localized display name of this result item.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Access      respjson.Field
		Address     respjson.Field
		Categories  respjson.Field
		Contacts    respjson.Field
		Distance    respjson.Field
		MapView     respjson.Field
		Position    respjson.Field
		Scoring     respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeBatchNewResponseItem) RawJSON() string { return r.JSON.raw }
func (r *GeocodeBatchNewResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Score of the result. A higher score indicates a closer match.
type GeocodeBatchNewResponseItemScoring struct {
	// A breakdown of how closely individual field of the result matched with the
	// provided query q.
	FieldScore any `json:"fieldScore"`
	// A score, out of 1, indicating how closely the result matches with the provided
	// query q .
	QueryScore float64 `json:"queryScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldScore  respjson.Field
		QueryScore  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeBatchNewResponseItemScoring) RawJSON() string { return r.JSON.raw }
func (r *GeocodeBatchNewResponseItemScoring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeStructuredGetResponse struct {
	// The results are presented as a JSON list of candidates in ranked order
	// (most-likely to least-likely) based on the matched location criteria.
	Items []GeocodeStructuredGetResponseItem `json:"items"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponse) RawJSON() string { return r.JSON.raw }
func (r *GeocodeStructuredGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeStructuredGetResponseItem struct {
	// The unique identifier for the result item.
	ID string `json:"id"`
	// An array returning the location coordinates of all the access points of the
	// search result.
	Access Access `json:"access"`
	// Postal address of the result item.
	Address Address `json:"address"`
	// The list of categories assigned to this place.
	Categories []Categories `json:"categories"`
	// Contact information like phone, email or website.
	Contacts []Contacts `json:"contacts"`
	// The distance "as the crow flies" from the search center to this result item in
	// meters.
	Distance int64 `json:"distance"`
	// The bounding box enclosing the geometric shape (area or line) that an individual
	// result covers. place typed results have no mapView.
	MapView MapView `json:"mapView"`
	// Returns the operating hours of the place, if available.
	OpeningHours GeocodeStructuredGetResponseItemOpeningHours `json:"openingHours"`
	// Returns the location coordinates of the result.
	Position Position `json:"position"`
	// Score of the result. A higher score indicates a closer match.
	Scoring GeocodeStructuredGetResponseItemScoring `json:"scoring"`
	// The localized display name of this result item.
	Title string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Access       respjson.Field
		Address      respjson.Field
		Categories   respjson.Field
		Contacts     respjson.Field
		Distance     respjson.Field
		MapView      respjson.Field
		OpeningHours respjson.Field
		Position     respjson.Field
		Scoring      respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItem) RawJSON() string { return r.JSON.raw }
func (r *GeocodeStructuredGetResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the operating hours of the place, if available.
type GeocodeStructuredGetResponseItemOpeningHours struct {
	// A collection of attributes with details about the opening and closing hours for
	// each day of the week.
	TimeRanges []GeocodeStructuredGetResponseItemOpeningHoursTimeRange `json:"timeRanges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimeRanges  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItemOpeningHours) RawJSON() string { return r.JSON.raw }
func (r *GeocodeStructuredGetResponseItemOpeningHours) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeStructuredGetResponseItemOpeningHoursTimeRange struct {
	// Returns the closing time details.
	EndTime GeocodeStructuredGetResponseItemOpeningHoursTimeRangeEndTime `json:"endTime"`
	// Returns the open time details.
	StartTime GeocodeStructuredGetResponseItemOpeningHoursTimeRangeStartTime `json:"startTime"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItemOpeningHoursTimeRange) RawJSON() string { return r.JSON.raw }
func (r *GeocodeStructuredGetResponseItemOpeningHoursTimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the closing time details.
type GeocodeStructuredGetResponseItemOpeningHoursTimeRangeEndTime struct {
	// The date to which the subsequent closing time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place closes.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place closes.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItemOpeningHoursTimeRangeEndTime) RawJSON() string {
	return r.JSON.raw
}
func (r *GeocodeStructuredGetResponseItemOpeningHoursTimeRangeEndTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Returns the open time details.
type GeocodeStructuredGetResponseItemOpeningHoursTimeRangeStartTime struct {
	// The date to which the subsequent open time details belong to.
	Date string `json:"date"`
	// The hour of the day when the place opens.
	Hour int64 `json:"hour"`
	// The minute of the hour when the place opens.
	Minute int64 `json:"minute"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date        respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItemOpeningHoursTimeRangeStartTime) RawJSON() string {
	return r.JSON.raw
}
func (r *GeocodeStructuredGetResponseItemOpeningHoursTimeRangeStartTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Score of the result. A higher score indicates a closer match.
type GeocodeStructuredGetResponseItemScoring struct {
	// A breakdown of how closely individual field of the result matched with the
	// provided query q.
	FieldScore any `json:"fieldScore"`
	// A score, out of 1, indicating how closely the result matches with the provided
	// query q .
	QueryScore float64 `json:"queryScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldScore  respjson.Field
		QueryScore  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeocodeStructuredGetResponseItemScoring) RawJSON() string { return r.JSON.raw }
func (r *GeocodeStructuredGetResponseItemScoring) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeGetParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the free-text search query.
	//
	// Please note that whitespace, urls, email addresses, or other out-of-scope
	// queries will yield no results.
	Q string `query:"q,required" json:"-"`
	// Specify the center of the search context expressed as coordinates.
	//
	// Please note that one of "at", "in=circle" or "in=bbox" should be provided for
	// relevant results.
	At param.Opt[string] `query:"at,omitzero" format:"latitude,longitude" json:"-"`
	// Search within a geographic area. This is a hard filter. Results will be returned
	// if they are located within the specified area.
	//
	// # A geographic area can be
	//
	//   - a country (or multiple countries), provided as comma-separated
	//     [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	//     codes
	//
	//     The country codes are to be provided in all uppercase.
	//
	//     Format: countryCode:{countryCode}[,{countryCode}]
	//
	//   - a circular area, provided as latitude, longitude, and radius (an integer with
	//     meters as unit)
	//
	//     Format: circle:{latitude},{longitude};r={radius}
	//
	//   - a bounding box, provided as _west longitude_, _south latitude_, _east
	//     longitude_, _north latitude_
	//
	//     Format: bbox:{west longitude},{south latitude},{east longitude},{north
	//     latitude}
	//
	// Please provide one of 'at', 'in=circle' or 'in=bbox' input for a relevant
	// result.
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	// Select the language to be used for result rendering from a list of
	// [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language
	// codes.
	Lang param.Opt[string] `query:"lang,omitzero" json:"-"`
	// Sets the maximum number of results to be returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeocodeGetParams]'s query parameters as `url.Values`.
func (r GeocodeGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeocodeBatchNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key  string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	Body []GeocodeBatchNewParamsBody
	paramObj
}

func (r GeocodeBatchNewParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *GeocodeBatchNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [GeocodeBatchNewParams]'s query parameters as `url.Values`.
func (r GeocodeBatchNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The property Q is required.
type GeocodeBatchNewParamsBody struct {
	// Specify the free-text search query. Please note that whitespace, urls, email
	// addresses, or other out-of-scope queries will yield no results.
	Q string `json:"q,required"`
	// Specify the center of the search context expressed as coordinates.
	//
	// Please note that one of "at", "in=circle" or "in=bbox" should be provided for
	// relevant results.
	At param.Opt[string] `json:"at,omitzero" format:"latitude,longitude"`
	// Search within a geographic area. This is a hard filter. Results will be returned
	// if they are located within the specified area.
	//
	// # A geographic area can be
	//
	//   - a country (or multiple countries), provided as comma-separated
	//     [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	//     codes
	//
	//     The country codes are to be provided in all uppercase.
	//
	//     Format: countryCode:{countryCode}[,{countryCode}]
	//
	//   - a circular area, provided as latitude, longitude, and radius (an integer with
	//     meters as unit)
	//
	//     Format: circle:{latitude},{longitude};r={radius}
	//
	//   - a bounding box, provided as _west longitude_, _south latitude_, _east
	//     longitude_, _north latitude_
	//
	//     Format: bbox:{west longitude},{south latitude},{east longitude},{north
	//     latitude}
	//
	// Please provide one of 'at', 'in=circle' or 'in=bbox' input for a relevant
	// result.
	In param.Opt[string] `json:"in,omitzero"`
	// Select the language to be used for result rendering from a list of
	// [BCP 47](https://en.wikipedia.org/wiki/IETF_language_tag) compliant language
	// codes.
	Lang param.Opt[string] `json:"lang,omitzero"`
	// Maximum number of results to be returned. Please note that the minimum value
	// that can be provided is 1 and the maximum that can be provided is 100.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	paramObj
}

func (r GeocodeBatchNewParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow GeocodeBatchNewParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GeocodeBatchNewParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeocodeStructuredGetParams struct {
	// Specify a valid
	// [ISO 3166-1 alpha-3](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3) country
	// code in which the place being searched should be located. Please note that this
	// is a case-sensitive field and the country code should be in all uppercase.
	CountryCode string `query:"countryCode,required" json:"-"`
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	Key string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the center of the search context expressed as coordinates.
	//
	// Please note that one of "at", "in=circle" or "in=bbox" should be provided for
	// relevant results.
	At param.Opt[string] `query:"at,omitzero" format:"latitude,longitude" json:"-"`
	// Specify the city in which the place being searched should be located.
	City param.Opt[string] `query:"city,omitzero" json:"-"`
	// Specify the county division of the state in which the place being searched
	// should be located.
	County param.Opt[string] `query:"county,omitzero" json:"-"`
	// Specify the house number of the place being searched.
	HouseNumber param.Opt[string] `query:"houseNumber,omitzero" json:"-"`
	// Search within a geographic area. This is a hard filter. Results will be returned
	// if they are located within the specified area.
	//
	// # A geographic area can be
	//
	//   - a circular area, provided as latitude, longitude, and radius (an integer with
	//     meters as unit)
	//
	//     Format: circle:{latitude},{longitude};r={radius}
	//
	//   - a bounding box, provided as _west longitude_, _south latitude_, _east
	//     longitude_, _north latitude_
	//
	//     Format: bbox:{west longitude},{south latitude},{east longitude},{north
	//     latitude}
	//
	// Please provide one of 'at', 'in=circle' or 'in=bbox' input for a relevant
	// result.
	In param.Opt[string] `query:"in,omitzero" json:"-"`
	// Sets the maximum number of results to be returned.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Specify the postal code in which the place being searched should be located.
	PostalCode param.Opt[string] `query:"postalCode,omitzero" json:"-"`
	// Specify the state division of the country in which the place being searched
	// should be located.
	State param.Opt[string] `query:"state,omitzero" json:"-"`
	// Specify the name of the street in which the place being searched should be
	// located.
	Street param.Opt[string] `query:"street,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GeocodeStructuredGetParams]'s query parameters as
// `url.Values`.
func (r GeocodeStructuredGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
