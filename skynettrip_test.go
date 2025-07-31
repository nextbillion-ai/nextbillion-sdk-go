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

func TestSkynetTripGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.Get(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetTripGetParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetTripGetParamsClusterAmerica,
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

func TestSkynetTripUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.Update(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetTripUpdateParams{
			Key:         "key=API_KEY",
			AssetID:     "asset_id",
			Cluster:     nextbillionsdk.SkynetTripUpdateParamsClusterAmerica,
			Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
			Description: nextbillionsdk.String("description"),
			MetaData:    `"meta_data":["Scheduled Trip", "Custom Deliveries"]`,
			Name:        nextbillionsdk.String(`"name": "Employee Pickup"`),
			Stops: []nextbillionsdk.SkynetTripUpdateParamsStop{{
				GeofenceID: nextbillionsdk.String("geofence_id"),
				MetaData:   `"meta_data":["Staff Entry Point", "Biometric checkpoint"]`,
				Name:       nextbillionsdk.String(`"name":"Head Office"`),
			}},
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

func TestSkynetTripDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.Delete(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetTripDeleteParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetTripDeleteParamsClusterAmerica,
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

func TestSkynetTripEndWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.End(context.TODO(), nextbillionsdk.SkynetTripEndParams{
		Key:     "key=API_KEY",
		ID:      "id",
		Cluster: nextbillionsdk.SkynetTripEndParamsClusterAmerica,
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSkynetTripGetSummaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.GetSummary(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetTripGetSummaryParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetTripGetSummaryParamsClusterAmerica,
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

func TestSkynetTripStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Trip.Start(context.TODO(), nextbillionsdk.SkynetTripStartParams{
		Key:         "key=API_KEY",
		AssetID:     "asset_id",
		Cluster:     nextbillionsdk.SkynetTripStartParamsClusterAmerica,
		Attributes:  "{\n  \"shift_timing\": \"0800-1700\",\n  \"driver_name\": \"John\"\n}",
		CustomID:    nextbillionsdk.String("custom_id"),
		Description: nextbillionsdk.String("description"),
		MetaData:    `"meta_data":["Scheduled Trip", "Custom Deliveries"]`,
		Name:        nextbillionsdk.String(`"name": "Employee Pickup"`),
		Stops: []nextbillionsdk.SkynetTripStartParamsStop{{
			GeofenceID: nextbillionsdk.String("geofence_id"),
			MetaData:   `"meta_data":["Staff Entry Point", "Biometric checkpoint"]`,
			Name:       nextbillionsdk.String(`"name":"Head Office"`),
		}},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
