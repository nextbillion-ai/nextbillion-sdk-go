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

func TestDistanceMatrixJsonNew(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	err := client.DistanceMatrix.Json.New(context.TODO())
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDistanceMatrixJsonGetWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.DistanceMatrix.Json.Get(context.TODO(), nextbillionai.DistanceMatrixJsonGetParams{
		Destinations:      "destinations=41.349302,2.136480|41.389925,2.136258|41.357961,2.097878",
		Key:               "key=API_KEY",
		Origins:           "origins:41.349302,2.136480|41.389925,2.136258|41.357961,2.097878",
		Approaches:        nextbillionai.DistanceMatrixJsonGetParamsApproachesUnrestricted,
		Avoid:             nextbillionai.DistanceMatrixJsonGetParamsAvoidToll,
		Bearings:          nextbillionai.String("bearings=0,180;0,180"),
		Mode:              nextbillionai.DistanceMatrixJsonGetParamsModeCar,
		RouteFailedPrompt: nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
