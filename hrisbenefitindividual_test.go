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

func TestHRISBenefitIndividualEnrolledIDs(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.HRIS.Benefits.Individuals.EnrolledIDs(context.TODO(), "string")
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISBenefitIndividualGetManyBenefitsWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.HRIS.Benefits.Individuals.GetManyBenefits(
		context.TODO(),
		"string",
		finchgo.HRISBenefitIndividualGetManyBenefitsParams{
			IndividualIDs: finchgo.F("d675d2b7-6d7b-41a8-b2d3-001eb3fb88f6,d02a6346-1f08-4312-a064-49ff3cafaa7a"),
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

func TestHRISBenefitIndividualUnenrollWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.HRIS.Benefits.Individuals.Unenroll(
		context.TODO(),
		"string",
		finchgo.HRISBenefitIndividualUnenrollParams{
			IndividualIDs: finchgo.F([]string{"string", "string", "string"}),
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
