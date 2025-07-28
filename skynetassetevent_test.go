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

func TestSkynetAssetEventGetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Skynet.Asset.Event.GetList(
		context.TODO(),
		"id",
		nextbillionsdk.SkynetAssetEventGetListParams{
			Key:       "key=API_KEY",
			Cluster:   nextbillionsdk.SkynetAssetEventGetListParamsClusterAmerica,
			EndTime:   nextbillionsdk.Int(0),
			MonitorID: nextbillionsdk.String("monitor_id"),
			Pn:        nextbillionsdk.Int(0),
			Ps:        nextbillionsdk.Int(100),
			StartTime: nextbillionsdk.Int(0),
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
