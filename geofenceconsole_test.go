// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nextbillion-ai/nextbillion-sdk-go"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/testutil"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
)

func TestGeofenceConsolePreviewWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.Console.Preview(context.TODO(), nextbillionsdk.GeofenceConsolePreviewParams{
		GeofenceEntityCreate: nextbillionsdk.GeofenceEntityCreateParam{
			Type: nextbillionsdk.GeofenceEntityCreateTypeCircle,
			Circle: nextbillionsdk.GeofenceEntityCreateCircleParam{
				Center: nextbillionsdk.GeofenceEntityCreateCircleCenterParam{
					Lat: 0,
					Lon: 0,
				},
				Radius: 0,
			},
			CustomID: nextbillionsdk.String("custom_id"),
			Isochrone: nextbillionsdk.GeofenceEntityCreateIsochroneParam{
				Coordinates:    `"coordinates": "13.25805884,77.91083661"`,
				ContoursMeter:  nextbillionsdk.Int(0),
				ContoursMinute: nextbillionsdk.Int(0),
				Denoise:        nextbillionsdk.Float(0),
				DepartureTime:  nextbillionsdk.Int(0),
				Mode:           "car",
			},
			MetaData: "{\n  \"country\": \"USA\",\n  \"state\": \"California\"\n}",
			Name:     nextbillionsdk.String(`"name":"Los Angeles Downtown"`),
			Polygon: nextbillionsdk.GeofenceEntityCreatePolygonParam{
				Geojson: nextbillionsdk.GeofenceEntityCreatePolygonGeojsonParam{
					Coordinates: [][]float64{{0}},
					Type:        "type",
				},
			},
			Tags: []string{`"tags":["tags_1", "O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"]`},
		},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeofenceConsoleSearch(t *testing.T) {
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
	_, err := client.Geofence.Console.Search(context.TODO(), nextbillionsdk.GeofenceConsoleSearchParams{
		Query: "query",
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
