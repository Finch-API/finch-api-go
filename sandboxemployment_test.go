// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestSandboxEmploymentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandbox.Employment.Update(
		context.TODO(),
		"individual_id",
		finchgo.SandboxEmploymentUpdateParams{
			ClassCode: finchgo.F("class_code"),
			CustomFields: finchgo.F([]finchgo.SandboxEmploymentUpdateParamsCustomField{{
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			Department: finchgo.F(finchgo.SandboxEmploymentUpdateParamsDepartment{
				Name: finchgo.F("name"),
			}),
			Employment: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmployment{
				Type:    finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentTypeEmployee),
				Subtype: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentSubtypeFullTime),
			}),
			EndDate:   finchgo.F("end_date"),
			FirstName: finchgo.F("first_name"),
			Income: finchgo.F(finchgo.IncomeParam{
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
			}),
			IncomeHistory: finchgo.F([]finchgo.IncomeParam{{
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
			}, {
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
			}, {
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
			}}),
			IsActive: finchgo.F(true),
			LastName: finchgo.F("last_name"),
			Location: finchgo.F(finchgo.LocationParam{
				Line1:      finchgo.F("line1"),
				Line2:      finchgo.F("line2"),
				City:       finchgo.F("city"),
				State:      finchgo.F("state"),
				PostalCode: finchgo.F("postal_code"),
				Country:    finchgo.F("country"),
				Name:       finchgo.F("name"),
				SourceID:   finchgo.F("source_id"),
			}),
			Manager: finchgo.F(finchgo.SandboxEmploymentUpdateParamsManager{
				ID: finchgo.F("id"),
			}),
			MiddleName: finchgo.F("middle_name"),
			SourceID:   finchgo.F("source_id"),
			StartDate:  finchgo.F("3/4/2020"),
			Title:      finchgo.F("title"),
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
