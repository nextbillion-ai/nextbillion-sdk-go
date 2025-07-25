// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nextbillion-sdk-go"
	"github.com/stainless-sdks/nextbillion-sdk-go/internal/testutil"
	"github.com/stainless-sdks/nextbillion-sdk-go/option"
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
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fleetify.Routes.Steps.New(
		context.TODO(),
		"routeID",
		nextbillionsdk.FleetifyRouteStepNewParams{
			Key:                "key",
			Arrival:            0,
			Location:           []float64{0},
			Position:           0,
			Type:               nextbillionsdk.FleetifyRouteStepNewParamsTypeStart,
			Address:            nextbillionsdk.String(`"address": "503, Dublin Drive, Los Angeles, California - 500674",`),
			CompletionMode:     nextbillionsdk.RouteStepCompletionModeManual,
			DocumentTemplateID: nextbillionsdk.String("document_template_id"),
			Duration:           nextbillionsdk.Int(0),
			GeofenceConfig: nextbillionsdk.RouteStepGeofenceConfigParam{
				Radius: nextbillionsdk.Float(0),
				Type:   nextbillionsdk.RouteStepGeofenceConfigTypeCircle,
			},
			Meta: nextbillionsdk.FleetifyRouteStepNewParamsMeta{
				CustomerName:        nextbillionsdk.String(`"customer_name": "Chandler Bing"`),
				CustomerPhoneNumber: nextbillionsdk.String(`"customer_phone_number": "+1 707 234 1234"`),
				Instructions:        nextbillionsdk.String(`"instructions": "Customer asked not to ring the doorbell."`),
			},
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyRouteStepGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fleetify.Routes.Steps.Get(
		context.TODO(),
		"stepID",
		nextbillionsdk.FleetifyRouteStepGetParams{
			RouteID: "routeID",
			Key:     "key",
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
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
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Fleetify.Routes.Steps.Update(
		context.TODO(),
		"stepsID",
		nextbillionsdk.FleetifyRouteStepUpdateParams{
			RuoteID:  "ruoteID",
			Key:      "key",
			Document: map[string]interface{}{},
			Mode:     nextbillionsdk.String("mode"),
			Status:   nextbillionsdk.String("status"),
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
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
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Fleetify.Routes.Steps.Delete(
		context.TODO(),
		"stepsID",
		nextbillionsdk.FleetifyRouteStepDeleteParams{
			RuoteID: "ruoteID",
			Key:     "key",
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
