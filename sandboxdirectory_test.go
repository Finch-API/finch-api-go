// File generated from our OpenAPI spec by Stainless.

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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Sandbox.Directory.New(context.TODO(), finchgo.SandboxDirectoryNewParams{
		Body: finchgo.F([]finchgo.SandboxDirectoryNewParamsBody{{
			FirstName:     finchgo.F("John"),
			MiddleName:    finchgo.F("string"),
			LastName:      finchgo.F("Smith"),
			PreferredName: finchgo.F("string"),
			Emails: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyEmail{{
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmailsTypeWork),
			}}),
			PhoneNumbers: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyPhoneNumber{{
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyPhoneNumbersTypeWork),
			}}),
			Gender:       finchgo.F(finchgo.SandboxDirectoryNewParamsBodyGenderFemale),
			Ethnicity:    finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEthnicityAsian),
			Dob:          finchgo.F("01/01/2000"),
			Ssn:          finchgo.F("string"),
			EncryptedSsn: finchgo.F("string"),
			Residence: finchgo.F(finchgo.LocationParam{
				Line1:      finchgo.F("string"),
				Line2:      finchgo.F("string"),
				City:       finchgo.F("string"),
				State:      finchgo.F("string"),
				PostalCode: finchgo.F("string"),
				Country:    finchgo.F("string"),
				Name:       finchgo.F("string"),
				SourceID:   finchgo.F("string"),
			}),
			Title: finchgo.F("string"),
			Manager: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyManager{
				ID: finchgo.F("string"),
			}),
			Department: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyDepartment{
				Name: finchgo.F("string"),
			}),
			Employment: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmployment{
				Type:    finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentTypeEmployee),
				Subtype: finchgo.F(finchgo.SandboxDirectoryNewParamsBodyEmploymentSubtypeFullTime),
			}),
			StartDate: finchgo.F("string"),
			EndDate:   finchgo.F("string"),
			IsActive:  finchgo.F(true),
			ClassCode: finchgo.F("string"),
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
			CustomFields: finchgo.F([]finchgo.SandboxDirectoryNewParamsBodyCustomField{{
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}, {
				Name:  finchgo.F("string"),
				Value: finchgo.F[any](map[string]interface{}{}),
			}}),
			SourceID: finchgo.F("string"),
		}}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
