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

func TestAccessTokenNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AccessTokens.New(context.TODO(), finchgo.AccessTokenNewParams{
		Code:         finchgo.F("<your_authorization_code>"),
		ClientID:     finchgo.F("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
		ClientSecret: finchgo.F("<your_client_secret>"),
		RedirectUri:  finchgo.F("https://example.com"),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
