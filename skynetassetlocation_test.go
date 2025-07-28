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

func TestSkynetAssetLocationListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Location.List(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetLocationListParams{
			Key:          "key=API_KEY",
			Cluster:      nextbillionsdk.SkynetAssetLocationListParamsClusterAmerica,
			Correction:   nextbillionsdk.String("correction=mapmatch=1,interpolate=0,mode=`car`"),
			EndTime:      nextbillionsdk.Int(0),
			GeometryType: nextbillionsdk.SkynetAssetLocationListParamsGeometryTypePolyline,
			Pn:           nextbillionsdk.Int(0),
			Ps:           nextbillionsdk.Int(500),
			StartTime:    nextbillionsdk.Int(0),
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

func TestSkynetAssetLocationGetLastWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Location.GetLast(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetLocationGetLastParams{
			Key:     "key=API_KEY",
			Cluster: nextbillionsdk.SkynetAssetLocationGetLastParamsClusterAmerica,
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
