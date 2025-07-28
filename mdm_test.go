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

func TestMdmNewDistanceMatrixWithOptionalParams(t *testing.T) {
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
	_, err := client.Mdm.NewDistanceMatrix(context.TODO(), nextbillionsdk.MdmNewDistanceMatrixParams{
		Key:                  "key=API_KEY",
		Option:               nextbillionsdk.MdmNewDistanceMatrixParamsOptionFlexible,
		Origins:              "origins",
		Spliter:              nextbillionsdk.MdmNewDistanceMatrixParamsSpliterOdNumberSpliter,
		Area:                 nextbillionsdk.MdmNewDistanceMatrixParamsAreaSingapore,
		Avoid:                nextbillionsdk.MdmNewDistanceMatrixParamsAvoidToll,
		CrossBorder:          nextbillionsdk.Bool(true),
		DepartureTime:        nextbillionsdk.Int(0),
		Destinations:         nextbillionsdk.String("destinations"),
		DestinationsApproach: nextbillionsdk.MdmNewDistanceMatrixParamsDestinationsApproachUnrestricted,
		HazmatType:           nextbillionsdk.MdmNewDistanceMatrixParamsHazmatTypeGeneral,
		Mode:                 nextbillionsdk.MdmNewDistanceMatrixParamsModeCar,
		OriginsApproach:      nextbillionsdk.MdmNewDistanceMatrixParamsOriginsApproachUnrestricted,
		RouteType:            nextbillionsdk.MdmNewDistanceMatrixParamsRouteTypeFastest,
		TruckAxleLoad:        nextbillionsdk.Float(0),
		TruckSize:            nextbillionsdk.String(`"truck_size"=200,210,600`),
		TruckWeight:          nextbillionsdk.Int(0),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMdmGetDistanceMatrixStatus(t *testing.T) {
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
	_, err := client.Mdm.GetDistanceMatrixStatus(context.TODO(), nextbillionsdk.MdmGetDistanceMatrixStatusParams{
		ID:  "id",
		Key: "key=API_KEY",
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
