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

func TestRestrictionsItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.RestrictionsItems.List(context.TODO(), nextbillionai.RestrictionsItemListParams{
		MaxLat:          0,
		MaxLon:          0,
		MinLat:          0,
		MinLon:          0,
		GroupID:         nextbillionai.Float(0),
		Mode:            nextbillionai.RestrictionsItemListParamsMode0w,
		RestrictionType: nextbillionai.RestrictionsItemListParamsRestrictionTypeTurn,
		Source:          nextbillionai.String("source"),
		State:           nextbillionai.RestrictionsItemListParamsStateEnabled,
		Status:          nextbillionai.RestrictionsItemListParamsStatusActive,
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
