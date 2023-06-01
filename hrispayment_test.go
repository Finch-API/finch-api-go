package finchgo_test

import (
	"context"
	"errors"
	"testing"
	"time"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISPaymentList(t *testing.T) {
	c := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.HRIS.Payments.List(context.TODO(), finchgo.HRISPaymentListParams{StartDate: finchgo.F(time.Now()), EndDate: finchgo.F(time.Now())})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			println(apierr.DumpRequest(true))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
