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

func TestRestrictionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.New(
		context.TODO(),
		nextbillionai.RestrictionNewParamsRestrictionTypeTurn,
		nextbillionai.RestrictionNewParams{
			Key: "key=API_KEY",
			RichGroupRequest: nextbillionai.RichGroupRequestParam{
				Area:      "area",
				Name:      "name",
				Comment:   nextbillionai.String("comment"),
				Direction: nextbillionai.RichGroupRequestDirectionForward,
				EndTime:   nextbillionai.Float(0),
				Geofence:  [][]float64{{0}},
				Height:    nextbillionai.Int(0),
				Length:    nextbillionai.Int(0),
				Mode:      []string{"0w"},
				RepeatOn:  nextbillionai.String(`repeatOn="Mo-Fr 07:00-09:00,17:00-19:00"`),
				Segments: []nextbillionai.RichGroupRequestSegmentParam{{
					From: nextbillionai.Float(0),
					To:   nextbillionai.Float(0),
				}},
				Speed:      nextbillionai.Float(0),
				SpeedLimit: nextbillionai.Float(0),
				StartTime:  nextbillionai.Float(0),
				Tracks:     [][]float64{{0}},
				Turns: []nextbillionai.RichGroupRequestTurnParam{{
					From: nextbillionai.Int(0),
					To:   nextbillionai.Int(0),
					Via:  nextbillionai.Int(0),
				}},
				Weight: nextbillionai.Int(0),
				Width:  nextbillionai.Int(0),
			},
			Latlon: nextbillionai.Bool(true),
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

func TestRestrictionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.Get(
		context.TODO(),
		0,
		nextbillionai.RestrictionGetParams{
			Key:       "key=API_KEY",
			Transform: nextbillionai.Bool(true),
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

func TestRestrictionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.Update(
		context.TODO(),
		0,
		nextbillionai.RestrictionUpdateParams{
			Key: "key=API_KEY",
			RichGroupRequest: nextbillionai.RichGroupRequestParam{
				Area:      "area",
				Name:      "name",
				Comment:   nextbillionai.String("comment"),
				Direction: nextbillionai.RichGroupRequestDirectionForward,
				EndTime:   nextbillionai.Float(0),
				Geofence:  [][]float64{{0}},
				Height:    nextbillionai.Int(0),
				Length:    nextbillionai.Int(0),
				Mode:      []string{"0w"},
				RepeatOn:  nextbillionai.String(`repeatOn="Mo-Fr 07:00-09:00,17:00-19:00"`),
				Segments: []nextbillionai.RichGroupRequestSegmentParam{{
					From: nextbillionai.Float(0),
					To:   nextbillionai.Float(0),
				}},
				Speed:      nextbillionai.Float(0),
				SpeedLimit: nextbillionai.Float(0),
				StartTime:  nextbillionai.Float(0),
				Tracks:     [][]float64{{0}},
				Turns: []nextbillionai.RichGroupRequestTurnParam{{
					From: nextbillionai.Int(0),
					To:   nextbillionai.Int(0),
					Via:  nextbillionai.Int(0),
				}},
				Weight: nextbillionai.Int(0),
				Width:  nextbillionai.Int(0),
			},
			Latlon: nextbillionai.Bool(true),
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

func TestRestrictionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.List(context.TODO(), nextbillionai.RestrictionListParams{
		Area:            "area",
		Key:             "key=API_KEY",
		Limit:           0,
		Offset:          0,
		Mode:            nextbillionai.RestrictionListParamsMode0w,
		Name:            nextbillionai.String("name"),
		RestrictionType: nextbillionai.RestrictionListParamsRestrictionTypeTurn,
		Source:          nextbillionai.RestrictionListParamsSourceRrt,
		State:           nextbillionai.RestrictionListParamsStateEnabled,
		Status:          nextbillionai.RestrictionListParamsStatusActive,
		Transform:       nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRestrictionDelete(t *testing.T) {
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
	_, err := client.Restrictions.Delete(
		context.TODO(),
		0,
		nextbillionai.RestrictionDeleteParams{
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

func TestRestrictionListByBboxWithOptionalParams(t *testing.T) {
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
	_, err := client.Restrictions.ListByBbox(context.TODO(), nextbillionai.RestrictionListByBboxParams{
		Key:             "key=API_KEY",
		MaxLat:          0,
		MaxLon:          0,
		MinLat:          0,
		MinLon:          0,
		Mode:            []string{"0w"},
		RestrictionType: nextbillionai.RestrictionListByBboxParamsRestrictionTypeTurn,
		Source:          nextbillionai.RestrictionListByBboxParamsSourceRrt,
		State:           nextbillionai.RestrictionListByBboxParamsStateEnabled,
		Status:          nextbillionai.RestrictionListByBboxParamsStatusActive,
		Transform:       nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRestrictionSetState(t *testing.T) {
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
	_, err := client.Restrictions.SetState(
		context.TODO(),
		0,
		nextbillionai.RestrictionSetStateParams{
			Key:   "key=API_KEY",
			State: nextbillionai.RestrictionSetStateParamsStateEnabled,
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
