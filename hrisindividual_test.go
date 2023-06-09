// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISIndividualGetManyWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.HRIS.Individuals.GetMany(context.TODO(), finchgo.HRISIndividualGetManyParams{
		Options:  finchgo.F(finchgo.HRISIndividualGetManyParamsOptions{Include: finchgo.F([]string{"string", "string", "string"})}),
		Requests: finchgo.F([]finchgo.HRISIndividualGetManyParamsRequests{{IndividualID: finchgo.F("string")}, {IndividualID: finchgo.F("string")}, {IndividualID: finchgo.F("string")}}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
