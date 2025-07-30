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

func TestFleetifyRouteStepNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fleetify.Routes.Steps.New(
		context.TODO(),
		"routeID",
		nextbillionai.FleetifyRouteStepNewParams{
			Key:                "key",
			Arrival:            0,
			Location:           []float64{0},
			Position:           0,
			Type:               nextbillionai.FleetifyRouteStepNewParamsTypeStart,
			Address:            nextbillionai.String(`"address": "503, Dublin Drive, Los Angeles, California - 500674",`),
			CompletionMode:     nextbillionai.RouteStepCompletionModeManual,
			DocumentTemplateID: nextbillionai.String("document_template_id"),
			Duration:           nextbillionai.Int(0),
			GeofenceConfig: nextbillionai.RouteStepGeofenceConfigParam{
				Radius: nextbillionai.Float(0),
				Type:   nextbillionai.RouteStepGeofenceConfigTypeCircle,
			},
			Meta: nextbillionai.FleetifyRouteStepNewParamsMeta{
				CustomerName:        nextbillionai.String(`"customer_name": "Chandler Bing"`),
				CustomerPhoneNumber: nextbillionai.String(`"customer_phone_number": "+1 707 234 1234"`),
				Instructions:        nextbillionai.String(`"instructions": "Customer asked not to ring the doorbell."`),
			},
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

func TestFleetifyRouteStepUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fleetify.Routes.Steps.Update(
		context.TODO(),
		"stepID",
		nextbillionai.FleetifyRouteStepUpdateParams{
			RouteID:            "routeID",
			Key:                "key",
			Arrival:            0,
			Position:           0,
			Address:            nextbillionai.String(`"address": "503, Dublin Drive, Los Angeles, California - 500674",`),
			CompletionMode:     nextbillionai.RouteStepCompletionModeManual,
			DocumentTemplateID: nextbillionai.String("document_template_id"),
			Duration:           nextbillionai.Int(0),
			GeofenceConfig: nextbillionai.RouteStepGeofenceConfigParam{
				Radius: nextbillionai.Float(0),
				Type:   nextbillionai.RouteStepGeofenceConfigTypeCircle,
			},
			Location: []float64{0},
			Meta: nextbillionai.FleetifyRouteStepUpdateParamsMeta{
				CustomerName:        nextbillionai.String(`"customer_name": "Chandler Bing"`),
				CustomerPhoneNumber: nextbillionai.String(`"customer_phone_number": "+1 707 234 1234"`),
				Instructions:        nextbillionai.String(`"instructions": "Customer asked not to ring the doorbell."`),
			},
			Type: nextbillionai.FleetifyRouteStepUpdateParamsTypeStart,
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

func TestFleetifyRouteStepDelete(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Fleetify.Routes.Steps.Delete(
		context.TODO(),
		"stepID",
		nextbillionai.FleetifyRouteStepDeleteParams{
			RouteID: "routeID",
			Key:     "key",
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

func TestFleetifyRouteStepCompleteWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	err := client.Fleetify.Routes.Steps.Complete(
		context.TODO(),
		"stepID",
		nextbillionai.FleetifyRouteStepCompleteParams{
			RouteID:  "routeID",
			Key:      "key",
			Document: map[string]interface{}{},
			Mode:     nextbillionai.String("mode"),
			Status:   nextbillionai.String("status"),
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
