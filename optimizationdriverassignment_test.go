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

func TestOptimizationDriverAssignmentAssignWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.DriverAssignment.Assign(context.TODO(), nextbillionai.OptimizationDriverAssignmentAssignParams{
		Key: "key=API_KEY",
		Filter: nextbillionai.OptimizationDriverAssignmentAssignParamsFilter{
			DrivingDistance: nextbillionai.Float(0),
			PickupEta:       nextbillionai.Int(0),
			Radius:          nextbillionai.Float(0),
		},
		Orders: []nextbillionai.OptimizationDriverAssignmentAssignParamsOrder{{
			ID: "id",
			Pickup: nextbillionai.OptimizationDriverAssignmentAssignParamsOrderPickup{
				Lat: nextbillionai.Float(0),
				Lng: nextbillionai.Float(0),
			},
			Attributes: map[string]interface{}{},
			Dropoffs: []nextbillionai.OptimizationDriverAssignmentAssignParamsOrderDropoff{{
				Lat: nextbillionai.Float(0),
				Lng: nextbillionai.Float(0),
			}},
			Priority:    nextbillionai.Int(0),
			ServiceTime: nextbillionai.Int(0),
			VehiclePreferences: nextbillionai.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferences{
				ExcludeAllOfAttributes: []nextbillionai.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesExcludeAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"<"`,
					Value:     `"value": "4"`,
				}},
				RequiredAllOfAttributes: []nextbillionai.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesRequiredAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"=="`,
					Value:     `"value": "4"`,
				}},
				RequiredAnyOfAttributes: []nextbillionai.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesRequiredAnyOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":">"`,
					Value:     `"value": "4"`,
				}},
			},
		}},
		Vehicles: []nextbillionai.VehicleParam{{
			ID: "id",
			Location: nextbillionai.LocationParam{
				Lat: -90,
				Lon: -180,
			},
			Attributes: "\"attributes\":{\n    \"driver_rating\": \"4.0\",\n    \"trip_types\": \"premium\"\n  }",
			Priority:   nextbillionai.Int(0),
			RemainingWaypoints: []nextbillionai.LocationParam{{
				Lat: -90,
				Lon: -180,
			}},
		}},
		Options: nextbillionai.OptimizationDriverAssignmentAssignParamsOptions{
			AlternateAssignments: nextbillionai.Int(0),
			DropoffDetails:       nextbillionai.Bool(true),
			OrderAttributePriorityMappings: []nextbillionai.OptimizationDriverAssignmentAssignParamsOptionsOrderAttributePriorityMapping{{
				Attribute: `"attribute": "driver_rating"`,
				Operator:  `"operator":"=="`,
				Priority:  "priority",
				Value:     `"value": "4"`,
			}},
			TravelCost: "driving_eta",
			VehicleAttributePriorityMappings: []nextbillionai.OptimizationDriverAssignmentAssignParamsOptionsVehicleAttributePriorityMapping{{
				Attribute: `"attribute": "driver_rating"`,
				Operator:  `"operator":"=="`,
				Priority:  "priority",
				Value:     `"value": "4"`,
			}},
		},
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
