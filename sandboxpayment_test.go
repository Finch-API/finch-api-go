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

func TestSandboxPaymentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandbox.Payment.New(context.TODO(), finchgo.SandboxPaymentNewParams{
		EndDate: finchgo.F(time.Now()),
		PayStatements: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatement{{
			IndividualID: finchgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
			Earnings: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEarning{{
				Amount: finchgo.F(int64(0)),
				Hours:  finchgo.F(0.000000),
				Name:   finchgo.F("x"),
				Type:   finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeBonus),
			}}),
			EmployeeDeductions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeduction{{
				Amount: finchgo.F(int64(0)),
				Name:   finchgo.F("x"),
				PreTax: finchgo.F(true),
				Type:   finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeductionsType_457),
			}}),
			EmployerContributions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployerContribution{{
				Amount: finchgo.F(int64(0)),
				Name:   finchgo.F("x"),
				Type:   finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployerContributionsType_457),
			}}),
			GrossPay:      finchgo.F(int64(1)),
			NetPay:        finchgo.F(int64(9007199254740991)),
			PaymentMethod: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsPaymentMethodCheck),
			Taxes: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsTax{{
				Amount:   finchgo.F(int64(0)),
				Employer: finchgo.F(true),
				Name:     finchgo.F("x"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeFederal),
			}}),
			TotalHours: finchgo.F(1.000000),
			Type:       finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTypeOffCyclePayroll),
		}}),
		StartDate: finchgo.F(time.Now()),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
