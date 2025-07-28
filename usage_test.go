// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/nextbillion-ai/nextbillion-sdk-go"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/testutil"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	route, err := client.Fleetify.Routes.New(context.TODO(), nextbillionsdk.FleetifyRouteNewParams{
		Key:         "REPLACE_ME",
		DriverEmail: "REPLACE_ME",
		Steps: []nextbillionsdk.RouteStepsRequestParam{{
			Arrival:  0,
			Location: []float64{0},
			Type:     nextbillionsdk.RouteStepsRequestTypeStart,
		}},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", route.Data)
}
