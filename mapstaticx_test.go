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

func TestMapStaticXGetBoundsImageWithOptionalParams(t *testing.T) {
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
	err := client.Maps.Static.X.GetBoundsImage(
		context.TODO(),
		"format",
		nextbillionsdk.MapStaticXGetBoundsImageParams{
			MapID:       nextbillionsdk.MapStaticXGetBoundsImageParamsMapIDHybrid,
			Miny:        0,
			Minx:        0,
			Maxy:        0,
			Maxx:        0,
			Width:       0,
			Height:      0,
			Scale:       "scale",
			Key:         "key=API_KEY",
			Attribution: nextbillionsdk.MapStaticXGetBoundsImageParamsAttributionBottomright,
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

func TestMapStaticXGetCenteredImageWithOptionalParams(t *testing.T) {
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
	err := client.Maps.Static.X.GetCenteredImage(
		context.TODO(),
		nextbillionsdk.MapStaticXGetCenteredImageParamsFormatPng,
		nextbillionsdk.MapStaticXGetCenteredImageParams{
			MapID:       nextbillionsdk.MapStaticXGetCenteredImageParamsMapIDHybrid,
			Lat:         0,
			Lon:         0,
			Zoom:        0,
			Width:       0,
			Height:      0,
			Scale:       "scale",
			Key:         "key=API_KEY",
			Attribution: nextbillionsdk.MapStaticXGetCenteredImageParamsAttributionBottomright,
			Markers:     nextbillionsdk.String("markers=14.4,50.1,red|8.6,47.4,blue"),
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
