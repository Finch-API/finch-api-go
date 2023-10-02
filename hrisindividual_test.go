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

func TestHRISIndividualGetManyWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Individuals.GetMany(context.TODO(), finchgo.HRISIndividualGetManyParams{
		Options: finchgo.F(finchgo.HRISIndividualGetManyParamsOptions{
			Include: finchgo.F([]string{"string", "string", "string"}),
		}),
		Requests: finchgo.F([]finchgo.HRISIndividualGetManyParamsRequest{{
			IndividualID: finchgo.F("string"),
		}, {
			IndividualID: finchgo.F("string"),
		}, {
			IndividualID: finchgo.F("string"),
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
