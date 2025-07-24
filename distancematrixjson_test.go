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

func TestDistancematrixJsonNew(t *testing.T) {
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
	err := client.Distancematrix.Json.New(context.TODO())
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDistancematrixJsonGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Distancematrix.Json.Get(context.TODO(), nextbillionsdk.DistancematrixJsonGetParams{
		Destinations:      "destinations=41.349302,2.136480|41.389925,2.136258|41.357961,2.097878",
		Key:               "key=API_KEY",
		Origins:           "origins:41.349302,2.136480|41.389925,2.136258|41.357961,2.097878",
		Approaches:        nextbillionsdk.DistancematrixJsonGetParamsApproachesUnrestricted,
		Avoid:             nextbillionsdk.DistancematrixJsonGetParamsAvoidToll,
		Bearings:          nextbillionsdk.String("bearings=0,180;0,180"),
		Mode:              nextbillionsdk.DistancematrixJsonGetParamsModeCar,
		RouteFailedPrompt: nextbillionsdk.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
