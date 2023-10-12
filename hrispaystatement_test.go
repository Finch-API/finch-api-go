// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISPayStatementGetMany(t *testing.T) {
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
	)
	_, err := client.HRIS.PayStatements.GetMany(context.TODO(), finchgo.HRISPayStatementGetManyParams{
		Requests: finchgo.F([]finchgo.HRISPayStatementGetManyParamsRequest{{
			PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"),
			Limit:     finchgo.F(int64(0)),
			Offset:    finchgo.F(int64(0)),
		}, {
			PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"),
			Limit:     finchgo.F(int64(0)),
			Offset:    finchgo.F(int64(0)),
		}, {
			PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"),
			Limit:     finchgo.F(int64(0)),
			Offset:    finchgo.F(int64(0)),
		}}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
