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

func TestMultigeocodePlaceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Place.New(context.TODO(), nextbillionsdk.MultigeocodePlaceNewParams{
		Key: "key=API_KEY",
		Place: []nextbillionsdk.MultigeocodePlaceNewParamsPlace{{
			Geopoint: nextbillionsdk.MultigeocodePlaceNewParamsPlaceGeopoint{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			},
			Address:  nextbillionsdk.String("address"),
			Building: nextbillionsdk.String("building"),
			City:     nextbillionsdk.String("city"),
			Country:  nextbillionsdk.String(`"country":"IND"`),
			District: nextbillionsdk.String("district"),
			House:    nextbillionsdk.String("house"),
			Poi: nextbillionsdk.MultigeocodePlaceNewParamsPlacePoi{
				Title: nextbillionsdk.String("title"),
			},
			PostalCode:  nextbillionsdk.String("postalCode"),
			State:       nextbillionsdk.String("state"),
			Street:      nextbillionsdk.String("street"),
			SubDistrict: nextbillionsdk.String("subDistrict"),
		}},
		DataSource: nextbillionsdk.MultigeocodePlaceNewParamsDataSource{
			RefID:  nextbillionsdk.String("refId"),
			Source: nextbillionsdk.String("source"),
			Status: "enable",
		},
		Force: nextbillionsdk.Bool(true),
		Score: nextbillionsdk.Int(0),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceGet(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Get(
		context.TODO(),
		"docId",
		nextbillionsdk.MultigeocodePlaceGetParams{
			Key: "key=API_KEY",
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Update(
		context.TODO(),
		"docId",
		nextbillionsdk.MultigeocodePlaceUpdateParams{
			Key: "key=API_KEY",
			DataSource: nextbillionsdk.MultigeocodePlaceUpdateParamsDataSource{
				RefID:  nextbillionsdk.String("refId"),
				Source: nextbillionsdk.String("source"),
				Status: "enable",
			},
			Place: []nextbillionsdk.PlaceItemParam{{
				Address:  nextbillionsdk.String("address"),
				Building: nextbillionsdk.String("building"),
				City:     nextbillionsdk.String("city"),
				Country:  nextbillionsdk.String("country"),
				District: nextbillionsdk.String("district"),
				Geopoint: nextbillionsdk.PlaceItemGeopointParam{
					Lat: nextbillionsdk.Float(0),
					Lng: nextbillionsdk.Float(0),
				},
				House: nextbillionsdk.String("house"),
				Poi: nextbillionsdk.PlaceItemPoiParam{
					Title: nextbillionsdk.String("title"),
				},
				PostalCode:  nextbillionsdk.String("postalCode"),
				State:       nextbillionsdk.String("state"),
				Street:      nextbillionsdk.String("street"),
				SubDistrict: nextbillionsdk.String("subDistrict"),
			}},
			Score: nextbillionsdk.Int(0),
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceDelete(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Delete(
		context.TODO(),
		"docId",
		nextbillionsdk.MultigeocodePlaceDeleteParams{
			Key: "key=API_KEY",
		},
	)
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
