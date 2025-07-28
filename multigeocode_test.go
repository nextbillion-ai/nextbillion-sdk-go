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

func TestMultigeocodeSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Search(context.TODO(), nextbillionsdk.MultigeocodeSearchParams{
		Key: "key=API_KEY",
		At: nextbillionsdk.MultigeocodeSearchParamsAt{
			Lat: 0,
			Lng: 0,
		},
		Query:       "“query”: “Taj Mahal”",
		City:        nextbillionsdk.String("“city”: “Glendale”"),
		Country:     nextbillionsdk.String("“country”:”IND”"),
		District:    nextbillionsdk.String("“district”: “City Center”"),
		Limit:       nextbillionsdk.Int(0),
		Radius:      nextbillionsdk.String("“radius”: “10m”"),
		State:       nextbillionsdk.String("“state”: “California”"),
		Street:      nextbillionsdk.String("“street”: “Americana Way”"),
		SubDistrict: nextbillionsdk.String("“subDistrict”: “Golkonda”"),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
