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

func TestGeocodeGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.Get(context.TODO(), nextbillionsdk.GeocodeGetParams{
		Key:   "key=API_KEY",
		Q:     "q=125, Berliner, berlin",
		At:    nextbillionsdk.String("at=52.5308,13.3856"),
		In:    nextbillionsdk.String("in=countryCode:CAN,MEX,USA"),
		Lang:  nextbillionsdk.String("lang=en"),
		Limit: nextbillionsdk.Int(0),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeocodeBatchNew(t *testing.T) {
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
	_, err := client.Geocode.BatchNew(context.TODO(), nextbillionsdk.GeocodeBatchNewParams{
		Key: "key=API_KEY",
		Body: []nextbillionsdk.GeocodeBatchNewParamsBody{{
			Q:     `"q":"125, Berliner, berlin"`,
			At:    nextbillionsdk.String(`"at": "52.5308,13.3856"`),
			In:    nextbillionsdk.String(`"in":"countryCode:CAN,MEX,USA"`),
			Lang:  nextbillionsdk.String(`"lang":"en"`),
			Limit: nextbillionsdk.Int(0),
		}},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeocodeStructuredGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.StructuredGet(context.TODO(), nextbillionsdk.GeocodeStructuredGetParams{
		CountryCode: "countryCode",
		Key:         "key=API_KEY",
		At:          nextbillionsdk.String("at=52.5308,13.3856"),
		City:        nextbillionsdk.String("city"),
		County:      nextbillionsdk.String("county"),
		HouseNumber: nextbillionsdk.String("houseNumber"),
		In:          nextbillionsdk.String("in=circle:52.53,13.38;r=10000"),
		Limit:       nextbillionsdk.Int(0),
		PostalCode:  nextbillionsdk.String("postalCode"),
		State:       nextbillionsdk.String("state"),
		Street:      nextbillionsdk.String("street"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
