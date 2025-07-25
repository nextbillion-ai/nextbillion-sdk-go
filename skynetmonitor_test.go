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

func TestSkynetMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.New(context.TODO(), nextbillionsdk.SkynetMonitorNewParams{
		Key:         "key=API_KEY",
		Tags:        []string{"string"},
		Type:        nextbillionsdk.SkynetMonitorNewParamsTypeEnter,
		Cluster:     nextbillionsdk.SkynetMonitorNewParamsClusterAmerica,
		CustomID:    nextbillionsdk.String("custom_id"),
		Description: nextbillionsdk.String("description"),
		GeofenceConfig: nextbillionsdk.SkynetMonitorNewParamsGeofenceConfig{
			GeofenceIDs: []string{"string"},
		},
		GeofenceIDs: []string{"string"},
		IdleConfig: nextbillionsdk.SkynetMonitorNewParamsIdleConfig{
			DistanceTolerance: 0,
			TimeTolerance:     nextbillionsdk.Int(0),
		},
		MatchFilter: nextbillionsdk.SkynetMonitorNewParamsMatchFilter{
			IncludeAllOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
			IncludeAnyOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
		},
		MetaData: map[string]interface{}{},
		Name:     nextbillionsdk.String("name"),
		SpeedingConfig: nextbillionsdk.SkynetMonitorNewParamsSpeedingConfig{
			CustomerSpeedLimit: nextbillionsdk.Int(0),
			TimeTolerance:      nextbillionsdk.Int(0),
			UseAdminSpeedLimit: nextbillionsdk.Bool(true),
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

func TestSkynetMonitorGet(t *testing.T) {
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
	_, err := client.Skynet.Monitor.Get(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetMonitorGetParams{
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

func TestSkynetMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.Update(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetMonitorUpdateParams{
			Key:         "key=API_KEY",
			Description: nextbillionsdk.String("description"),
			GeofenceConfig: nextbillionsdk.SkynetMonitorUpdateParamsGeofenceConfig{
				GeofenceIDs: []string{"string"},
			},
			GeofenceIDs: []string{"string"},
			IdleConfig: nextbillionsdk.SkynetMonitorUpdateParamsIdleConfig{
				DistanceTolerance: 0,
				TimeTolerance:     nextbillionsdk.Int(0),
			},
			MatchFilter: nextbillionsdk.SkynetMonitorUpdateParamsMatchFilter{
				IncludeAllOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
				IncludeAnyOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
			},
			MetaData: map[string]interface{}{},
			Name:     nextbillionsdk.String(`"name":"warehouse_exit"`),
			SpeedingConfig: nextbillionsdk.SkynetMonitorUpdateParamsSpeedingConfig{
				CustomerSpeedLimit: nextbillionsdk.String(`"customer_speed_limit":8`),
				TimeTolerance:      nextbillionsdk.Int(0),
				UseAdminSpeedLimit: nextbillionsdk.Bool(true),
			},
			Tags: []string{"string"},
			Type: nextbillionsdk.SkynetMonitorUpdateParamsTypeEnter,
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

func TestSkynetMonitorDelete(t *testing.T) {
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
	_, err := client.Skynet.Monitor.Delete(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetMonitorDeleteParams{
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

func TestSkynetMonitorGetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.GetList(context.TODO(), nextbillionsdk.SkynetMonitorGetListParams{
		Key:     "key=API_KEY",
		Cluster: nextbillionsdk.SkynetMonitorGetListParamsClusterAmerica,
		Pn:      nextbillionsdk.Int(0),
		Ps:      nextbillionsdk.Int(100),
		Sort:    nextbillionsdk.String("updated_at:desc"),
		Tags:    nextbillionsdk.String("tags=tag_1,tag_2"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
