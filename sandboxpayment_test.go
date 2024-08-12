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
		option.WithClientID("4ab15e51-11ad-49f4-acae-f343b7794375"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Sandbox.Payment.New(context.TODO(), finchgo.SandboxPaymentNewParams{
		EndDate: finchgo.F("end_date"),
		PayStatements: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatement{{
			Earnings: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEarning{{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Hours:    finchgo.F(0.000000),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Hours:    finchgo.F(0.000000),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Hours:    finchgo.F(0.000000),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
			}}),
			EmployeeDeductions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeduction{{
				Amount:   finchgo.F(int64(2000)),
				Currency: finchgo.F("usd"),
				Name:     finchgo.F("401k test"),
				PreTax:   finchgo.F(true),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
			EmployerContributions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployerContribution{{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
			GrossPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
			}),
			IndividualID: finchgo.F("b2338cfb-472f-4f72-9faa-e028c083144a"),
			NetPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
			}),
			PaymentMethod: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsPaymentMethodCheck),
			Taxes: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsTax{{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Employer: finchgo.F(true),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Employer: finchgo.F(true),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
			}, {
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Employer: finchgo.F(true),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesTypeState),
			}}),
			TotalHours: finchgo.F(0.000000),
			Type:       finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTypeRegularPayroll),
		}}),
		StartDate: finchgo.F("start_date"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
