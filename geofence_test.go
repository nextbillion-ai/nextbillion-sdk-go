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

func TestGeofenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.New(context.TODO(), nextbillionsdk.GeofenceNewParams{
		Key: "key=API_KEY",
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
				Mode:           "`car`",
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

func TestGeofenceGet(t *testing.T) {
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
	_, err := client.Geofence.Get(
		context.TODO(),
		"id",
		nextbillionsdk.GeofenceGetParams{
			Key: "key=API_KEY",
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

func TestGeofenceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.Update(
		context.TODO(),
		"id",
		nextbillionsdk.GeofenceUpdateParams{
			Key: "key=API_KEY",
			Circle: nextbillionsdk.GeofenceUpdateParamsCircle{
				Center: nextbillionsdk.GeofenceUpdateParamsCircleCenter{
					Lat: nextbillionsdk.Float(0),
					Lon: nextbillionsdk.Float(0),
				},
				Radius: nextbillionsdk.Float(0),
			},
			Isochrone: nextbillionsdk.GeofenceUpdateParamsIsochrone{
				ContoursMeter:  nextbillionsdk.Int(0),
				ContoursMinute: nextbillionsdk.Int(0),
				Coordinates:    nextbillionsdk.String(`"coordinates": "13.25805884388484,77.91083661048299"`),
				Denoise:        nextbillionsdk.Float(0),
				DepartureTime:  nextbillionsdk.Int(0),
				Mode:           nextbillionsdk.String("“mode”:”car”"),
			},
			MetaData: "",
			Name:     nextbillionsdk.String(`"name":"Los Angeles Downtown"`),
			Polygon: nextbillionsdk.GeofenceUpdateParamsPolygon{
				Geojson: nextbillionsdk.GeofenceUpdateParamsPolygonGeojson{
					Geometry: [][]float64{{0}},
					Type:     nextbillionsdk.String("type"),
				},
			},
			Tags: []string{`"tags":["tags_1", "O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"]`},
			Type: nextbillionsdk.GeofenceUpdateParamsTypeCircle,
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

func TestGeofenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.List(context.TODO(), nextbillionsdk.GeofenceListParams{
		Key:  "key=API_KEY",
		Pn:   nextbillionsdk.Int(0),
		Ps:   nextbillionsdk.Int(100),
		Tags: nextbillionsdk.String("tags=tags_1,O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeofenceDelete(t *testing.T) {
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
	_, err := client.Geofence.Delete(
		context.TODO(),
		"id",
		nextbillionsdk.GeofenceDeleteParams{
			Key: "key=API_KEY",
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

func TestGeofenceContainsWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.Contains(context.TODO(), nextbillionsdk.GeofenceContainsParams{
		Key:       "key=API_KEY",
		Locations: "13.25805884388484,77.91083661048299|13.25805884388484,77.91083661048299",
		Geofences: nextbillionsdk.String("geofences=80d1fa55-6287-4da0-93ac-2fc162d0a228,70d1fa55-1287-4da0-93ac-2fc162d0a228"),
		Verbose:   nextbillionsdk.String("verbose=true"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
