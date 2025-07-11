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

func TestSandboxDirectoryNewWithOptionalParams(t *testing.T) {
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
			Dob: finchgo.F("dob"),
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
			FirstName:        finchgo.F("first_name"),
			Gender:           finchgo.F(finchgo.SandboxDirectoryNewParamsBodyGenderFemale),
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
			Manager: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyManager{
				ID: finchgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
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
				PostalCode: finchgo.F("postal_code"),
				State:      finchgo.F("state"),
				Name:       finchgo.F("name"),
				SourceID:   finchgo.F("source_id"),
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
