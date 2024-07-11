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
		option.WithClientID("4ab15e51-11ad-49f4-acae-f343b7794375"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Sandbox.Directory.New(context.TODO(), finchgo.SandboxDirectoryNewParams{
		Body: []finchgo.SandboxDirectoryNewParamsBody{{
			FirstName:     finchgo.F("John"),
			MiddleName:    finchgo.F("middle_name"),
			LastName:      finchgo.F("Smith"),
			PreferredName: finchgo.F("preferred_name"),
			Emails: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyEmail{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}}),
			PhoneNumbers: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyPhoneNumber{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}}),
			Gender:       finchgo.F(finchgo.SandboxDirectoryNewParamsBodyGenderFemale),
			Ethnicity:    finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEthnicityAsian),
			Dob:          finchgo.F("01/01/2000"),
			Ssn:          finchgo.F("ssn"),
			EncryptedSsn: finchgo.F("encrypted_ssn"),
			Residence: finchgo.F(finchgo.LocationParam{
				Line1:      finchgo.F("line1"),
				Line2:      finchgo.F("line2"),
				City:       finchgo.F("city"),
				State:      finchgo.F("state"),
				PostalCode: finchgo.F("postal_code"),
				Country:    finchgo.F("country"),
				Name:       finchgo.F("name"),
				SourceID:   finchgo.F("source_id"),
			}),
			Title: finchgo.F("title"),
			Manager: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyManager{
				ID: finchgo.F("id"),
			}),
			Department: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyDepartment{
				Name: finchgo.F("name"),
			}),
			Employment: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmployment{
				Type:    finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentTypeEmployee),
				Subtype: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentSubtypeFullTime),
			}),
			StartDate: finchgo.F("start_date"),
			EndDate:   finchgo.F("end_date"),
			IsActive:  finchgo.F(true),
			ClassCode: finchgo.F("class_code"),
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
			CustomFields: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyCustomField{{
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("name"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			SourceID: finchgo.F("source_id"),
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
