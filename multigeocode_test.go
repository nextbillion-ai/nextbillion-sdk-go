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

func TestMultigeocodeSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Multigeocode.Search(context.TODO(), nextbillionai.MultigeocodeSearchParams{
		Key: "key=API_KEY",
		At: nextbillionai.MultigeocodeSearchParamsAt{
			Lat: 0,
			Lng: 0,
		},
		Query:       "“query”: “Taj Mahal”",
		City:        nextbillionai.String("“city”: “Glendale”"),
		Country:     nextbillionai.String("“country”:”IND”"),
		District:    nextbillionai.String("“district”: “City Center”"),
		Limit:       nextbillionai.Int(0),
		Radius:      nextbillionai.String("“radius”: “10m”"),
		State:       nextbillionai.String("“state”: “California”"),
		Street:      nextbillionai.String("“street”: “Americana Way”"),
		SubDistrict: nextbillionai.String("“subDistrict”: “Golkonda”"),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
