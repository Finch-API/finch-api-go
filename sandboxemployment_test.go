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
	)
	_, err := client.Sandbox.Employment.Update(
		context.TODO(),
		"string",
		finchgo.SandboxEmploymentUpdateParams{
			ClassCode: finchgo.F("string"),
			CustomFields: finchgo.F([]finchgo.SandboxEmploymentUpdateParamsCustomField{{
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			Department: finchgo.F(finchgo.SandboxEmploymentUpdateParamsDepartment{
				Name: finchgo.F("string"),
			}),
			Employment: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmployment{
				Type:    finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentTypeEmployee),
				Subtype: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentSubtypeFullTime),
			}),
			EndDate:   finchgo.F("string"),
			FirstName: finchgo.F("string"),
			Income: finchgo.F(finchgo.IncomeParam{
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("string"),
				EffectiveDate: finchgo.F("string"),
			}),
			IncomeHistory: finchgo.F([]finchgo.IncomeParam{{
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("string"),
				EffectiveDate: finchgo.F("string"),
			}, {
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("string"),
				EffectiveDate: finchgo.F("string"),
			}, {
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("string"),
				EffectiveDate: finchgo.F("string"),
			}}),
			IsActive: finchgo.F(true),
			LastName: finchgo.F("string"),
			Location: finchgo.F(finchgo.LocationParam{
				Line1:      finchgo.F("string"),
				Line2:      finchgo.F("string"),
				City:       finchgo.F("string"),
				State:      finchgo.F("string"),
				PostalCode: finchgo.F("string"),
				Country:    finchgo.F("string"),
				Name:       finchgo.F("string"),
				SourceID:   finchgo.F("string"),
			}),
			Manager: finchgo.F(finchgo.SandboxEmploymentUpdateParamsManager{
				ID: finchgo.F("string"),
			}),
			MiddleName: finchgo.F("string"),
			SourceID:   finchgo.F("string"),
			StartDate:  finchgo.F("string"),
			Title:      finchgo.F("string"),
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
