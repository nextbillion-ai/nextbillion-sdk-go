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

func TestOptimizationV2GetResult(t *testing.T) {
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
	_, err := client.Optimization.V2.GetResult(context.TODO(), nextbillionai.OptimizationV2GetResultParams{
		ID:  "id",
		Key: "key=API_KEY",
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOptimizationV2SubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.V2.Submit(context.TODO(), nextbillionai.OptimizationV2SubmitParams{
		Key: "key=API_KEY",
		Locations: nextbillionai.OptimizationV2SubmitParamsLocations{
			Location:   []string{"string"},
			ID:         nextbillionai.Int(0),
			Approaches: []string{"unrestricted"},
		},
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
		CostMatrix: [][]int64{{0}},
		Depots: []nextbillionai.OptimizationV2SubmitParamsDepot{{
			ID:            `"id":"depot 1"`,
			LocationIndex: 0,
			Description:   nextbillionai.String("“description”:”Los_Angeles_depot”"),
			Service:       nextbillionai.Int(0),
			TimeWindows:   [][]int64{{0}},
		}},
		Description:        nextbillionai.String(`"description": "Sample Optimization"`),
		DistanceMatrix:     [][]int64{{0}},
		DurationMatrix:     [][]int64{{0}},
		ExistingSolutionID: nextbillionai.String("existing_solution_id"),
		Jobs: []nextbillionai.JobParam{{
			ID:                    `"id":"Job 1"`,
			LocationIndex:         0,
			Delivery:              []int64{0},
			DepotIDs:              []string{"string"},
			Description:           nextbillionai.String("description"),
			FollowLifoOrder:       nextbillionai.Bool(true),
			IncompatibleLoadTypes: []string{"string"},
			JointOrder:            nextbillionai.Int(0),
			LoadTypes:             []string{"string"},
			MaxVisitLateness:      nextbillionai.Int(0),
			Metadata:              "{\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
			OutsourcingCost:       nextbillionai.Int(0),
			Pickup:                []int64{0},
			Priority:              nextbillionai.Int(0),
			Revenue:               nextbillionai.Int(0),
			SequenceOrder:         nextbillionai.Int(0),
			Service:               nextbillionai.Int(0),
			Setup:                 nextbillionai.Int(0),
			Skills:                []int64{1},
			TimeWindows:           [][]int64{{0}},
			Volume: nextbillionai.JobVolumeParam{
				Alignment: "strict",
				Depth:     nextbillionai.Float(0),
				Height:    nextbillionai.Float(0),
				Width:     nextbillionai.Float(0),
			},
			Zones: []int64{0},
		}},
		Options: nextbillionai.OptimizationV2SubmitParamsOptions{
			Constraint: nextbillionai.OptimizationV2SubmitParamsOptionsConstraint{
				MaxActivityWaitingTime: nextbillionai.Int(0),
				MaxVehicleOvertime:     nextbillionai.Int(0),
				MaxVisitLateness:       nextbillionai.Int(0),
			},
			Grouping: nextbillionai.OptimizationV2SubmitParamsOptionsGrouping{
				OrderGrouping: nextbillionai.OptimizationV2SubmitParamsOptionsGroupingOrderGrouping{
					GroupingDiameter: nextbillionai.Float(0),
				},
				ProximityFactor: nextbillionai.Float(0),
				RouteGrouping: nextbillionai.OptimizationV2SubmitParamsOptionsGroupingRouteGrouping{
					PenaltyFactor: nextbillionai.Float(0),
					ZoneDiameter:  nextbillionai.Float(0),
					ZoneSource:    "system_generated",
				},
			},
			Objective: nextbillionai.OptimizationV2SubmitParamsOptionsObjective{
				AllowEarlyArrival: nextbillionai.Bool(true),
				Custom: nextbillionai.OptimizationV2SubmitParamsOptionsObjectiveCustom{
					Type:  "min",
					Value: "vehicles",
				},
				MinimiseNumDepots: nextbillionai.Bool(true),
				SolverMode:        "flexible",
				SolvingTimeLimit:  nextbillionai.Int(0),
				TravelCost:        "duration",
			},
			Routing: nextbillionai.OptimizationV2SubmitParamsOptionsRouting{
				Allow:            []string{"taxi"},
				Avoid:            []string{"toll"},
				Context:          "avgspeed",
				CrossBorder:      nextbillionai.Bool(true),
				DisableCache:     nextbillionai.Bool(true),
				HazmatType:       []string{"general"},
				Mode:             "car",
				Profiles:         "\"profiles\":{\n    \"mini-van\":{\n        \"mode\": \"car\",\n        \"avoid\":[\"highway, toll\"]\n        },\n    \"trailer\":{\n        \"mode\": \"truck\",\n        \"truck_weight\":12000,\n        \"truck_size\":\"200, 210, 600\",\n        \"hazmat_type\": [\"general\", \"harmful_to_water\"]\n        }\n    }\n",
				TrafficTimestamp: nextbillionai.Int(0),
				TruckAxleLoad:    nextbillionai.Float(0),
				TruckSize:        nextbillionai.String(`"truck_size":"200,210,600"`),
				TruckWeight:      nextbillionai.Int(0),
			},
		},
		Relations: []nextbillionai.OptimizationV2SubmitParamsRelation{{
			Steps: []nextbillionai.OptimizationV2SubmitParamsRelationStep{{
				Type: "start",
				ID:   nextbillionai.String(`"id":"Job 1"`),
			}},
			Type:        "in_same_route",
			ID:          nextbillionai.Int(0),
			MaxDuration: nextbillionai.Int(0),
			MinDuration: nextbillionai.Int(0),
			Vehicle:     nextbillionai.String(`"vehicle": "Vehicle 10"`),
		}},
		Shipments: []nextbillionai.ShipmentParam{{
			Delivery: nextbillionai.ShipmentDeliveryParam{
				ID:               `"id":"Shipment Delivery 1"`,
				LocationIndex:    0,
				Description:      nextbillionai.String("description"),
				MaxVisitLateness: nextbillionai.Int(0),
				Metadata:         "{\n  \"notes\": \"dropoff at the patio\",\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
				SequenceOrder:    nextbillionai.Int(0),
				Service:          nextbillionai.Int(0),
				Setup:            nextbillionai.Int(0),
				TimeWindows:      [][]int64{{0}},
			},
			Pickup: nextbillionai.ShipmentPickupParam{
				ID:               `"id": "Shipment Pickup 1"`,
				LocationIndex:    0,
				Description:      nextbillionai.String("description"),
				MaxVisitLateness: nextbillionai.Int(0),
				Metadata:         "{\n  \"notes\": \"involves fragile items\",\n  \"contact\": \"212-456-7890\",\n  \"metaId\": 1234\n}",
				SequenceOrder:    nextbillionai.Int(0),
				Service:          nextbillionai.Int(0),
				Setup:            nextbillionai.Int(0),
				TimeWindows:      [][]int64{{0}},
			},
			Amount:                []int64{0},
			FollowLifoOrder:       nextbillionai.Bool(true),
			IncompatibleLoadTypes: []string{"string"},
			JointOrder:            nextbillionai.Int(0),
			LoadTypes:             []string{"string"},
			MaxTimeInVehicle:      nextbillionai.Int(0),
			OutsourcingCost:       nextbillionai.Int(0),
			Priority:              nextbillionai.Int(0),
			Revenue:               nextbillionai.Int(0),
			Skills:                []int64{0},
			Volume: nextbillionai.ShipmentVolumeParam{
				Alignment: "strict",
				Depth:     nextbillionai.Float(0),
				Height:    nextbillionai.Float(0),
				Width:     nextbillionai.Float(0),
			},
			Zones: []int64{0},
		}},
		Solution: []nextbillionai.OptimizationV2SubmitParamsSolution{{
			Cost: 0,
			Steps: []nextbillionai.OptimizationV2SubmitParamsSolutionStep{{
				ID:            `"id": "Job 10"`,
				Arrival:       0,
				Type:          "start",
				Description:   nextbillionai.String("description"),
				Distance:      nextbillionai.Int(0),
				Duration:      nextbillionai.Int(0),
				Load:          []int64{0},
				Location:      []float64{0},
				LocationIndex: nextbillionai.Int(0),
				Service:       nextbillionai.Int(0),
				Setup:         nextbillionai.Int(0),
				WaitingTime:   nextbillionai.Int(0),
			}},
			Vehicle:     "vehicle",
			Delivery:    []int64{0},
			Description: nextbillionai.String("description"),
			Distance:    nextbillionai.Int(0),
			Duration:    nextbillionai.Int(0),
			Geometry:    nextbillionai.String(`"geometry": "}ebGgcsxRE?CuDOYCYAG???"`),
			Pickup:      []int64{0},
			Priority:    nextbillionai.Int(0),
			Service:     nextbillionai.Int(0),
			Setup:       nextbillionai.Int(0),
			WaitingTime: nextbillionai.Int(0),
		}},
		Unassigned: nextbillionai.OptimizationV2SubmitParamsUnassigned{
			Jobs:      []string{"string"},
			Shipments: [][]string{{"string"}},
		},
		Zones: []nextbillionai.OptimizationV2SubmitParamsZone{{
			ID:         0,
			GeofenceID: nextbillionai.String("geofence_id"),
			Geometry: nextbillionai.OptimizationV2SubmitParamsZoneGeometry{
				Coordinates: [][]float64{{0}},
				Description: nextbillionai.String("description"),
				Type:        "Polygon",
			},
		}},
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
