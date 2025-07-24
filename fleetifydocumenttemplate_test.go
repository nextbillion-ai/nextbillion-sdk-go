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

func TestFleetifyDocumentTemplateNew(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.New(context.TODO(), nextbillionsdk.FleetifyDocumentTemplateNewParams{
		Key: "key",
		Content: []nextbillionsdk.DocumentTemplateContentRequestParam{{
			Label: `"label": "Specify Completion Time"`,
			Type:  nextbillionsdk.DocumentTemplateContentRequestTypeString,
			Meta: nextbillionsdk.DocumentTemplateContentRequestMetaParam{
				Options: []nextbillionsdk.DocumentTemplateContentRequestMetaOptionParam{{
					Label: `"label": "Option 1"`,
					Value: `"value": "Car"`,
				}},
			},
			Name:     nextbillionsdk.String(`"name" : "Completion DateTime"`),
			Required: nextbillionsdk.Bool(true),
			Validation: nextbillionsdk.DocumentTemplateContentRequestValidationParam{
				Max:      nextbillionsdk.Int(0),
				MaxItems: nextbillionsdk.Int(0),
				Min:      nextbillionsdk.Int(0),
				MinItems: nextbillionsdk.Int(0),
			},
		}},
		Name: "name",
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyDocumentTemplateGet(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Get(
		context.TODO(),
		"id",
		nextbillionsdk.FleetifyDocumentTemplateGetParams{
			Key: "key",
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

func TestFleetifyDocumentTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Update(
		context.TODO(),
		"id",
		nextbillionsdk.FleetifyDocumentTemplateUpdateParams{
			Key: "key",
			Content: []nextbillionsdk.DocumentTemplateContentRequestParam{{
				Label: `"label": "Specify Completion Time"`,
				Type:  nextbillionsdk.DocumentTemplateContentRequestTypeString,
				Meta: nextbillionsdk.DocumentTemplateContentRequestMetaParam{
					Options: []nextbillionsdk.DocumentTemplateContentRequestMetaOptionParam{{
						Label: `"label": "Option 1"`,
						Value: `"value": "Car"`,
					}},
				},
				Name:     nextbillionsdk.String(`"name" : "Completion DateTime"`),
				Required: nextbillionsdk.Bool(true),
				Validation: nextbillionsdk.DocumentTemplateContentRequestValidationParam{
					Max:      nextbillionsdk.Int(0),
					MaxItems: nextbillionsdk.Int(0),
					Min:      nextbillionsdk.Int(0),
					MinItems: nextbillionsdk.Int(0),
				},
			}},
			Name: nextbillionsdk.String("name"),
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

func TestFleetifyDocumentTemplateList(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.List(context.TODO(), nextbillionsdk.FleetifyDocumentTemplateListParams{
		Key: "key",
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyDocumentTemplateDelete(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Delete(
		context.TODO(),
		"id",
		nextbillionsdk.FleetifyDocumentTemplateDeleteParams{
			Key: "key",
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
