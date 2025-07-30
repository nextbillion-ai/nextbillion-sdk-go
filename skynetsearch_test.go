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

func TestSkynetSearchAroundWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Search.Around(context.TODO(), nextbillionai.SkynetSearchAroundParams{
		Center:                 "56.597801,43.967836",
		Key:                    "key=API_KEY",
		Radius:                 0,
		Filter:                 nextbillionai.String("filter=tag:delivery,truck"),
		IncludeAllOfAttributes: nextbillionai.String("include_all_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		IncludeAnyOfAttributes: nextbillionai.String("include_any_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		MaxSearchLimit:         nextbillionai.Bool(true),
		Pn:                     nextbillionai.Int(0),
		Ps:                     nextbillionai.Int(100),
		SortBy:                 nextbillionai.SkynetSearchAroundParamsSortByDistance,
		SortDestination:        nextbillionai.String("sort_destination= 34.0241,-118.2550"),
		SortDrivingMode:        nextbillionai.SkynetSearchAroundParamsSortDrivingModeCar,
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkynetSearchBoundWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Search.Bound(context.TODO(), nextbillionai.SkynetSearchBoundParams{
		Bound:                  "bounds=44.7664,-0.6941|44.9206,-0.4639",
		Key:                    "key=API_KEY",
		Filter:                 nextbillionai.String("filter=tag:delivery,truck"),
		IncludeAllOfAttributes: nextbillionai.String("include_all_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		IncludeAnyOfAttributes: nextbillionai.String("include_any_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		MaxSearchLimit:         nextbillionai.Bool(true),
		Pn:                     nextbillionai.Int(0),
		Ps:                     nextbillionai.Int(100),
		SortBy:                 nextbillionai.SkynetSearchBoundParamsSortByDistance,
		SortDestination:        nextbillionai.String("sort_destination= 34.0241,-118.2550"),
		SortDrivingMode:        nextbillionai.SkynetSearchBoundParamsSortDrivingModeCar,
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
