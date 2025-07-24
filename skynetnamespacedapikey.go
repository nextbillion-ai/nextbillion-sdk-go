// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apijson"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/apiquery"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/param"
	"github.com/stainless-sdks/nextbillion-sdk-go/packages/respjson"
)

// SkynetNamespacedApikeyService contains methods and other services that help with
// interacting with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetNamespacedApikeyService] method instead.
type SkynetNamespacedApikeyService struct {
	Options []option.RequestOption
}

// NewSkynetNamespacedApikeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSkynetNamespacedApikeyService(opts ...option.RequestOption) (r SkynetNamespacedApikeyService) {
	r = SkynetNamespacedApikeyService{}
	r.Options = opts
	return
}

// Create namespace under a parent key
func (r *SkynetNamespacedApikeyService) New(ctx context.Context, body SkynetNamespacedApikeyNewParams, opts ...option.RequestOption) (res *SkynetNamespacedApikeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/namespaced-apikeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete namespace under a parent key
func (r *SkynetNamespacedApikeyService) Delete(ctx context.Context, body SkynetNamespacedApikeyDeleteParams, opts ...option.RequestOption) (res *SkynetNamespacedApikeyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "skynet/namespaced-apikeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type SkynetNamespacedApikeyNewResponse struct {
	// Returns the error type in case of any error. If there is no error, then this
	// field is absent in the response.
	Error string `json:"error"`
	// Returns the error message in case of any error. If there is no error, then this
	// field is absent in the response.
	Message string `json:"message"`
	// An object to return the details about the namespace key created.
	Result SkynetNamespacedApikeyNewResponseResult `json:"result"`
	// Returns the API response code.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Message     respjson.Field
		Result      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetNamespacedApikeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetNamespacedApikeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object to return the details about the namespace key created.
type SkynetNamespacedApikeyNewResponseResult struct {
	// Returns the unique `key` created for the specified namespace.
	Apikey string `json:"apikey"`
	// Returns the time, expressed as UNIX epoch timestamp in seconds, when the
	// namespace key was created.
	CreatedAt int64 `json:"created_at"`
	// Returns the time, expressed as UNIX epoch timestamp in seconds, when the
	// namespace key will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Returns the name of the `namespace` for which the key is created.
	Namespace string `json:"namespace"`
	// An internal subscription ID.
	SubID string `json:"sub_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Apikey      respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Namespace   respjson.Field
		SubID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetNamespacedApikeyNewResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SkynetNamespacedApikeyNewResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetNamespacedApikeyDeleteResponse struct {
	// Its value is `OK` in case of a successful delete operation. Indicative error
	// messages are returned otherwise, for different errors.
	Msg string `json:"msg"`
	// A string indicating the state of the response. A successful delete operation ins
	// indicated by an HTTP code of`200`. See the
	// [API Error Codes](https://docs.nextbillion.ai/docs/tracking/api/live-tracking-api#api-error-codes)
	// section below for possible values in case of errors.
	Status int64 `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Msg         respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetNamespacedApikeyDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetNamespacedApikeyDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetNamespacedApikeyNewParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API.
	QueryKey1 string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify a name for the `namespace`. If the namespace specified is unique then a
	// new namespace along with a new key is created. Whereas if the specified
	// `namespace` is not unique, a new key will be created in the existing
	// `namespace`. Please note that a `namespace` cannot be created using another
	// namespace key.
	Namespace string            `query:"namespace,required" json:"-"`
	QueryKey2 param.Opt[string] `query:"{key},omitzero" json:"-"`
	paramObj
}

func (r SkynetNamespacedApikeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetNamespacedApikeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetNamespacedApikeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SkynetNamespacedApikeyNewParams]'s query parameters as
// `url.Values`.
func (r SkynetNamespacedApikeyNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SkynetNamespacedApikeyDeleteParams struct {
	// A key is a unique identifier that is required to authenticate a request to the
	// API. Please note for the delete namespace key operation another namespace key
	// cannot be used.
	//
	// The namespace created using this key can be managed using the APIs & Services >
	// Credentials section of user’s
	// [NextBillion Console](https://console.nextbillion.ai).
	QueryKey1 string `query:"key,required" format:"32 character alphanumeric string" json:"-"`
	// Specify the key to be deleted.
	KeyToDelete string `query:"key_to_delete,required" json:"-"`
	// Specify the name of the `namespace` to which the \`key_to_delete\` belongs.
	// Please note that a namespace key cannot be deleted using another namespace key.
	Namespace string            `query:"namespace,required" json:"-"`
	QueryKey2 param.Opt[string] `query:"{key},omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SkynetNamespacedApikeyDeleteParams]'s query parameters as
// `url.Values`.
func (r SkynetNamespacedApikeyDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
