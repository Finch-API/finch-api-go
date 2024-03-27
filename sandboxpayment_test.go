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
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Sandbox.Payment.New(context.TODO(), finchgo.SandboxPaymentNewParams{
		EndDate: finchgo.F("string"),
		PayStatements: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatement{{
			IndividualID:  finchgo.F("b2338cfb-472f-4f72-9faa-e028c083144a"),
			Type:          finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTypeRegularPayroll),
			PaymentMethod: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsPaymentMethodCheck),
			TotalHours:    finchgo.F(0.000000),
			GrossPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
			}),
			NetPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
			}),
			Earnings: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEarning{{
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Hours:    finchgo.F(0.000000),
			}, {
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Hours:    finchgo.F(0.000000),
			}, {
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Hours:    finchgo.F(0.000000),
			}}),
			Taxes: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsTax{{
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
				Name:     finchgo.F("string"),
				Employer: finchgo.F(true),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
			}, {
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
				Name:     finchgo.F("string"),
				Employer: finchgo.F(true),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
			}, {
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
				Name:     finchgo.F("string"),
				Employer: finchgo.F(true),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
			}}),
			EmployeeDeductions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeduction{{
				Name:     finchgo.F("401k test"),
				Amount:   finchgo.F(int64(2000)),
				Currency: finchgo.F("usd"),
				PreTax:   finchgo.F(true),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
			EmployerContributions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployerContribution{{
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}, {
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}, {
				Name:     finchgo.F("string"),
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("string"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
		}}),
		StartDate: finchgo.F("string"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
