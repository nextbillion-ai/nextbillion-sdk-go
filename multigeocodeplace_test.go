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

func TestMultigeocodePlaceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Place.New(context.TODO(), nextbillionai.MultigeocodePlaceNewParams{
		Key: "key=API_KEY",
		Place: []nextbillionai.MultigeocodePlaceNewParamsPlace{{
			Geopoint: nextbillionai.MultigeocodePlaceNewParamsPlaceGeopoint{
				Lat: nextbillionai.Float(0),
				Lng: nextbillionai.Float(0),
			},
			Address:  nextbillionai.String("address"),
			Building: nextbillionai.String("building"),
			City:     nextbillionai.String("city"),
			Country:  nextbillionai.String(`"country":"IND"`),
			District: nextbillionai.String("district"),
			House:    nextbillionai.String("house"),
			Poi: nextbillionai.MultigeocodePlaceNewParamsPlacePoi{
				Title: nextbillionai.String("title"),
			},
			PostalCode:  nextbillionai.String("postalCode"),
			State:       nextbillionai.String("state"),
			Street:      nextbillionai.String("street"),
			SubDistrict: nextbillionai.String("subDistrict"),
		}},
		DataSource: nextbillionai.MultigeocodePlaceNewParamsDataSource{
			RefID:  nextbillionai.String("refId"),
			Source: nextbillionai.String("source"),
			Status: "enable",
		},
		Force: nextbillionai.Bool(true),
		Score: nextbillionai.Int(0),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceGet(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Get(
		context.TODO(),
		"docId",
		nextbillionai.MultigeocodePlaceGetParams{
			Key: "key=API_KEY",
		},
	)
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Update(
		context.TODO(),
		"docId",
		nextbillionai.MultigeocodePlaceUpdateParams{
			Key: "key=API_KEY",
			DataSource: nextbillionai.MultigeocodePlaceUpdateParamsDataSource{
				RefID:  nextbillionai.String("refId"),
				Source: nextbillionai.String("source"),
				Status: "enable",
			},
			Place: []nextbillionai.PlaceItemParam{{
				Address:  nextbillionai.String("address"),
				Building: nextbillionai.String("building"),
				City:     nextbillionai.String("city"),
				Country:  nextbillionai.String("country"),
				District: nextbillionai.String("district"),
				Geopoint: nextbillionai.PlaceItemGeopointParam{
					Lat: nextbillionai.Float(0),
					Lng: nextbillionai.Float(0),
				},
				House: nextbillionai.String("house"),
				Poi: nextbillionai.PlaceItemPoiParam{
					Title: nextbillionai.String("title"),
				},
				PostalCode:  nextbillionai.String("postalCode"),
				State:       nextbillionai.String("state"),
				Street:      nextbillionai.String("street"),
				SubDistrict: nextbillionai.String("subDistrict"),
			}},
			Score: nextbillionai.Int(0),
		},
	)
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMultigeocodePlaceDelete(t *testing.T) {
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
	_, err := client.Multigeocode.Place.Delete(
		context.TODO(),
		"docId",
		nextbillionai.MultigeocodePlaceDeleteParams{
			Key: "key=API_KEY",
		},
	)
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
