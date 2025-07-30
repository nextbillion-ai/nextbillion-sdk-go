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

func TestOptimizationComputeWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Optimization.Compute(context.TODO(), nextbillionai.OptimizationComputeParams{
		Coordinates:  "coordinates=41.35544869444527,2.0747669962025292|41.37498154684205,2.103705 4530396886|41.38772862000152,2.1311887061315526",
		Key:          "key=API_KEY",
		Approaches:   nextbillionai.OptimizationComputeParamsApproachesUnrestricted,
		Destination:  nextbillionai.OptimizationComputeParamsDestinationAny,
		Geometries:   nextbillionai.OptimizationComputeParamsGeometriesPolyline,
		Mode:         nextbillionai.OptimizationComputeParamsModeCar,
		Roundtrip:    nextbillionai.Bool(true),
		Source:       nextbillionai.OptimizationComputeParamsSourceAny,
		WithGeometry: nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
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
	client := nextbillionai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Optimization.ReOptimize(context.TODO(), nextbillionai.OptimizationReOptimizeParams{
		Key:               "key=API_KEY",
		ExistingRequestID: "existing_request_id",
		JobChanges: nextbillionai.OptimizationReOptimizeParamsJobChanges{
			Add: []nextbillionai.JobParam{{
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
					Alignment: "`strict`",
					Depth:     nextbillionai.Float(0),
					Height:    nextbillionai.Float(0),
					Width:     nextbillionai.Float(0),
				},
				Zones: []int64{0},
			}},
			Modify: []nextbillionai.JobParam{{
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
					Alignment: "`strict`",
					Depth:     nextbillionai.Float(0),
					Height:    nextbillionai.Float(0),
					Width:     nextbillionai.Float(0),
				},
				Zones: []int64{0},
			}},
			Remove: []string{"string"},
		},
		Locations: []string{"string"},
		ShipmentChanges: nextbillionai.OptimizationReOptimizeParamsShipmentChanges{
			Add: []nextbillionai.ShipmentParam{{
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
					Alignment: "`strict`",
					Depth:     nextbillionai.Float(0),
					Height:    nextbillionai.Float(0),
					Width:     nextbillionai.Float(0),
				},
				Zones: []int64{0},
			}},
			Modify: []nextbillionai.ShipmentParam{{
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
					Alignment: "`strict`",
					Depth:     nextbillionai.Float(0),
					Height:    nextbillionai.Float(0),
					Width:     nextbillionai.Float(0),
				},
				Zones: []int64{0},
			}},
			Remove: []string{"string"},
		},
		VehicleChanges: nextbillionai.OptimizationReOptimizeParamsVehicleChanges{
			Add: []nextbillionai.VehicleParam{{
				ID: "id",
				Location: nextbillionai.VehicleLocationParam{
					Lat: nextbillionai.Float(0),
					Lng: nextbillionai.Float(0),
				},
				Attributes: "\"attributes\":{\n    \"driver_rating\": \"4.0\",\n    \"trip_types\": \"premium\"\n  }",
				Priority:   nextbillionai.Int(0),
				RemainingWaypoints: []nextbillionai.LocationParam{{
					Lat: -90,
					Lon: -180,
				}},
			}},
			Modify: nextbillionai.VehicleParam{
				ID: "id",
				Location: nextbillionai.VehicleLocationParam{
					Lat: nextbillionai.Float(0),
					Lng: nextbillionai.Float(0),
				},
				Attributes: "\"attributes\":{\n    \"driver_rating\": \"4.0\",\n    \"trip_types\": \"premium\"\n  }",
				Priority:   nextbillionai.Int(0),
				RemainingWaypoints: []nextbillionai.LocationParam{{
					Lat: -90,
					Lon: -180,
				}},
			},
			Remove: []string{"string"},
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
