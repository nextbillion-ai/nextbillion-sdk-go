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

func TestSkynetAssetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.New(context.TODO(), nextbillionai.SkynetAssetNewParams{
		Key:         "key=API_KEY",
		Cluster:     nextbillionai.SkynetAssetNewParamsClusterAmerica,
		Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
		CustomID:    nextbillionai.String("custom_id"),
		Description: nextbillionai.String("description"),
		MetaData:    map[string]interface{}{},
		Name:        nextbillionai.String("name"),
		Tags:        []string{"string"},
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkynetAssetGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Get(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetGetParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionai.SkynetAssetGetParamsClusterAmerica,
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

func TestSkynetAssetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Update(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetUpdateParams{
			Key:         "key=API_KEY",
			Cluster:     nextbillionai.SkynetAssetUpdateParamsClusterAmerica,
			Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
			Description: nextbillionai.String("description"),
			MetaData:    map[string]interface{}{},
			Name:        nextbillionai.String("name"),
			Tags:        []string{"string"},
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

func TestSkynetAssetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.List(context.TODO(), nextbillionai.SkynetAssetListParams{
		Key:                    "key=API_KEY",
		Cluster:                nextbillionai.SkynetAssetListParamsClusterAmerica,
		IncludeAllOfAttributes: nextbillionai.String("include_all_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		IncludeAnyOfAttributes: nextbillionai.String("include_any_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		Pn:                     nextbillionai.Int(0),
		Ps:                     nextbillionai.Int(100),
		Sort:                   nextbillionai.String("updated_at:desc"),
		Tags:                   nextbillionai.String("tags=tag_1,tag_2"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkynetAssetDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Delete(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetDeleteParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionai.SkynetAssetDeleteParamsClusterAmerica,
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

func TestSkynetAssetBind(t *testing.T) {
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
	_, err := client.Skynet.Asset.Bind(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetBindParams{
			Key:      "key=API_KEY",
			DeviceID: "device_id",
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

func TestSkynetAssetTrackWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Track(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetTrackParams{
			Key:      "key=API_KEY",
			DeviceID: "device_id",
			Locations: nextbillionai.SkynetAssetTrackParamsLocations{
				Location: nextbillionai.SkynetAssetTrackParamsLocationsLocation{
					Lat: 0,
					Lon: 0,
				},
				Timestamp:    0,
				Accuracy:     nextbillionai.Float(0),
				Altitude:     nextbillionai.Float(0),
				BatteryLevel: nextbillionai.Int(0),
				Bearing:      nextbillionai.Float(0),
				MetaData:     "{\n  \"driver_name\": \"Tyler Durden\",\n  \"type\": \"parcel\"\n}",
				Speed:        nextbillionai.Float(0),
				TrackingMode: nextbillionai.String("tracking_mode"),
			},
			Cluster: nextbillionai.SkynetAssetTrackParamsClusterAmerica,
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

func TestSkynetAssetUpdateAttributes(t *testing.T) {
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
	_, err := client.Skynet.Asset.UpdateAttributes(
		context.TODO(),
		"id",
		nextbillionai.SkynetAssetUpdateAttributesParams{
			Key:        "key=API_KEY",
			Attributes: "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
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
