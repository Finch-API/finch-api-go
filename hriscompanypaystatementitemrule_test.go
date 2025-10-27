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

func TestHRISCompanyPayStatementItemRuleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Company.PayStatementItem.Rules.New(context.TODO(), finchgo.HRISCompanyPayStatementItemRuleNewParams{
		EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
		Attributes: finchgo.F(finchgo.HRISCompanyPayStatementItemRuleNewParamsAttributes{
			Metadata: finchgo.F(map[string]interface{}{
				"foo": "bar",
			}),
		}),
		Conditions: finchgo.F([]finchgo.HRISCompanyPayStatementItemRuleNewParamsCondition{{
			Field:    finchgo.F("field"),
			Operator: finchgo.F(finchgo.HRISCompanyPayStatementItemRuleNewParamsConditionsOperatorEquals),
			Value:    finchgo.F("value"),
		}}),
		EffectiveEndDate:   finchgo.F("effective_end_date"),
		EffectiveStartDate: finchgo.F("effective_start_date"),
		EntityType:         finchgo.F(finchgo.HRISCompanyPayStatementItemRuleNewParamsEntityTypePayStatementItem),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISCompanyPayStatementItemRuleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Company.PayStatementItem.Rules.Update(
		context.TODO(),
		"rule_id",
		finchgo.HRISCompanyPayStatementItemRuleUpdateParams{
			EntityIDs:        finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
			OptionalProperty: finchgo.F[any](map[string]interface{}{}),
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

func TestHRISCompanyPayStatementItemRuleListWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Company.PayStatementItem.Rules.List(context.TODO(), finchgo.HRISCompanyPayStatementItemRuleListParams{
		EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestHRISCompanyPayStatementItemRuleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.Company.PayStatementItem.Rules.Delete(
		context.TODO(),
		"rule_id",
		finchgo.HRISCompanyPayStatementItemRuleDeleteParams{
			EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
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
