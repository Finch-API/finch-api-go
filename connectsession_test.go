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

func TestConnectSessionNew(t *testing.T) {
	t.Skip("prism tests are broken")
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
	_, err := client.Connect.Sessions.New(context.TODO(), finchgo.ConnectSessionNewParams{
		CustomerEmail: finchgo.F("dev@stainless.com"),
		CustomerID:    finchgo.F("x"),
		CustomerName:  finchgo.F("x"),
		Integration: finchgo.F(finchgo.ConnectSessionNewParamsIntegration{
			AuthMethod: finchgo.F(finchgo.ConnectSessionNewParamsIntegrationAuthMethodAssisted),
			Provider:   finchgo.F("provider"),
		}),
		Manual:          finchgo.F(true),
		MinutesToExpire: finchgo.F(1.000000),
		Products:        finchgo.F([]finchgo.ConnectSessionNewParamsProduct{finchgo.ConnectSessionNewParamsProductBenefits}),
		RedirectUri:     finchgo.F("redirect_uri"),
		Sandbox:         finchgo.F(finchgo.ConnectSessionNewParamsSandboxFinch),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConnectSessionReauthenticate(t *testing.T) {
	t.Skip("prism tests are broken")
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
	_, err := client.Connect.Sessions.Reauthenticate(context.TODO(), finchgo.ConnectSessionReauthenticateParams{
		ConnectionID:    finchgo.F("connection_id"),
		MinutesToExpire: finchgo.F(int64(0)),
		Products:        finchgo.F([]finchgo.ConnectSessionReauthenticateParamsProduct{finchgo.ConnectSessionReauthenticateParamsProductBenefits}),
		RedirectUri:     finchgo.F("https://example.com"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
