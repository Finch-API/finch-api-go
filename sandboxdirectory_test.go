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

func TestSandboxDirectoryNew(t *testing.T) {
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
	_, err := client.Sandbox.Directory.New(context.TODO(), finchgo.SandboxDirectoryNewParams{
		Body: []finchgo.SandboxDirectoryNewParamsBody{{
			ClassCode: finchgo.F("class_code"),
			CustomFields: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyCustomField{{
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			Department: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyDepartment{
				Name: finchgo.F("name"),
			}),
			Dob: finchgo.F("01/01/2000"),
			Emails: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyEmail{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}}),
			Employment: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmployment{
				Subtype: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentSubtypeFullTime),
				Type:    finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentTypeEmployee),
			}),
			EmploymentStatus: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentStatusActive),
			EncryptedSsn:     finchgo.F("encrypted_ssn"),
			EndDate:          finchgo.F("end_date"),
			Ethnicity:        finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEthnicityAsian),
			FirstName:        finchgo.F("John"),
			Gender:           finchgo.F(finchgo.SandboxDirectoryNewParamsBodyGenderFemale),
			Income: finchgo.F(finchgo.IncomeParam{
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
			}),
			IncomeHistory: finchgo.F([]finchgo.IncomeParam{{
				Amount:        finchgo.F(int64(0)),
				Currency:      finchgo.F("currency"),
				EffectiveDate: finchgo.F("effective_date"),
				Unit:          finchgo.F(finchgo.IncomeUnitYearly),
			}}),
			IsActive:         finchgo.F(true),
			LastName:         finchgo.F("Smith"),
			LatestRehireDate: finchgo.F("latest_rehire_date"),
			Location: finchgo.F(finchgo.LocationParam{
				City:       finchgo.F("city"),
				Country:    finchgo.F("country"),
				Line1:      finchgo.F("line1"),
				Line2:      finchgo.F("line2"),
				Name:       finchgo.F("name"),
				PostalCode: finchgo.F("postal_code"),
				SourceID:   finchgo.F("source_id"),
				State:      finchgo.F("state"),
			}),
			Manager: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyManager{
				ID: finchgo.F("id"),
			}),
			MiddleName: finchgo.F("middle_name"),
			PhoneNumbers: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyPhoneNumber{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}}),
			PreferredName: finchgo.F("preferred_name"),
			Residence: finchgo.F(finchgo.LocationParam{
				City:       finchgo.F("city"),
				Country:    finchgo.F("country"),
				Line1:      finchgo.F("line1"),
				Line2:      finchgo.F("line2"),
				Name:       finchgo.F("name"),
				PostalCode: finchgo.F("postal_code"),
				SourceID:   finchgo.F("source_id"),
				State:      finchgo.F("state"),
			}),
			SourceID:  finchgo.F("source_id"),
			Ssn:       finchgo.F("ssn"),
			StartDate: finchgo.F("start_date"),
			Title:     finchgo.F("title"),
		}},
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
