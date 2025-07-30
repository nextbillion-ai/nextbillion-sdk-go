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

func TestOptimizationComputeWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.Compute(context.TODO(), nextbillionsdk.OptimizationComputeParams{
		Coordinates:  "coordinates=41.35544869444527,2.0747669962025292|41.37498154684205,2.103705 4530396886|41.38772862000152,2.1311887061315526",
		Key:          "key=API_KEY",
		Approaches:   nextbillionsdk.OptimizationComputeParamsApproachesUnrestricted,
		Destination:  nextbillionsdk.OptimizationComputeParamsDestinationAny,
		Geometries:   nextbillionsdk.OptimizationComputeParamsGeometriesPolyline,
		Mode:         nextbillionsdk.OptimizationComputeParamsModeCar,
		Roundtrip:    nextbillionsdk.Bool(true),
		Source:       nextbillionsdk.OptimizationComputeParamsSourceAny,
		WithGeometry: nextbillionsdk.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOptimizationReOptimizeWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimization.ReOptimize(context.TODO(), nextbillionsdk.OptimizationReOptimizeParams{
		Key:               "key=API_KEY",
		ExistingRequestID: "existing_request_id",
		JobChanges: nextbillionsdk.OptimizationReOptimizeParamsJobChanges{
			Add: []nextbillionsdk.JobParam{{
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
			Modify: []nextbillionsdk.JobParam{{
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
			Remove: []string{"string"},
		},
		Locations: []string{"string"},
		ShipmentChanges: nextbillionsdk.OptimizationReOptimizeParamsShipmentChanges{
			Add: []nextbillionsdk.ShipmentParam{{
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
			Modify: []nextbillionsdk.ShipmentParam{{
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
			Remove: []string{"string"},
		},
		VehicleChanges: nextbillionsdk.OptimizationReOptimizeParamsVehicleChanges{
			Add: []nextbillionsdk.VehicleParam{{
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
			Modify: nextbillionsdk.VehicleParam{
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
			},
			Remove: []string{"string"},
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
