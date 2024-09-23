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

func TestRequestForwardingForwardWithOptionalParams(t *testing.T) {
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
	_, err := client.RequestForwarding.Forward(context.TODO(), finchgo.RequestForwardingForwardParams{
		Method: finchgo.F("POST"),
		Route:  finchgo.F("/people/search"),
		Data:   finchgo.Null[string](),
		Headers: finchgo.F[any](map[string]interface{}{
			"content-type": "application/json",
		}),
		Params: finchgo.F[any](map[string]interface{}{
			"showInactive":  true,
			"humanReadable": true,
		}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
