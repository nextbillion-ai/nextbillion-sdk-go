// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai

import (
	"context"
	"net/http"
	"slices"

	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/apijson"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/requestconfig"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/param"
	"github.com/nextbillion-ai/nextbillion-sdk-go/packages/respjson"
)

// SkynetService contains methods and other services that help with interacting
// with the nextbillion-sdk API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSkynetService] method instead.
type SkynetService struct {
	Options           []option.RequestOption
	Asset             SkynetAssetService
	Monitor           SkynetMonitorService
	Trip              SkynetTripService
	NamespacedApikeys SkynetNamespacedApikeyService
	Config            SkynetConfigService
	Search            SkynetSearchService
}

// NewSkynetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSkynetService(opts ...option.RequestOption) (r SkynetService) {
	r = SkynetService{}
	r.Options = opts
	r.Asset = NewSkynetAssetService(opts...)
	r.Monitor = NewSkynetMonitorService(opts...)
	r.Trip = NewSkynetTripService(opts...)
	r.NamespacedApikeys = NewSkynetNamespacedApikeyService(opts...)
	r.Config = NewSkynetConfigService(opts...)
	r.Search = NewSkynetSearchService(opts...)
	return
}

// POST Action
func (r *SkynetService) Subscribe(ctx context.Context, body SkynetSubscribeParams, opts ...option.RequestOption) (res *SkynetSubscribeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "skynet/subscribe"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SkynetSubscribeResponse struct {
	// Subscription ID as provided in the input action message.
	ID string `json:"id"`
	// Returns the error message when status: error. Otherwise, response doesn't
	// contain this field.
	Error string `json:"error"`
	// Status of the action. It can have only two values - "success" or "error".
	Status string `json:"status"`
	// Returns the UNIX timestamp, in milliseconds format, when the web-socket returns
	// the action response.
	Timestamp int64 `json:"timestamp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Error       respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkynetSubscribeResponse) RawJSON() string { return r.JSON.raw }
func (r *SkynetSubscribeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SkynetSubscribeParams struct {
	// Specify the behavior that needs to be achieved for the subscription. Following
	// values are accepted:
	//
	//   - TRIP_SUBSCRIBE: Enable a trip subscription.
	//   - TRIP_UNSUBSCRIBE: Unsubscribe from a trip
	//   - HEARTBEAT: Enable heartbeat mechanism for a web-socket connection. The action
	//     message need to be sent at a frequency higher than every 5 mins to keep the
	//     connection alive. Alternatively, users can chose to respond to the ping frame
	//     sent by web socket server to keep the connection alive. Refer to
	//     [connection details](https://188--nbai-docs-stg.netlify.app/docs/tracking/api/live-tracking-api#connect-to-web-socket-server)
	//     for more details.
	//
	// Any of "TRIP_SUBSCRIBE", "TRIP_UNSUBSCRIBE", "HEARTBEAT".
	Action SkynetSubscribeParamsAction `json:"action,omitzero,required"`
	// Specify a custom ID for the subscription. It can be used to reference a given
	// subscription in the downstream applications / systems.
	ID     param.Opt[string]           `json:"id,omitzero"`
	Params SkynetSubscribeParamsParams `json:"params,omitzero"`
	paramObj
}

func (r SkynetSubscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSubscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSubscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specify the behavior that needs to be achieved for the subscription. Following
// values are accepted:
//
//   - TRIP_SUBSCRIBE: Enable a trip subscription.
//   - TRIP_UNSUBSCRIBE: Unsubscribe from a trip
//   - HEARTBEAT: Enable heartbeat mechanism for a web-socket connection. The action
//     message need to be sent at a frequency higher than every 5 mins to keep the
//     connection alive. Alternatively, users can chose to respond to the ping frame
//     sent by web socket server to keep the connection alive. Refer to
//     [connection details](https://188--nbai-docs-stg.netlify.app/docs/tracking/api/live-tracking-api#connect-to-web-socket-server)
//     for more details.
type SkynetSubscribeParamsAction string

const (
	SkynetSubscribeParamsActionTripSubscribe   SkynetSubscribeParamsAction = "TRIP_SUBSCRIBE"
	SkynetSubscribeParamsActionTripUnsubscribe SkynetSubscribeParamsAction = "TRIP_UNSUBSCRIBE"
	SkynetSubscribeParamsActionHeartbeat       SkynetSubscribeParamsAction = "HEARTBEAT"
)

// The property ID is required.
type SkynetSubscribeParamsParams struct {
	// Specify the ID of an active trip that needs to be subscribed. The ID of a trip
	// is returned in the response when _Start A Trip_ request is acknowledged.
	//
	// This attribute is mandatory when action is set to either "TRIP_SUBSCRIBE" or
	// "TRIP_UNSUBSCRIBE"
	ID string `json:"id,required"`
	paramObj
}

func (r SkynetSubscribeParamsParams) MarshalJSON() (data []byte, err error) {
	type shadow SkynetSubscribeParamsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkynetSubscribeParamsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
