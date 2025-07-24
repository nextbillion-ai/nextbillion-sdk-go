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

func TestMapStaticAutoGetAutoImageWithOptionalParams(t *testing.T) {
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
	err := client.Maps.Static.Auto.GetAutoImage(
		context.TODO(),
		nextbillionsdk.MapStaticAutoGetAutoImageParamsFormatPng,
		nextbillionsdk.MapStaticAutoGetAutoImageParams{
			MapID:       nextbillionsdk.MapStaticAutoGetAutoImageParamsMapIDHybrid,
			Width:       0,
			Height:      0,
			Scale:       "scale",
			Key:         "key=API_KEY",
			Attribution: nextbillionsdk.MapStaticAutoGetAutoImageParamsAttributionBottomright,
			Markers:     nextbillionsdk.String("markers=14.4,50.1,red|8.6,47.4,blue"),
			Padding:     nextbillionsdk.Float(0),
			Path:        nextbillionsdk.String("path=stroke:green|width:3|fill:none|5.9,45.8|5.9,47.8|10.5,47.8|10.5,45.8|5.9,45.8"),
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
