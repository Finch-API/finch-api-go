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

func TestHRISBenefitNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Benefits.New(context.TODO(), finchgo.HRISBenefitNewParams{
		Description: finchgo.F("string"),
		Frequency:   finchgo.F(finchgo.BenefitFrequencyOneTime),
		Type:        finchgo.F(finchgo.BenefitType401k),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISBenefitGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Benefits.Get(context.TODO(), "string")
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISBenefitUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Benefits.Update(
		context.TODO(),
		"string",
		finchgo.HRISBenefitUpdateParams{
			Description: finchgo.F("string"),
		},
	)
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISBenefitList(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Benefits.List(context.TODO())
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISBenefitListSupportedBenefits(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.HRIS.Benefits.ListSupportedBenefits(context.TODO())
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
