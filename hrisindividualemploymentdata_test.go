// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISIndividualEmploymentDataGetMany(t *testing.T) {
	c := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.HRIS.Individuals.EmploymentData.GetMany(context.TODO(), finchgo.HRISIndividualEmploymentDataGetManyParams{
		Requests: finchgo.F([]finchgo.HRISIndividualEmploymentDataGetManyParamsRequests{{IndividualID: finchgo.F("string")}, {IndividualID: finchgo.F("string")}, {IndividualID: finchgo.F("string")}}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
