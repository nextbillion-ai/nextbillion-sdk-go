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

func TestSnapToRoadSnapWithOptionalParams(t *testing.T) {
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
	_, err := client.SnapToRoads.Snap(context.TODO(), nextbillionai.SnapToRoadSnapParams{
		Key:             "key=API_KEY",
		Path:            "path=41.38602272,2.17621539|41.38312885,2.17207083|41.38157854,2.17906668|41.38288511,2.18186215",
		Approaches:      nextbillionai.SnapToRoadSnapParamsApproachesUnrestricted,
		Avoid:           nextbillionai.SnapToRoadSnapParamsAvoidToll,
		Geometry:        nextbillionai.SnapToRoadSnapParamsGeometryPolyline,
		Mode:            nextbillionai.SnapToRoadSnapParamsModeCar,
		Option:          nextbillionai.SnapToRoadSnapParamsOptionFlexible,
		Radiuses:        nextbillionai.String("radiuses=14|16|14"),
		RoadInfo:        nextbillionai.SnapToRoadSnapParamsRoadInfoMaxSpeed,
		Timestamps:      nextbillionai.String("timestamps=1656570000|1656570015|1656570030"),
		TolerateOutlier: nextbillionai.Bool(true),
	})
	if err != nil {
		var apierr *nextbillionai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
