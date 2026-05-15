// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestHRISBenefitIndividualEnrollManyWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Benefits.Individuals.EnrollMany(
		context.TODO(),
		"benefit_id",
		finchgo.HRISBenefitIndividualEnrollManyParams{
			EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
			Individuals: []finchgo.HRISBenefitIndividualEnrollManyParamsIndividual{{
				Configuration: finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfiguration{
					AnnualContributionLimit: finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationAnnualContributionLimitIndividual),
					AnnualMaximum:           finchgo.Null[int64](),
					CatchUp:                 finchgo.F(true),
					CompanyContribution: finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContribution{
						Amount: finchgo.F(int64(0)),
						Tiers: finchgo.F([]finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTier{{
							Match:     finchgo.F(int64(0)),
							Threshold: finchgo.F(int64(0)),
						}}),
						Type: finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationCompanyContributionTypeFixed),
					}),
					EffectiveDate: finchgo.F(time.Now()),
					EmployeeDeduction: finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeduction{
						Amount: finchgo.F(int64(10000)),
						Type:   finchgo.F(finchgo.HRISBenefitIndividualEnrollManyParamsIndividualsConfigurationEmployeeDeductionTypeFixed),
					}),
				}),
				IndividualID: finchgo.F("d02a6346-1f08-4312-a064-49ff3cafaa7a"),
			}},
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

func TestHRISBenefitIndividualEnrolledIDsWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Benefits.Individuals.EnrolledIDs(
		context.TODO(),
		"benefit_id",
		finchgo.HRISBenefitIndividualEnrolledIDsParams{
			EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
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

func TestHRISBenefitIndividualGetManyBenefitsWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Benefits.Individuals.GetManyBenefits(
		context.TODO(),
		"benefit_id",
		finchgo.HRISBenefitIndividualGetManyBenefitsParams{
			EntityIDs:     finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
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

func TestHRISBenefitIndividualUnenrollManyWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Benefits.Individuals.UnenrollMany(
		context.TODO(),
		"benefit_id",
		finchgo.HRISBenefitIndividualUnenrollManyParams{
			EntityIDs:     finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
			IndividualIDs: finchgo.F([]string{"string"}),
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
