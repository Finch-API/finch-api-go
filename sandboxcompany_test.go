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
	)
	_, err := client.Sandbox.Company.Update(context.TODO(), finchgo.SandboxCompanyUpdateParams{
		Accounts: finchgo.F([]finchgo.SandboxCompanyUpdateParamsAccount{{
			AccountName:     finchgo.F("account_name"),
			AccountNumber:   finchgo.F("account_number"),
			AccountType:     finchgo.F(finchgo.SandboxCompanyUpdateParamsAccountsAccountTypeChecking),
			InstitutionName: finchgo.F("institution_name"),
			RoutingNumber:   finchgo.F("routing_number"),
		}}),
		Departments: finchgo.F([]finchgo.SandboxCompanyUpdateParamsDepartment{{
			Name: finchgo.F("name"),
			Parent: finchgo.F(finchgo.SandboxCompanyUpdateParamsDepartmentsParent{
				Name: finchgo.F("name"),
			}),
		}}),
		Ein: finchgo.F("ein"),
		Entity: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntity{
			Subtype: finchgo.F(finchgo.SandboxCompanyUpdateParamsEntitySubtypeSCorporation),
			Type:    finchgo.F(finchgo.SandboxCompanyUpdateParamsEntityTypeLlc),
		}),
		LegalName: finchgo.F("legal_name"),
		Locations: finchgo.F([]finchgo.LocationParam{{
			City:       finchgo.F("city"),
			Country:    finchgo.F("country"),
			Line1:      finchgo.F("line1"),
			Line2:      finchgo.F("line2"),
			PostalCode: finchgo.F("postal_code"),
			State:      finchgo.F("state"),
			Name:       finchgo.F("name"),
			SourceID:   finchgo.F("source_id"),
		}}),
		PrimaryEmail:       finchgo.F("dev@stainless.com"),
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
