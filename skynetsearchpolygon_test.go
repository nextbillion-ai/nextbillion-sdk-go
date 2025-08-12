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

func TestSkynetSearchPolygonNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Search.Polygon.New(context.TODO(), nextbillionai.SkynetSearchPolygonNewParams{
		Key: "key=API_KEY",
		Polygon: nextbillionai.SkynetSearchPolygonNewParamsPolygon{
			Coordinates: []float64{0},
			Type:        "type",
		},
		Filter: nextbillionai.String(`"tag:delivery,truck"`),
		MatchFilter: nextbillionai.SkynetSearchPolygonNewParamsMatchFilter{
			IncludeAllOfAttributes: nextbillionai.String(`"shift_timing": "0800-1700","driver_name": "John"`),
			IncludeAnyOfAttributes: nextbillionai.String("include_any_of_attributes"),
		},
		MaxSearchLimit: nextbillionai.Bool(true),
		Pn:             nextbillionai.Int(0),
		Ps:             nextbillionai.Int(0),
		Sort: nextbillionai.SkynetSearchPolygonNewParamsSort{
			SortBy: "distance",
			SortDestination: nextbillionai.SkynetSearchPolygonNewParamsSortSortDestination{
				Lat: 0,
				Lon: 0,
			},
			SortDrivingMode: "car",
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

func TestSkynetSearchPolygonGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Search.Polygon.Get(context.TODO(), nextbillionai.SkynetSearchPolygonGetParams{
		Key:                    "key=API_KEY",
		Polygon:                "polygon=17.4239,78.4590|17.4575,78.4624|17.4547,78.5483|17.4076,78.5527|17.4239,78.4590",
		Filter:                 nextbillionai.String("filter=tag:delivery,truck"),
		IncludeAllOfAttributes: nextbillionai.String("include_all_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		IncludeAnyOfAttributes: nextbillionai.String("include_any_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		MaxSearchLimit:         nextbillionai.Bool(true),
		Pn:                     nextbillionai.Int(0),
		Ps:                     nextbillionai.Int(100),
		SortBy:                 nextbillionai.SkynetSearchPolygonGetParamsSortByDistance,
		SortDestination:        nextbillionai.String("sort_destination= 34.0241,-118.2550"),
		SortDrivingMode:        nextbillionai.SkynetSearchPolygonGetParamsSortDrivingModeCar,
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
