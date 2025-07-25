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

func TestOptimizationDriverAssignmentAssignDriversWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.DriverAssignment.AssignDrivers(context.TODO(), nextbillionsdk.OptimizationDriverAssignmentAssignDriversParams{
		Key: "key=API_KEY",
		Filter: nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsFilter{
			DrivingDistance: nextbillionsdk.Float(0),
			PickupEta:       nextbillionsdk.Int(0),
			Radius:          nextbillionsdk.Float(0),
		},
		Orders: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrder{{
			ID: "id",
			Pickup: nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderPickup{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			},
			Attributes: map[string]interface{}{},
			Dropoffs: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderDropoff{{
				Lat: nextbillionsdk.Float(0),
				Lng: nextbillionsdk.Float(0),
			}},
			Priority:    nextbillionsdk.Int(0),
			ServiceTime: nextbillionsdk.Int(0),
			VehiclePreferences: nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferences{
				ExcludeAllOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesExcludeAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"<"`,
					Value:     `"value": "4"`,
				}},
				RequiredAllOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAllOfAttribute{{
					Attribute: `"attribute": "driver_rating"`,
					Operator:  `"operator":"=="`,
					Value:     `"value": "4"`,
				}},
				RequiredAnyOfAttributes: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOrderVehiclePreferencesRequiredAnyOfAttribute{{
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
		Options: nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOptions{
			AlternateAssignments: nextbillionsdk.Int(0),
			DropoffDetails:       nextbillionsdk.Bool(true),
			OrderAttributePriorityMappings: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOptionsOrderAttributePriorityMapping{{
				Attribute: `"attribute": "driver_rating"`,
				Operator:  `"operator":"=="`,
				Priority:  "priority",
				Value:     `"value": "4"`,
			}},
			TravelCost: "driving_eta",
			VehicleAttributePriorityMappings: []nextbillionsdk.OptimizationDriverAssignmentAssignDriversParamsOptionsVehicleAttributePriorityMapping{{
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
