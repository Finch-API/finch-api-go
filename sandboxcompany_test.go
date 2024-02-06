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

func TestSandboxCompanyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandbox.Company.Update(context.TODO(), finchgo.SandboxCompanyUpdateParams{
		Accounts: finchgo.F([]finchgo.SandboxCompanyUpdateParamsAccount{{
			RoutingNumber:   finchgo.F("string"),
			AccountName:     finchgo.F("string"),
			InstitutionName: finchgo.F("string"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("string"),
		}, {
			RoutingNumber:   finchgo.F("string"),
			AccountName:     finchgo.F("string"),
			InstitutionName: finchgo.F("string"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("string"),
		}, {
			RoutingNumber:   finchgo.F("string"),
			AccountName:     finchgo.F("string"),
			InstitutionName: finchgo.F("string"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("string"),
		}}),
		Departments: finchgo.F([]finchgo.SandboxCompanyUpdateParamsDepartment{{
			Name: finchgo.F("string"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("string"),
			}),
		}, {
			Name: finchgo.F("string"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("string"),
			}),
		}, {
			Name: finchgo.F("string"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("string"),
			}),
		}}),
		Ein: finchgo.F("string"),
		Entity: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntity{
			Type:    finchgo.F(finchgo.SandboxCompanyUpdateParamsEntityTypeLlc),
			Subtype: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntitySubtypeSCorporation),
		}),
		LegalName: finchgo.F("string"),
		Locations: finchgo.F([]finchgo.LocationParam{{
			Line1:      finchgo.F("string"),
			Line2:      finchgo.F("string"),
			City:       finchgo.F("string"),
			State:      finchgo.F("string"),
			PostalCode: finchgo.F("string"),
			Country:    finchgo.F("string"),
			Name:       finchgo.F("string"),
			SourceID:   finchgo.F("string"),
		}, {
			Line1:      finchgo.F("string"),
			Line2:      finchgo.F("string"),
			City:       finchgo.F("string"),
			State:      finchgo.F("string"),
			PostalCode: finchgo.F("string"),
			Country:    finchgo.F("string"),
			Name:       finchgo.F("string"),
			SourceID:   finchgo.F("string"),
		}, {
			Line1:      finchgo.F("string"),
			Line2:      finchgo.F("string"),
			City:       finchgo.F("string"),
			State:      finchgo.F("string"),
			PostalCode: finchgo.F("string"),
			Country:    finchgo.F("string"),
			Name:       finchgo.F("string"),
			SourceID:   finchgo.F("string"),
		}}),
		PrimaryEmail:       finchgo.F("string"),
		PrimaryPhoneNumber: finchgo.F("string"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
