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

func TestSandboxIndividualUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandbox.Individual.Update(
		context.TODO(),
		"string",
		finchgo.SandboxIndividualUpdateParams{
			Dob: finchgo.F("12/20/1989"),
			Emails: finchgo.F([]finchgo.SandboxIndividualUpdateParamsEmail{{
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}}),
			EncryptedSsn: finchgo.F("string"),
			Ethnicity:    finchgo.F(finchgo.SandboxIndividualUpdateParamsEthnicityAsian),
			FirstName:    finchgo.F("string"),
			Gender:       finchgo.F(finchgo.SandboxIndividualUpdateParamsGenderFemale),
			LastName:     finchgo.F("string"),
			MiddleName:   finchgo.F("string"),
			PhoneNumbers: finchgo.F([]finchgo.SandboxIndividualUpdateParamsPhoneNumber{{
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("string"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
			}}),
			PreferredName: finchgo.F("string"),
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
			Ssn: finchgo.F("string"),
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
