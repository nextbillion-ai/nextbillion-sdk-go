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

func TestGeocodeGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Geocode.Get(context.TODO(), nextbillionai.GeocodeGetParams{
		Key:   "key=API_KEY",
		Q:     "q=125, Berliner, berlin",
		At:    nextbillionai.String("at=52.5308,13.3856"),
		In:    nextbillionai.String("in=countryCode:CAN,MEX,USA"),
		Lang:  nextbillionai.String("lang=en"),
		Limit: nextbillionai.Int(0),
	})
	if err != nil {
		var apierr *nextbillionai.Error
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
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Geocode.BatchNew(context.TODO(), nextbillionai.GeocodeBatchNewParams{
		Key: "key=API_KEY",
		Body: []nextbillionai.GeocodeBatchNewParamsBody{{
			Q:     `"q":"125, Berliner, berlin"`,
			At:    nextbillionai.String(`"at": "52.5308,13.3856"`),
			In:    nextbillionai.String(`"in":"countryCode:CAN,MEX,USA"`),
			Lang:  nextbillionai.String(`"lang":"en"`),
			Limit: nextbillionai.Int(0),
		}},
	})
	if err != nil {
		var apierr *nextbillionai.Error
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
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Geocode.StructuredGet(context.TODO(), nextbillionai.GeocodeStructuredGetParams{
		CountryCode: "countryCode",
		Key:         "key=API_KEY",
		At:          nextbillionai.String("at=52.5308,13.3856"),
		City:        nextbillionai.String("city"),
		County:      nextbillionai.String("county"),
		HouseNumber: nextbillionai.String("houseNumber"),
		In:          nextbillionai.String("in=circle:52.53,13.38;r=10000"),
		Limit:       nextbillionai.Int(0),
		PostalCode:  nextbillionai.String("postalCode"),
		State:       nextbillionai.String("state"),
		Street:      nextbillionai.String("street"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
