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

func TestOptimizationDriverAssignmentAssignWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.DriverAssignment.Assign(context.TODO(), nextbillionsdk.OptimizationDriverAssignmentAssignParams{
		Key: "key=API_KEY",
		Filter: nextbillionsdk.OptimizationDriverAssignmentAssignParamsFilter{
			DrivingDistance: nextbillionsdk.Float(0),
			PickupEta:       nextbillionsdk.Int(0),
			Radius:          nextbillionsdk.Float(0),
		},
		Orders: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrder{{
			ID: "id",
			Pickup: nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderPickup{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			},
			Attributes: map[string]interface{}{},
			Dropoffs: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderDropoff{{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			}},
			Priority:    nextbillionsdk.Int(0),
			ServiceTime: nextbillionsdk.Int(0),
			VehiclePreferences: nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferences{
				ExcludeAllOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesExcludeAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"<"`,
					Value:     `"value": "4"`,
				}},
				RequiredAllOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesRequiredAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"=="`,
					Value:     `"value": "4"`,
				}},
				RequiredAnyOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOrderVehiclePreferencesRequiredAnyOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":">"`,
					Value:     `"value": "4"`,
				}},
			},
		}},
		Vehicles: []nextbillionsdk.VehicleParam{{
			ID: "id",
			Location: nextbillionsdk.VehicleLocationParam{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			},
			Attributes: "\"attributes\":{\n    \"driver_rating\": \"4.0\",\n    \"trip_types\": \"premium\"\n  }",
			Priority:   nextbillionsdk.Int(0),
			RemainingWaypoints: []nextbillionsdk.LocationParam{{
				Lat: -90,
				Lon: -180,
			}},
		}},
		Options: nextbillionsdk.OptimizationDriverAssignmentAssignParamsOptions{
			AlternateAssignments: nextbillionsdk.Int(0),
			DropoffDetails:       nextbillionsdk.Bool(true),
			OrderAttributePriorityMappings: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOptionsOrderAttributePriorityMapping{{
				Attribute: `"attribute": "driver_rating"`,
				Operator:  `"operator":"=="`,
				Priority:  "priority",
				Value:     `"value": "4"`,
			}},
			TravelCost: "driving_eta",
			VehicleAttributePriorityMappings: []nextbillionsdk.OptimizationDriverAssignmentAssignParamsOptionsVehicleAttributePriorityMapping{{
				Attribute: `"attribute": "driver_rating"`,
				Operator:  `"operator":"=="`,
				Priority:  "priority",
				Value:     `"value": "4"`,
			}},
		},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
