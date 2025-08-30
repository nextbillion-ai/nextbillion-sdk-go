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

func TestPostalcodeGetCoordinatesWithOptionalParams(t *testing.T) {
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
	_, err := client.Postalcode.GetCoordinates(context.TODO(), nextbillionai.PostalcodeGetCoordinatesParams{
		Key: "key=API_KEY",
		At: nextbillionai.PostalcodeGetCoordinatesParamsAt{
			Lat: nextbillionai.Float(0),
			Lng: nextbillionai.Float(0),
		},
		Country:    nextbillionai.String("country"),
		Format:     nextbillionai.PostalcodeGetCoordinatesParamsFormatGeojson,
		Postalcode: nextbillionai.String("postalcode"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
