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

func TestRestrictionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.New(
		context.TODO(),
		nextbillionsdk.RestrictionNewParamsRestrictionTypeTurn,
		nextbillionsdk.RestrictionNewParams{
			Key: "key=API_KEY",
			RichGroupRequest: nextbillionsdk.RichGroupRequestParam{
				Area:      "area",
				Name:      "name",
				Comment:   nextbillionsdk.String("comment"),
				Direction: nextbillionsdk.RichGroupRequestDirectionForward,
				EndTime:   nextbillionsdk.Float(0),
				Geofence:  [][]float64{{0}},
				Height:    nextbillionsdk.Int(0),
				Length:    nextbillionsdk.Int(0),
				Mode:      []string{"0w"},
				RepeatOn:  nextbillionsdk.String(`repeatOn="Mo-Fr 07:00-09:00,17:00-19:00"`),
				Segments: []nextbillionsdk.RichGroupRequestSegmentParam{{
					From: nextbillionsdk.Float(0),
					To:   nextbillionsdk.Float(0),
				}},
				Speed:      nextbillionsdk.Float(0),
				SpeedLimit: nextbillionsdk.Float(0),
				StartTime:  nextbillionsdk.Float(0),
				Tracks:     [][]float64{{0}},
				Turns: []nextbillionsdk.RichGroupRequestTurnParam{{
					From: nextbillionsdk.Int(0),
					To:   nextbillionsdk.Int(0),
					Via:  nextbillionsdk.Int(0),
				}},
				Weight: nextbillionsdk.Int(0),
				Width:  nextbillionsdk.Int(0),
			},
			Latlon: nextbillionsdk.Bool(true),
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

func TestRestrictionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.Get(
		context.TODO(),
		0,
		nextbillionsdk.RestrictionGetParams{
			Key:       "key=API_KEY",
			Transform: nextbillionsdk.Bool(true),
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

func TestRestrictionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.Update(
		context.TODO(),
		0,
		nextbillionsdk.RestrictionUpdateParams{
			Key: "key=API_KEY",
			RichGroupRequest: nextbillionsdk.RichGroupRequestParam{
				Area:      "area",
				Name:      "name",
				Comment:   nextbillionsdk.String("comment"),
				Direction: nextbillionsdk.RichGroupRequestDirectionForward,
				EndTime:   nextbillionsdk.Float(0),
				Geofence:  [][]float64{{0}},
				Height:    nextbillionsdk.Int(0),
				Length:    nextbillionsdk.Int(0),
				Mode:      []string{"0w"},
				RepeatOn:  nextbillionsdk.String(`repeatOn="Mo-Fr 07:00-09:00,17:00-19:00"`),
				Segments: []nextbillionsdk.RichGroupRequestSegmentParam{{
					From: nextbillionsdk.Float(0),
					To:   nextbillionsdk.Float(0),
				}},
				Speed:      nextbillionsdk.Float(0),
				SpeedLimit: nextbillionsdk.Float(0),
				StartTime:  nextbillionsdk.Float(0),
				Tracks:     [][]float64{{0}},
				Turns: []nextbillionsdk.RichGroupRequestTurnParam{{
					From: nextbillionsdk.Int(0),
					To:   nextbillionsdk.Int(0),
					Via:  nextbillionsdk.Int(0),
				}},
				Weight: nextbillionsdk.Int(0),
				Width:  nextbillionsdk.Int(0),
			},
			Latlon: nextbillionsdk.Bool(true),
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

func TestRestrictionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.List(context.TODO(), nextbillionsdk.RestrictionListParams{
		Area:            "area",
		Key:             "key=API_KEY",
		Limit:           0,
		Offset:          0,
		Mode:            nextbillionsdk.RestrictionListParamsMode0w,
		Name:            nextbillionsdk.String("name"),
		RestrictionType: nextbillionsdk.RestrictionListParamsRestrictionTypeTurn,
		Source:          nextbillionsdk.RestrictionListParamsSourceRrt,
		State:           nextbillionsdk.RestrictionListParamsStateEnabled,
		Status:          nextbillionsdk.RestrictionListParamsStatusActive,
		Transform:       nextbillionsdk.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRestrictionDelete(t *testing.T) {
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
	_, err := client.Restrictions.Delete(
		context.TODO(),
		0,
		nextbillionsdk.RestrictionDeleteParams{
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

func TestRestrictionListByBboxWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.ListByBbox(context.TODO(), nextbillionsdk.RestrictionListByBboxParams{
		Key:             "key=API_KEY",
		MaxLat:          0,
		MaxLon:          0,
		MinLat:          0,
		MinLon:          0,
		Mode:            []string{"0w"},
		RestrictionType: nextbillionsdk.RestrictionListByBboxParamsRestrictionTypeTurn,
		Source:          nextbillionsdk.RestrictionListByBboxParamsSourceRrt,
		State:           nextbillionsdk.RestrictionListByBboxParamsStateEnabled,
		Status:          nextbillionsdk.RestrictionListByBboxParamsStatusActive,
		Transform:       nextbillionsdk.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRestrictionSetState(t *testing.T) {
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
	_, err := client.Restrictions.SetState(
		context.TODO(),
		0,
		nextbillionsdk.RestrictionSetStateParams{
			Key:   "key=API_KEY",
			State: nextbillionsdk.RestrictionSetStateParamsStateEnabled,
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
