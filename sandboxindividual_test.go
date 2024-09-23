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
	)
	_, err := client.Sandbox.Individual.Update(
		context.TODO(),
		"individual_id",
		finchgo.SandboxIndividualUpdateParams{
			Dob: finchgo.F("12/20/1989"),
			Emails: finchgo.F([]finchgo.SandboxIndividualUpdateParamsEmail{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsEmailsTypeWork),
			}}),
			EncryptedSsn: finchgo.F("encrypted_ssn"),
			Ethnicity:    finchgo.F(finchgo.SandboxIndividualUpdateParamsEthnicityAsian),
			FirstName:    finchgo.F("first_name"),
			Gender:       finchgo.F(finchgo.SandboxIndividualUpdateParamsGenderFemale),
			LastName:     finchgo.F("last_name"),
			MiddleName:   finchgo.F("middle_name"),
			PhoneNumbers: finchgo.F([]finchgo.SandboxIndividualUpdateParamsPhoneNumber{{
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
			}, {
				Data: finchgo.F("data"),
				Type: finchgo.F(finchgo.SandboxIndividualUpdateParamsPhoneNumbersTypeWork),
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
			Ssn: finchgo.F("ssn"),
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
