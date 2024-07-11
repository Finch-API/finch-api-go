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
		option.WithClientID("4ab15e51-11ad-49f4-acae-f343b7794375"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Sandbox.Company.Update(context.TODO(), finchgo.SandboxCompanyUpdateParams{
		Accounts: finchgo.F([]finchgo.SandboxCompanyUpdateParamsAccount{{
			RoutingNumber:   finchgo.F("routing_number"),
			AccountName:     finchgo.F("account_name"),
			InstitutionName: finchgo.F("institution_name"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("account_number"),
		}, {
			RoutingNumber:   finchgo.F("routing_number"),
			AccountName:     finchgo.F("account_name"),
			InstitutionName: finchgo.F("institution_name"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("account_number"),
		}, {
			RoutingNumber:   finchgo.F("routing_number"),
			AccountName:     finchgo.F("account_name"),
			InstitutionName: finchgo.F("institution_name"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			AccountNumber:   finchgo.F("account_number"),
		}}),
		Departments: finchgo.F([]finchgo.SandboxCompanyUpdateParamsDepartment{{
			Name: finchgo.F("name"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("name"),
			}),
		}, {
			Name: finchgo.F("name"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("name"),
			}),
		}, {
			Name: finchgo.F("name"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("name"),
			}),
		}}),
		Ein: finchgo.F("ein"),
		Entity: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntity{
			Type:    finchgo.F(finchgo.SandboxCompanyUpdateParamsEntityTypeLlc),
			Subtype: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntitySubtypeSCorporation),
		}),
		LegalName: finchgo.F("legal_name"),
		Locations: finchgo.F([]finchgo.LocationParam{{
			Line1:      finchgo.F("line1"),
			Line2:      finchgo.F("line2"),
			City:       finchgo.F("city"),
			State:      finchgo.F("state"),
			PostalCode: finchgo.F("postal_code"),
			Country:    finchgo.F("country"),
			Name:       finchgo.F("name"),
			SourceID:   finchgo.F("source_id"),
		}, {
			Line1:      finchgo.F("line1"),
			Line2:      finchgo.F("line2"),
			City:       finchgo.F("city"),
			State:      finchgo.F("state"),
			PostalCode: finchgo.F("postal_code"),
			Country:    finchgo.F("country"),
			Name:       finchgo.F("name"),
			SourceID:   finchgo.F("source_id"),
		}, {
			Line1:      finchgo.F("line1"),
			Line2:      finchgo.F("line2"),
			City:       finchgo.F("city"),
			State:      finchgo.F("state"),
			PostalCode: finchgo.F("postal_code"),
			Country:    finchgo.F("country"),
			Name:       finchgo.F("name"),
			SourceID:   finchgo.F("source_id"),
		}}),
		PrimaryEmail:       finchgo.F("primary_email"),
		PrimaryPhoneNumber: finchgo.F("primary_phone_number"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
