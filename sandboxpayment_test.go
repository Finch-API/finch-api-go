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
	)
	_, err := client.Sandbox.Payment.New(context.TODO(), finchgo.SandboxPaymentNewParams{
		EndDate: finchgo.F("end_date"),
		PayStatements: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatement{{
			Earnings: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEarning{{
				Amount: finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsAttributes{
					Metadata: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsAttributesMetadata{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": "bar",
						}),
					}),
				}),
				Currency: finchgo.F("currency"),
				Hours:    finchgo.F(0.000000),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEarningsTypeSalary),
			}}),
			EmployeeDeductions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeduction{{
				Amount: finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributes{
					Metadata: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployeeDeductionsAttributesMetadata{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": "bar",
						}),
					}),
				}),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				PreTax:   finchgo.F(true),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
			EmployerContributions: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsEmployerContribution{{
				Amount: finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributes{
					Metadata: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsEmployerContributionsAttributesMetadata{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": "bar",
						}),
					}),
				}),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.BenefitType_401k),
			}}),
			GrossPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
			}),
			IndividualID: finchgo.F("individual_id"),
			NetPay: finchgo.F(finchgo.MoneyParam{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
			}),
			PaymentMethod: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsPaymentMethodCheck),
			Taxes: finchgo.F([]finchgo.SandboxPaymentNewParamsPayStatementsTax{{
				Amount: finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesAttributes{
					Metadata: finchgo.F(finchgo.SandboxPaymentNewParamsPayStatementsTaxesAttributesMetadata{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": "bar",
						}),
					}),
				}),
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
