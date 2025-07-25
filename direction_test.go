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

func TestDirectionComputeRouteWithOptionalParams(t *testing.T) {
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
	_, err := client.Directions.ComputeRoute(context.TODO(), nextbillionsdk.DirectionComputeRouteParams{
		QueryKey:        "key",
		Destination:     "41.349302,2.136480",
		BodyKey:         "key",
		Origin:          "41.349302,2.136480",
		Altcount:        nextbillionsdk.Int(1),
		Alternatives:    nextbillionsdk.Bool(true),
		Approaches:      nextbillionsdk.String("unrestricted;;curb;"),
		Avoid:           nextbillionsdk.DirectionComputeRouteParamsAvoidToll,
		Bearings:        nextbillionsdk.String("0,180;0,180"),
		CrossBorder:     nextbillionsdk.Bool(true),
		DepartureTime:   nextbillionsdk.Int(0),
		DriveTimeLimits: nextbillionsdk.String("500,400,400"),
		EmissionClass:   nextbillionsdk.DirectionComputeRouteParamsEmissionClassEuro0,
		Exclude:         nextbillionsdk.DirectionComputeRouteParamsExcludeToll,
		Geometry:        nextbillionsdk.DirectionComputeRouteParamsGeometryPolyline,
		HazmatType:      nextbillionsdk.DirectionComputeRouteParamsHazmatTypeGeneral,
		Mode:            nextbillionsdk.DirectionComputeRouteParamsModeCar,
		Option:          nextbillionsdk.DirectionComputeRouteParamsOptionFast,
		Overview:        nextbillionsdk.DirectionComputeRouteParamsOverviewFull,
		RestTimes:       nextbillionsdk.String("500,300,100"),
		RoadInfo:        nextbillionsdk.DirectionComputeRouteParamsRoadInfoMaxSpeed,
		RouteType:       nextbillionsdk.DirectionComputeRouteParamsRouteTypeFastest,
		Steps:           nextbillionsdk.Bool(true),
		TruckAxleLoad:   nextbillionsdk.Float(0),
		TruckSize:       nextbillionsdk.String("200,210,600"),
		TruckWeight:     nextbillionsdk.Int(1),
		TurnAngleRange:  nextbillionsdk.Int(0),
		Waypoints:       nextbillionsdk.String("41.349302,2.136480|41.349303,2.136481|41.349304,2.136482"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
