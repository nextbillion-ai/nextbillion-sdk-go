// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nextbillion-ai/nextbillion-sdk-go"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/testutil"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
)

func TestFleetifyRouteNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fleetify.Routes.New(context.TODO(), nextbillionai.FleetifyRouteNewParams{
		Key:         "key",
		DriverEmail: "johndoe@abc.com",
		Steps: []nextbillionai.RouteStepsRequestParam{{
			Arrival:            0,
			Location:           []float64{0},
			Type:               nextbillionai.RouteStepsRequestTypeStart,
			Address:            nextbillionai.String(`"address": "503, Dublin Drive, Los Angeles, California - 500674",`),
			CompletionMode:     nextbillionai.RouteStepCompletionModeManual,
			DocumentTemplateID: nextbillionai.String("document_template_id"),
			Duration:           nextbillionai.Int(0),
			GeofenceConfig: nextbillionai.RouteStepGeofenceConfigParam{
				Radius: nextbillionai.Float(0),
				Type:   nextbillionai.RouteStepGeofenceConfigTypeCircle,
			},
			Meta: nextbillionai.RouteStepsRequestMetaParam{
				CustomerName:        nextbillionai.String(`"customer_name": "Chandler Bing"`),
				CustomerPhoneNumber: nextbillionai.String(`"customer_phone_number": "+1 707 234 1234"`),
				Instructions:        nextbillionai.String(`"instructions": "Customer asked not to ring the doorbell."`),
			},
		}},
		Distance:           nextbillionai.Int(0),
		DocumentTemplateID: nextbillionai.String(`"document_template_id": "bfbc4799-bc2f-4515-9054-d888560909bf"`),
		RoRequestID:        nextbillionai.String("ro_request_id"),
		Routing: nextbillionai.FleetifyRouteNewParamsRouting{
			Approaches:    "unrestricted",
			Avoid:         "toll",
			HazmatType:    "general",
			Mode:          "car",
			TruckAxleLoad: nextbillionai.Int(0),
			TruckSize:     nextbillionai.String(`"truck_size": "200, 210, 600"`),
			TruckWeight:   nextbillionai.Int(0),
		},
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyRouteRedispatchWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fleetify.Routes.Redispatch(
		context.TODO(),
		"routeID",
		nextbillionai.FleetifyRouteRedispatchParams{
			Key: "key",
			Operations: []nextbillionai.FleetifyRouteRedispatchParamsOperation{{
				Data: nextbillionai.FleetifyRouteRedispatchParamsOperationData{
					CompletionMode:     nextbillionai.RouteStepCompletionModeManual,
					DocumentTemplateID: nextbillionai.String("document_template_id"),
					Step: nextbillionai.RouteStepsRequestParam{
						Arrival:            0,
						Location:           []float64{0},
						Type:               nextbillionai.RouteStepsRequestTypeStart,
						Address:            nextbillionai.String(`"address": "503, Dublin Drive, Los Angeles, California - 500674",`),
						CompletionMode:     nextbillionai.RouteStepCompletionModeManual,
						DocumentTemplateID: nextbillionai.String("document_template_id"),
						Duration:           nextbillionai.Int(0),
						GeofenceConfig: nextbillionai.RouteStepGeofenceConfigParam{
							Radius: nextbillionai.Float(0),
							Type:   nextbillionai.RouteStepGeofenceConfigTypeCircle,
						},
						Meta: nextbillionai.RouteStepsRequestMetaParam{
							CustomerName:        nextbillionai.String(`"customer_name": "Chandler Bing"`),
							CustomerPhoneNumber: nextbillionai.String(`"customer_phone_number": "+1 707 234 1234"`),
							Instructions:        nextbillionai.String(`"instructions": "Customer asked not to ring the doorbell."`),
						},
					},
					StepID: nextbillionai.String("step_id"),
				},
				Operation: "create",
			}},
			Distance: nextbillionai.Float(0),
		},
	)
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
