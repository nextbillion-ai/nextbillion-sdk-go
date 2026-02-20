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

func TestMdmNewDistanceMatrixWithOptionalParams(t *testing.T) {
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
	_, err := client.Mdm.NewDistanceMatrix(context.TODO(), nextbillionai.MdmNewDistanceMatrixParams{
		Key:                  "key=API_KEY",
		Option:               nextbillionai.MdmNewDistanceMatrixParamsOptionFlexible,
		Origins:              "origins",
		Spliter:              nextbillionai.MdmNewDistanceMatrixParamsSpliterOdNumberSpliter,
		Area:                 nextbillionai.MdmNewDistanceMatrixParamsAreaSingapore,
		Avoid:                nextbillionai.MdmNewDistanceMatrixParamsAvoidToll,
		CrossBorder:          nextbillionai.Bool(true),
		DepartureTime:        nextbillionai.Int(0),
		Destinations:         nextbillionai.String("destinations"),
		DestinationsApproach: nextbillionai.MdmNewDistanceMatrixParamsDestinationsApproachUnrestricted,
		HazmatType:           nextbillionai.MdmNewDistanceMatrixParamsHazmatTypeGeneral,
		Mode:                 nextbillionai.MdmNewDistanceMatrixParamsModeCar,
		OriginsApproach:      nextbillionai.MdmNewDistanceMatrixParamsOriginsApproachUnrestricted,
		RouteType:            nextbillionai.MdmNewDistanceMatrixParamsRouteTypeFastest,
		TruckAxleLoad:        nextbillionai.Float(0),
		TruckSize:            nextbillionai.String(`"truck_size"=200,210,600`),
		TruckWeight:          nextbillionai.Int(0),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMdmGetDistanceMatrixStatus(t *testing.T) {
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
	_, err := client.Mdm.GetDistanceMatrixStatus(context.TODO(), nextbillionai.MdmGetDistanceMatrixStatusParams{
		ID:  "id",
		Key: "key=API_KEY",
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
