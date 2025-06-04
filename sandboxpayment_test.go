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
		PayStatements: finchgo.F([]finchgo.PayStatementParam{{
			Earnings: finchgo.F([]finchgo.PayStatementEarningParam{{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Hours:    finchgo.F(0.000000),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.PayStatementEarningsTypeSalary),
				Attributes: finchgo.F(finchgo.PayStatementEarningsAttributesParam{
					Metadata: finchgo.F(finchgo.PayStatementEarningsAttributesMetadataParam{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": map[string]interface{}{},
						}),
					}),
				}),
			}}),
			EmployeeDeductions: finchgo.F([]finchgo.PayStatementEmployeeDeductionParam{{
				Amount:   finchgo.F(int64(0)),
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				PreTax:   finchgo.F(true),
				Type:     finchgo.F(finchgo.BenefitType_457),
				Attributes: finchgo.F(finchgo.PayStatementEmployeeDeductionsAttributesParam{
					Metadata: finchgo.F(finchgo.PayStatementEmployeeDeductionsAttributesMetadataParam{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": map[string]interface{}{},
						}),
					}),
				}),
			}}),
			EmployerContributions: finchgo.F([]finchgo.PayStatementEmployerContributionParam{{
				Currency: finchgo.F("currency"),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.BenefitType_457),
				Amount:   finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.PayStatementEmployerContributionsAttributesParam{
					Metadata: finchgo.F(finchgo.PayStatementEmployerContributionsAttributesMetadataParam{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": map[string]interface{}{},
						}),
					}),
				}),
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
			PaymentMethod: finchgo.F(finchgo.PayStatementPaymentMethodCheck),
			Taxes: finchgo.F([]finchgo.PayStatementTaxParam{{
				Currency: finchgo.F("currency"),
				Employer: finchgo.F(true),
				Name:     finchgo.F("name"),
				Type:     finchgo.F(finchgo.PayStatementTaxesTypeState),
				Amount:   finchgo.F(int64(0)),
				Attributes: finchgo.F(finchgo.PayStatementTaxesAttributesParam{
					Metadata: finchgo.F(finchgo.PayStatementTaxesAttributesMetadataParam{
						Metadata: finchgo.F(map[string]interface{}{
							"foo": map[string]interface{}{},
						}),
					}),
				}),
			}}),
			TotalHours: finchgo.F(0.000000),
			Type:       finchgo.F(finchgo.PayStatementTypeOffCyclePayroll),
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
