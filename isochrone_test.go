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

func TestIsochroneComputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Isochrone.Compute(context.TODO(), nextbillionai.IsochroneComputeParams{
		ContoursMeters:  0,
		ContoursMinutes: 0,
		Coordinates:     "coordinates=1.29363713,103.8383112",
		Key:             "key=API_KEY",
		ContoursColors:  nextbillionai.String("contours_colors=ff0000,bf4040"),
		Denoise:         nextbillionai.Float(0),
		DepartureTime:   nextbillionai.Int(0),
		Generalize:      nextbillionai.Float(0),
		Mode:            nextbillionai.IsochroneComputeParamsModeCar,
		Polygons:        nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
