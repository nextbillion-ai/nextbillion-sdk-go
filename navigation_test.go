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

func TestNavigationGetRouteWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigation.GetRoute(context.TODO(), nextbillionsdk.NavigationGetRouteParams{
		Key:               "key=API_KEY",
		Altcount:          nextbillionsdk.Int(1),
		Alternatives:      nextbillionsdk.Bool(true),
		Approaches:        nextbillionsdk.NavigationGetRouteParamsApproachesUnrestricted,
		Avoid:             nextbillionsdk.NavigationGetRouteParamsAvoidToll,
		Bearings:          nextbillionsdk.String("bearings=0,180;0,180"),
		Destination:       nextbillionsdk.String("destination=41.349302,2.136480"),
		Geometry:          nextbillionsdk.NavigationGetRouteParamsGeometryPolyline,
		Lang:              nextbillionsdk.String("lang=en"),
		Mode:              nextbillionsdk.NavigationGetRouteParamsModeCar,
		Origin:            nextbillionsdk.String("origin=41.349302,2.136480"),
		OriginalShape:     nextbillionsdk.String("original_shape=sbp}_Almgp`FnLuToKmKviB{eDlcGhpFvj@qbAwoA_mA"),
		OriginalShapeType: nextbillionsdk.NavigationGetRouteParamsOriginalShapeTypePolyline,
		Overview:          nextbillionsdk.NavigationGetRouteParamsOverviewFull,
		Waypoints:         nextbillionsdk.String("waypoints=41.349302,2.136480|41.349303,2.136481|41.349304,2.136482"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
