// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nextbillionsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nextbillion-ai/nextbillion-sdk-go"
	"github.com/nextbillion-ai/nextbillion-sdk-go/internal/testutil"
	"github.com/nextbillion-ai/nextbillion-sdk-go/option"
)

func TestOptimizationV2GetResult(t *testing.T) {
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
	_, err := client.Optimization.V2.GetResult(context.TODO(), nextbillionsdk.OptimizationV2GetResultParams{
		ID:  "id",
		Key: "key=API_KEY",
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOptimizationV2SubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.V2.Submit(context.TODO(), nextbillionsdk.OptimizationV2SubmitParams{
		Key: "key=API_KEY",
		Locations: nextbillionsdk.OptimizationV2SubmitParamsLocations{
			Location:   []string{"string"},
			ID:         nextbillionsdk.Int(0),
			Approaches: []string{"unrestricted"},
		},
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
		CostMatrix: [][]int64{{0}},
		Depots: []nextbillionsdk.OptimizationV2SubmitParamsDepot{{
			ID:            `"id":"depot 1"`,
			LocationIndex: 0,
			Description:   nextbillionsdk.String("“description”:”Los_Angeles_depot”"),
			Service:       nextbillionsdk.Int(0),
			TimeWindows:   [][]int64{{0}},
		}},
		Description:        nextbillionsdk.String(`"description": "Sample Optimization"`),
		DistanceMatrix:     [][]int64{{0}},
		DurationMatrix:     [][]int64{{0}},
		ExistingSolutionID: nextbillionsdk.String("existing_solution_id"),
		Jobs: []nextbillionsdk.JobParam{{
			ID:                    `"id":"Job 1"`,
			LocationIndex:         0,
			Delivery:              []int64{0},
			DepotIDs:              []string{"string"},
			Description:           nextbillionsdk.String("description"),
			FollowLifoOrder:       nextbillionsdk.Bool(true),
			IncompatibleLoadTypes: []string{"string"},
			JointOrder:            nextbillionsdk.Int(0),
			LoadTypes:             []string{"string"},
			MaxVisitLateness:      nextbillionsdk.Int(0),
			Metadata:              "{\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
			OutsourcingCost:       nextbillionsdk.Int(0),
			Pickup:                []int64{0},
			Priority:              nextbillionsdk.Int(0),
			Revenue:               nextbillionsdk.Int(0),
			SequenceOrder:         nextbillionsdk.Int(0),
			Service:               nextbillionsdk.Int(0),
			Setup:                 nextbillionsdk.Int(0),
			Skills:                []int64{1},
			TimeWindows:           [][]int64{{0}},
			Volume: nextbillionsdk.JobVolumeParam{
				Alignment: "strict",
				Depth:     nextbillionsdk.Float(0),
				Height:    nextbillionsdk.Float(0),
				Width:     nextbillionsdk.Float(0),
			},
			Zones: []int64{0},
		}},
		Options: nextbillionsdk.OptimizationV2SubmitParamsOptions{
			Constraint: nextbillionsdk.OptimizationV2SubmitParamsOptionsConstraint{
				MaxActivityWaitingTime: nextbillionsdk.Int(0),
				MaxVehicleOvertime:     nextbillionsdk.Int(0),
				MaxVisitLateness:       nextbillionsdk.Int(0),
			},
			Grouping: nextbillionsdk.OptimizationV2SubmitParamsOptionsGrouping{
				OrderGrouping: nextbillionsdk.OptimizationV2SubmitParamsOptionsGroupingOrderGrouping{
					GroupingDiameter: nextbillionsdk.Float(0),
				},
				ProximityFactor: nextbillionsdk.Float(0),
				RouteGrouping: nextbillionsdk.OptimizationV2SubmitParamsOptionsGroupingRouteGrouping{
					PenaltyFactor: nextbillionsdk.Float(0),
					ZoneDiameter:  nextbillionsdk.Float(0),
					ZoneSource:    "system_generated",
				},
			},
			Objective: nextbillionsdk.OptimizationV2SubmitParamsOptionsObjective{
				AllowEarlyArrival: nextbillionsdk.Bool(true),
				Custom: nextbillionsdk.OptimizationV2SubmitParamsOptionsObjectiveCustom{
					Type:  "min",
					Value: "vehicles",
				},
				MinimiseNumDepots: nextbillionsdk.Bool(true),
				SolverMode:        "flexible",
				SolvingTimeLimit:  nextbillionsdk.Int(0),
				TravelCost:        "duration",
			},
			Routing: nextbillionsdk.OptimizationV2SubmitParamsOptionsRouting{
				Allow:            []string{"taxi"},
				Avoid:            []string{"toll"},
				Context:          "avgspeed",
				CrossBorder:      nextbillionsdk.Bool(true),
				DisableCache:     nextbillionsdk.Bool(true),
				HazmatType:       []string{"general"},
				Mode:             "car",
				Profiles:         "\"profiles\":{\n    \"mini-van\":{\n        \"mode\": \"car\",\n        \"avoid\":[\"highway, toll\"]\n        },\n    \"trailer\":{\n        \"mode\": \"truck\",\n        \"truck_weight\":12000,\n        \"truck_size\":\"200, 210, 600\",\n        \"hazmat_type\": [\"general\", \"harmful_to_water\"]\n        }\n    }\n",
				TrafficTimestamp: nextbillionsdk.Int(0),
				TruckAxleLoad:    nextbillionsdk.Float(0),
				TruckSize:        nextbillionsdk.String(`"truck_size":"200,210,600"`),
				TruckWeight:      nextbillionsdk.Int(0),
			},
		},
		Relations: []nextbillionsdk.OptimizationV2SubmitParamsRelation{{
			Steps: []nextbillionsdk.OptimizationV2SubmitParamsRelationStep{{
				Type: "start",
				ID:   nextbillionsdk.String(`"id":"Job 1"`),
			}},
			Type:        "in_same_route",
			ID:          nextbillionsdk.Int(0),
			MaxDuration: nextbillionsdk.Int(0),
			MinDuration: nextbillionsdk.Int(0),
			Vehicle:     nextbillionsdk.String(`"vehicle": "Vehicle 10"`),
		}},
		Shipments: []nextbillionsdk.ShipmentParam{{
			Delivery: nextbillionsdk.ShipmentDeliveryParam{
				ID:               `"id":"Shipment Delivery 1"`,
				LocationIndex:    0,
				Description:      nextbillionsdk.String("description"),
				MaxVisitLateness: nextbillionsdk.Int(0),
				Metadata:         "{\n  \"notes\": \"dropoff at the patio\",\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
				SequenceOrder:    nextbillionsdk.Int(0),
				Service:          nextbillionsdk.Int(0),
				Setup:            nextbillionsdk.Int(0),
				TimeWindows:      [][]int64{{0}},
			},
			Pickup: nextbillionsdk.ShipmentPickupParam{
				ID:               `"id": "Shipment Pickup 1"`,
				LocationIndex:    0,
				Description:      nextbillionsdk.String("description"),
				MaxVisitLateness: nextbillionsdk.Int(0),
				Metadata:         "{\n  \"notes\": \"involves fragile items\",\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
				SequenceOrder:    nextbillionsdk.Int(0),
				Service:          nextbillionsdk.Int(0),
				Setup:            nextbillionsdk.Int(0),
				TimeWindows:      [][]int64{{0}},
			},
			Amount:                []int64{0},
			FollowLifoOrder:       nextbillionsdk.Bool(true),
			IncompatibleLoadTypes: []string{"string"},
			JointOrder:            nextbillionsdk.Int(0),
			LoadTypes:             []string{"string"},
			MaxTimeInVehicle:      nextbillionsdk.Int(0),
			OutsourcingCost:       nextbillionsdk.Int(0),
			Priority:              nextbillionsdk.Int(0),
			Revenue:               nextbillionsdk.Int(0),
			Skills:                []int64{0},
			Volume: nextbillionsdk.ShipmentVolumeParam{
				Alignment: "strict",
				Depth:     nextbillionsdk.Float(0),
				Height:    nextbillionsdk.Float(0),
				Width:     nextbillionsdk.Float(0),
			},
			Zones: []int64{0},
		}},
		Solution: []nextbillionsdk.OptimizationV2SubmitParamsSolution{{
			Cost: 0,
			Steps: []nextbillionsdk.OptimizationV2SubmitParamsSolutionStep{{
				ID:            `"id": "Job 10"`,
				Arrival:       0,
				Type:          "start",
				Description:   nextbillionsdk.String("description"),
				Distance:      nextbillionsdk.Int(0),
				Duration:      nextbillionsdk.Int(0),
				Load:          []int64{0},
				Location:      []float64{0},
				LocationIndex: nextbillionsdk.Int(0),
				Service:       nextbillionsdk.Int(0),
				Setup:         nextbillionsdk.Int(0),
				WaitingTime:   nextbillionsdk.Int(0),
			}},
			Vehicle:     "vehicle",
			Delivery:    []int64{0},
			Description: nextbillionsdk.String("description"),
			Distance:    nextbillionsdk.Int(0),
			Duration:    nextbillionsdk.Int(0),
			Geometry:    nextbillionsdk.String(`"geometry": "}ebGgcsxRE?CuDOYCYAG???"`),
			Pickup:      []int64{0},
			Priority:    nextbillionsdk.Int(0),
			Service:     nextbillionsdk.Int(0),
			Setup:       nextbillionsdk.Int(0),
			WaitingTime: nextbillionsdk.Int(0),
		}},
		Unassigned: nextbillionsdk.OptimizationV2SubmitParamsUnassigned{
			Jobs:      []string{"string"},
			Shipments: [][]string{{"string"}},
		},
		Zones: []nextbillionsdk.OptimizationV2SubmitParamsZone{{
			ID:         0,
			GeofenceID: nextbillionsdk.String("geofence_id"),
			Geometry: nextbillionsdk.OptimizationV2SubmitParamsZoneGeometry{
				Coordinates: [][]float64{{0}},
				Description: nextbillionsdk.String("description"),
				Type:        "Polygon",
			},
		}},
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
