// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"testing"
	"time"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISPaymentList(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	c := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.HRIS.Payments.List(context.TODO(), finchgo.HRISPaymentListParams{
		EndDate:   finchgo.F(time.Now()),
		StartDate: finchgo.F(time.Now()),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
