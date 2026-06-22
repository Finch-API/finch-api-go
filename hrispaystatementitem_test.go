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

func TestHRISPayStatementItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.HRIS.PayStatementItem.List(context.TODO(), finchgo.HRISPayStatementItemListParams{
		Categories: finchgo.F([]finchgo.HRISPayStatementItemListParamsCategory{finchgo.HRISPayStatementItemListParamsCategoryEarnings}),
		EndDate:    finchgo.F(time.Now()),
		EntityIDs:  finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
		Name:       finchgo.F("name"),
		StartDate:  finchgo.F(time.Now()),
		Type:       finchgo.F("base_compensation"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
