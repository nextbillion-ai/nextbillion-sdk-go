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

func TestSkynetAssetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.New(context.TODO(), nextbillionsdk.SkynetAssetNewParams{
		Key:         "key=API_KEY",
		Cluster:     nextbillionsdk.SkynetAssetNewParamsClusterAmerica,
		Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
		CustomID:    nextbillionsdk.String("custom_id"),
		Description: nextbillionsdk.String("description"),
		MetaData:    map[string]interface{}{},
		Name:        nextbillionsdk.String("name"),
		Tags:        []string{"string"},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
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
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Skynet.Asset.Get(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetGetParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetAssetGetParamsClusterAmerica,
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

func TestSkynetAssetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Update(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetUpdateParams{
			Key:         "key=API_KEY",
			Cluster:     nextbillionsdk.SkynetAssetUpdateParamsClusterAmerica,
			Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
			Description: nextbillionsdk.String("description"),
			MetaData:    map[string]interface{}{},
			Name:        nextbillionsdk.String("name"),
			Tags:        []string{"string"},
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

func TestSkynetAssetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.List(context.TODO(), nextbillionsdk.SkynetAssetListParams{
		Key:                    "key=API_KEY",
		Cluster:                nextbillionsdk.SkynetAssetListParamsClusterAmerica,
		IncludeAllOfAttributes: nextbillionsdk.String("include_all_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		IncludeAnyOfAttributes: nextbillionsdk.String("include_any_of_attributes=vehicle_type:pickup_truck|driver_name:John"),
		Pn:                     nextbillionsdk.Int(0),
		Ps:                     nextbillionsdk.Int(100),
		Sort:                   nextbillionsdk.String("updated_at:desc"),
		Tags:                   nextbillionsdk.String("tags=tag_1,tag_2"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
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
	client := nextbillionsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Skynet.Asset.Delete(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetDeleteParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetAssetDeleteParamsClusterAmerica,
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

func TestSkynetAssetBind(t *testing.T) {
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
	_, err := client.Skynet.Asset.Bind(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetBindParams{
			Key:      "key=API_KEY",
			DeviceID: "device_id",
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

func TestSkynetAssetTrackWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Track(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetTrackParams{
			Key:      "key=API_KEY",
			DeviceID: "device_id",
			Locations: nextbillionsdk.SkynetAssetTrackParamsLocations{
				Location: nextbillionsdk.SkynetAssetTrackParamsLocationsLocation{
					Lat: 0,
					Lon: 0,
				},
				Timestamp:    0,
				Accuracy:     nextbillionsdk.Float(0),
				Altitude:     nextbillionsdk.Float(0),
				BatteryLevel: nextbillionsdk.Int(0),
				Bearing:      nextbillionsdk.Float(0),
				MetaData:     "{\n  \"driver_name\": \"Tyler Durden\",\n  \"type\": \"parcel\"\n}",
				Speed:        nextbillionsdk.Float(0),
				TrackingMode: nextbillionsdk.String("tracking_mode"),
			},
			Cluster: nextbillionsdk.SkynetAssetTrackParamsClusterAmerica,
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

func TestSkynetAssetUpdateAttributes(t *testing.T) {
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
	_, err := client.Skynet.Asset.UpdateAttributes(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetUpdateAttributesParams{
			Key:        "key=API_KEY",
			Attributes: "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
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
