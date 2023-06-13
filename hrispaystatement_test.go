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

func TestHRISPayStatementGetMany(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.HRIS.PayStatements.GetMany(context.TODO(), finchgo.HRISPayStatementGetManyParams{
		Requests: finchgo.F([]finchgo.HRISPayStatementGetManyParamsRequests{{PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"), Limit: finchgo.F(int64(0)), Offset: finchgo.F(int64(0))}, {PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"), Limit: finchgo.F(int64(0)), Offset: finchgo.F(int64(0))}, {PaymentID: finchgo.F("e8b90071-0c11-471c-86e8-e303ef2f6782"), Limit: finchgo.F(int64(0)), Offset: finchgo.F(int64(0))}}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
