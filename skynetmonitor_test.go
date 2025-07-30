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

func TestSkynetMonitorNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.New(context.TODO(), nextbillionai.SkynetMonitorNewParams{
		Key:         "key=API_KEY",
		Tags:        []string{"string"},
		Type:        nextbillionai.SkynetMonitorNewParamsTypeEnter,
		Cluster:     nextbillionai.SkynetMonitorNewParamsClusterAmerica,
		CustomID:    nextbillionai.String("custom_id"),
		Description: nextbillionai.String("description"),
		GeofenceConfig: nextbillionai.SkynetMonitorNewParamsGeofenceConfig{
			GeofenceIDs: []string{"string"},
		},
		GeofenceIDs: []string{"string"},
		IdleConfig: nextbillionai.SkynetMonitorNewParamsIdleConfig{
			DistanceTolerance: 0,
			TimeTolerance:     nextbillionai.Int(0),
		},
		MatchFilter: nextbillionai.SkynetMonitorNewParamsMatchFilter{
			IncludeAllOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
			IncludeAnyOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
		},
		MetaData: map[string]interface{}{},
		Name:     nextbillionai.String("name"),
		SpeedingConfig: nextbillionai.SkynetMonitorNewParamsSpeedingConfig{
			CustomerSpeedLimit: nextbillionai.Int(0),
			TimeTolerance:      nextbillionai.Int(0),
			UseAdminSpeedLimit: nextbillionai.Bool(true),
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

func TestSkynetMonitorGet(t *testing.T) {
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
	_, err := client.Skynet.Monitor.Get(
		context.TODO(),
		"id",
		nextbillionai.SkynetMonitorGetParams{
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

func TestSkynetMonitorUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.Update(
		context.TODO(),
		"id",
		nextbillionai.SkynetMonitorUpdateParams{
			Key:         "key=API_KEY",
			Description: nextbillionai.String("description"),
			GeofenceConfig: nextbillionai.SkynetMonitorUpdateParamsGeofenceConfig{
				GeofenceIDs: []string{"string"},
			},
			GeofenceIDs: []string{"string"},
			IdleConfig: nextbillionai.SkynetMonitorUpdateParamsIdleConfig{
				DistanceTolerance: 0,
				TimeTolerance:     nextbillionai.Int(0),
			},
			MatchFilter: nextbillionai.SkynetMonitorUpdateParamsMatchFilter{
				IncludeAllOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
				IncludeAnyOfAttributes: "{\n  \"asset_type\": \"delivery\",\n  \"area\": \"Los Angeles downtown\"\n}",
			},
			MetaData: map[string]interface{}{},
			Name:     nextbillionai.String(`"name":"warehouse_exit"`),
			SpeedingConfig: nextbillionai.SkynetMonitorUpdateParamsSpeedingConfig{
				CustomerSpeedLimit: nextbillionai.String(`"customer_speed_limit":8`),
				TimeTolerance:      nextbillionai.Int(0),
				UseAdminSpeedLimit: nextbillionai.Bool(true),
			},
			Tags: []string{"string"},
			Type: nextbillionai.SkynetMonitorUpdateParamsTypeEnter,
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

func TestSkynetMonitorListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Monitor.List(context.TODO(), nextbillionai.SkynetMonitorListParams{
		Key:     "key=API_KEY",
		Cluster: nextbillionai.SkynetMonitorListParamsClusterAmerica,
		Pn:      nextbillionai.Int(0),
		Ps:      nextbillionai.Int(100),
		Sort:    nextbillionai.String("updated_at:desc"),
		Tags:    nextbillionai.String("tags=tag_1,tag_2"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
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
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Skynet.Monitor.Delete(
		context.TODO(),
		"id",
		nextbillionai.SkynetMonitorDeleteParams{
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
