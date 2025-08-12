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

func TestNavigationGetRouteWithOptionalParams(t *testing.T) {
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
	_, err := client.Navigation.GetRoute(context.TODO(), nextbillionai.NavigationGetRouteParams{
		Key:               "key=API_KEY",
		Altcount:          nextbillionai.Int(1),
		Alternatives:      nextbillionai.Bool(true),
		Approaches:        nextbillionai.NavigationGetRouteParamsApproachesUnrestricted,
		Avoid:             nextbillionai.NavigationGetRouteParamsAvoidToll,
		Bearings:          nextbillionai.String("bearings=0,180;0,180"),
		Destination:       nextbillionai.String("destination=41.349302,2.136480"),
		Geometry:          nextbillionai.NavigationGetRouteParamsGeometryPolyline,
		Lang:              nextbillionai.String("lang=en"),
		Mode:              nextbillionai.NavigationGetRouteParamsModeCar,
		Origin:            nextbillionai.String("origin=41.349302,2.136480"),
		OriginalShape:     nextbillionai.String("original_shape=sbp}_AlmgpFnLuToKmKviB{eDlcGhpFvj@qbAwoA_mA"),
		OriginalShapeType: nextbillionai.NavigationGetRouteParamsOriginalShapeTypePolyline,
		Overview:          nextbillionai.NavigationGetRouteParamsOverviewFull,
		Waypoints:         nextbillionai.String("waypoints=41.349302,2.136480|41.349303,2.136481|41.349304,2.136482"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
