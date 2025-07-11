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
		"individual_id",
		finchgo.SandboxEmploymentUpdateParams{
			ClassCode: finchgo.F("class_code"),
			CustomFields: finchgo.F([]finchgo.SandboxEmploymentUpdateParamsCustomField{{
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			Department: finchgo.F(finchgo.SandboxEmploymentUpdateParamsDepartment{
				Name: finchgo.F("name"),
			}),
			Employment: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmployment{
				Subtype: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentSubtypeFullTime),
				Type:    finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentTypeEmployee),
			}),
			EmploymentStatus: finchgo.F(finchgo.SandboxEmploymentUpdateParamsEmploymentStatusActive),
			EndDate:          finchgo.F("end_date"),
			FirstName:        finchgo.F("first_name"),
			Income: finchgo.F(finchgo.IncomeParam{
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F(time.Now()),
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
			}),
			IncomeHistory: finchgo.F([]finchgo.IncomeParam{{
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F(time.Now()),
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
			}}),
			IsActive:         finchgo.F(true),
			LastName:         finchgo.F("last_name"),
			LatestRehireDate: finchgo.F("latest_rehire_date"),
			Location: finchgo.F(finchgo.LocationParam{
				City:       finchgo.F("city"),
				Country:    finchgo.F("country"),
				Line1:      finchgo.F("line1"),
				Line2:      finchgo.F("line2"),
				PostalCode: finchgo.F("postal_code"),
				State:      finchgo.F("state"),
				Name:       finchgo.F("name"),
				SourceID:   finchgo.F("source_id"),
			}),
			Manager: finchgo.F(finchgo.SandboxEmploymentUpdateParamsManager{
				ID: finchgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			}),
			MiddleName: finchgo.F("middle_name"),
			SourceID:   finchgo.F("source_id"),
			StartDate:  finchgo.F("start_date"),
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
