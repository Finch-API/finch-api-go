// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Finch-API/finch-api-go/v2"
	"github.com/Finch-API/finch-api-go/v2/internal/testutil"
	"github.com/Finch-API/finch-api-go/v2/option"
)

func TestHRISPayStatementGetManyWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("My Access Token"),
		option.WithClientID("4ab15e51-11ad-49f4-acae-f343b7794375"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.HRIS.PayStatements.GetMany(context.TODO(), finchgo.HRISPayStatementGetManyParams{
		Requests: finchgo.F([]finchgo.HRISPayStatementGetManyParamsRequest{{
			PaymentID: finchgo.F("fc8b024e-d373-4c9c-80fc-f1625383d142"),
			Limit:     finchgo.F(int64(100)),
			Offset:    finchgo.F(int64(0)),
		}}),
		EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
