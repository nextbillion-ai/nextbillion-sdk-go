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

func TestSkynetTripGetWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.Get(
		context.TODO(),
		"id",
		nextbillionai.SkynetTripGetParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionai.SkynetTripGetParamsClusterAmerica,
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

func TestSkynetTripUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.Update(
		context.TODO(),
		"id",
		nextbillionai.SkynetTripUpdateParams{
			Key:         "key=API_KEY",
			AssetID:     "asset_id",
			Cluster:     nextbillionai.SkynetTripUpdateParamsClusterAmerica,
			Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
			Description: nextbillionai.String("description"),
			MetaData:    `"meta_data":["Scheduled Trip", "Custom Deliveries"]`,
			Name:        nextbillionai.String(`"name": "Employee Pickup"`),
			Stops: []nextbillionai.SkynetTripUpdateParamsStop{{
				GeofenceID: nextbillionai.String("geofence_id"),
				MetaData:   `"meta_data":["Staff Entry Point", "Biometric checkpoint"]`,
				Name:       nextbillionai.String(`"name":"Head Office"`),
			}},
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

func TestSkynetTripDeleteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.Delete(
		context.TODO(),
		"id",
		nextbillionai.SkynetTripDeleteParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionai.SkynetTripDeleteParamsClusterAmerica,
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

func TestSkynetTripEndWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.End(context.TODO(), nextbillionai.SkynetTripEndParams{
		Key:     "key=API_KEY",
		ID:      "id",
		Cluster: nextbillionai.SkynetTripEndParamsClusterAmerica,
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkynetTripGetSummaryWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.GetSummary(
		context.TODO(),
		"id",
		nextbillionai.SkynetTripGetSummaryParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionai.SkynetTripGetSummaryParamsClusterAmerica,
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

func TestSkynetTripStartWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Skynet.Trip.Start(context.TODO(), nextbillionai.SkynetTripStartParams{
		Key:         "key=API_KEY",
		AssetID:     "asset_id",
		Cluster:     nextbillionai.SkynetTripStartParamsClusterAmerica,
		Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
		CustomID:    nextbillionai.String("custom_id"),
		Description: nextbillionai.String("description"),
		MetaData:    `"meta_data":["Scheduled Trip", "Custom Deliveries"]`,
		Name:        nextbillionai.String(`"name": "Employee Pickup"`),
		Stops: []nextbillionai.SkynetTripStartParamsStop{{
			GeofenceID: nextbillionai.String("geofence_id"),
			MetaData:   `"meta_data":["Staff Entry Point", "Biometric checkpoint"]`,
			Name:       nextbillionai.String(`"name":"Head Office"`),
		}},
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
