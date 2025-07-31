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

func TestGeofenceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.New(context.TODO(), nextbillionai.GeofenceNewParams{
		Key: "key=API_KEY",
		GeofenceEntityCreate: nextbillionai.GeofenceEntityCreateParam{
			Type: nextbillionai.GeofenceEntityCreateTypeCircle,
			Circle: nextbillionai.GeofenceEntityCreateCircleParam{
				Center: nextbillionai.GeofenceEntityCreateCircleCenterParam{
					Lat: 0,
					Lon: 0,
				},
				Radius: 0,
			},
			CustomID: nextbillionai.String("custom_id"),
			Isochrone: nextbillionai.GeofenceEntityCreateIsochroneParam{
				Coordinates:    `"coordinates": "13.25805884,77.91083661"`,
				ContoursMeter:  nextbillionai.Int(0),
				ContoursMinute: nextbillionai.Int(0),
				Denoise:        nextbillionai.Float(0),
				DepartureTime:  nextbillionai.Int(0),
				Mode:           "car",
			},
			MetaData: "{\n  \"country\": \"USA\",\n  \"state\": \"California\"\n}",
			Name:     nextbillionai.String(`"name":"Los Angeles Downtown"`),
			Polygon: nextbillionai.GeofenceEntityCreatePolygonParam{
				Geojson: nextbillionai.GeofenceEntityCreatePolygonGeojsonParam{
					Coordinates: [][]float64{{0}},
					Type:        "type",
				},
			},
			Tags: []string{`"tags":["tags_1", "O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"]`},
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

func TestGeofenceGet(t *testing.T) {
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
	_, err := client.Geofence.Get(
		context.TODO(),
		"id",
		nextbillionai.GeofenceGetParams{
			Key: "key=API_KEY",
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

func TestGeofenceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.Update(
		context.TODO(),
		"id",
		nextbillionai.GeofenceUpdateParams{
			Key: "key=API_KEY",
			Circle: nextbillionai.GeofenceUpdateParamsCircle{
				Center: nextbillionai.GeofenceUpdateParamsCircleCenter{
					Lat: nextbillionai.Float(0),
					Lon: nextbillionai.Float(0),
				},
				Radius: nextbillionai.Float(0),
			},
			Isochrone: nextbillionai.GeofenceUpdateParamsIsochrone{
				ContoursMeter:  nextbillionai.Int(0),
				ContoursMinute: nextbillionai.Int(0),
				Coordinates:    nextbillionai.String(`"coordinates": "13.25805884388484,77.91083661048299"`),
				Denoise:        nextbillionai.Float(0),
				DepartureTime:  nextbillionai.Int(0),
				Mode:           nextbillionai.String("“mode”:”car”"),
			},
			MetaData: "",
			Name:     nextbillionai.String(`"name":"Los Angeles Downtown"`),
			Polygon: nextbillionai.GeofenceUpdateParamsPolygon{
				Geojson: nextbillionai.GeofenceUpdateParamsPolygonGeojson{
					Geometry: [][]float64{{0}},
					Type:     nextbillionai.String("type"),
				},
			},
			Tags: []string{`"tags":["tags_1", "O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"]`},
			Type: nextbillionai.GeofenceUpdateParamsTypeCircle,
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

func TestGeofenceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.List(context.TODO(), nextbillionai.GeofenceListParams{
		Key:  "key=API_KEY",
		Pn:   nextbillionai.Int(0),
		Ps:   nextbillionai.Int(100),
		Tags: nextbillionai.String("tags=tags_1,O69Am2Y4KL8q5Y5JuD-Fy-tdtEpkTRQo_ZYIK7"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
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
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Geofence.Delete(
		context.TODO(),
		"id",
		nextbillionai.GeofenceDeleteParams{
			Key: "key=API_KEY",
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

func TestGeofenceContainsWithOptionalParams(t *testing.T) {
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
	_, err := client.Geofence.Contains(context.TODO(), nextbillionai.GeofenceContainsParams{
		Key:       "key=API_KEY",
		Locations: "13.25805884388484,77.91083661048299|13.25805884388484,77.91083661048299",
		Geofences: nextbillionai.String("geofences=80d1fa55-6287-4da0-93ac-2fc162d0a228,70d1fa55-1287-4da0-93ac-2fc162d0a228"),
		Verbose:   nextbillionai.String("verbose=true"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
