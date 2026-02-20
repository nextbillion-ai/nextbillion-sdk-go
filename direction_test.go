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

func TestDirectionComputeRouteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Directions.ComputeRoute(context.TODO(), nextbillionai.DirectionComputeRouteParams{
		Destination:     "41.349302,2.136480",
		Origin:          "41.349302,2.136480",
		Altcount:        nextbillionai.Int(1),
		Alternatives:    nextbillionai.Bool(true),
		Approaches:      nextbillionai.String("unrestricted;;curb;"),
		Avoid:           nextbillionai.DirectionComputeRouteParamsAvoidToll,
		Bearings:        nextbillionai.String("0,180;0,180"),
		CrossBorder:     nextbillionai.Bool(true),
		DepartureTime:   nextbillionai.Int(0),
		DriveTimeLimits: nextbillionai.String("500,400,400"),
		EmissionClass:   nextbillionai.DirectionComputeRouteParamsEmissionClassEuro0,
		Exclude:         nextbillionai.DirectionComputeRouteParamsExcludeToll,
		Geometry:        nextbillionai.DirectionComputeRouteParamsGeometryPolyline,
		HazmatType:      nextbillionai.DirectionComputeRouteParamsHazmatTypeGeneral,
		Mode:            nextbillionai.DirectionComputeRouteParamsModeCar,
		Option:          nextbillionai.DirectionComputeRouteParamsOptionFast,
		Overview:        nextbillionai.DirectionComputeRouteParamsOverviewFull,
		RestTimes:       nextbillionai.String("500,300,100"),
		RoadInfo:        nextbillionai.DirectionComputeRouteParamsRoadInfoMaxSpeed,
		RouteType:       nextbillionai.DirectionComputeRouteParamsRouteTypeFastest,
		Steps:           nextbillionai.Bool(true),
		TruckAxleLoad:   nextbillionai.Float(0),
		TruckSize:       nextbillionai.String("200,210,600"),
		TruckWeight:     nextbillionai.Int(1),
		TurnAngleRange:  nextbillionai.Int(0),
		Waypoints:       nextbillionai.String("41.349302,2.136480|41.349303,2.136481|41.349304,2.136482"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
