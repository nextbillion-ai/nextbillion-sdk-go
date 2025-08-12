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

func TestFleetifyDocumentTemplateNew(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.New(context.TODO(), nextbillionai.FleetifyDocumentTemplateNewParams{
		Key: "key",
		Content: []nextbillionai.DocumentTemplateContentRequestParam{{
			Label: `"label": "Specify Completion Time"`,
			Type:  nextbillionai.DocumentTemplateContentRequestTypeString,
			Meta: nextbillionai.DocumentTemplateContentRequestMetaParam{
				Options: []nextbillionai.DocumentTemplateContentRequestMetaOptionParam{{
					Label: `"label": "Option 1"`,
					Value: `"value": "Car"`,
				}},
			},
			Name:     nextbillionai.String(`"name" : "Completion DateTime"`),
			Required: nextbillionai.Bool(true),
			Validation: nextbillionai.DocumentTemplateContentRequestValidationParam{
				Max:      nextbillionai.Int(0),
				MaxItems: nextbillionai.Int(0),
				Min:      nextbillionai.Int(0),
				MinItems: nextbillionai.Int(0),
			},
		}},
		Name: "name",
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyDocumentTemplateGet(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Get(
		context.TODO(),
		"id",
		nextbillionai.FleetifyDocumentTemplateGetParams{
			Key: "key",
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

func TestFleetifyDocumentTemplateUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Update(
		context.TODO(),
		"id",
		nextbillionai.FleetifyDocumentTemplateUpdateParams{
			Key: "key",
			Content: []nextbillionai.DocumentTemplateContentRequestParam{{
				Label: `"label": "Specify Completion Time"`,
				Type:  nextbillionai.DocumentTemplateContentRequestTypeString,
				Meta: nextbillionai.DocumentTemplateContentRequestMetaParam{
					Options: []nextbillionai.DocumentTemplateContentRequestMetaOptionParam{{
						Label: `"label": "Option 1"`,
						Value: `"value": "Car"`,
					}},
				},
				Name:     nextbillionai.String(`"name" : "Completion DateTime"`),
				Required: nextbillionai.Bool(true),
				Validation: nextbillionai.DocumentTemplateContentRequestValidationParam{
					Max:      nextbillionai.Int(0),
					MaxItems: nextbillionai.Int(0),
					Min:      nextbillionai.Int(0),
					MinItems: nextbillionai.Int(0),
				},
			}},
			Name: nextbillionai.String("name"),
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

func TestFleetifyDocumentTemplateList(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.List(context.TODO(), nextbillionai.FleetifyDocumentTemplateListParams{
		Key: "key",
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFleetifyDocumentTemplateDelete(t *testing.T) {
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
	_, err := client.Fleetify.DocumentTemplates.Delete(
		context.TODO(),
		"id",
		nextbillionai.FleetifyDocumentTemplateDeleteParams{
			Key: "key",
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
