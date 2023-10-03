// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"os"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestFinchgoForwardWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAccessToken("AccessToken"),
	)
	_, err := client.Forward(context.TODO(), finchgo.FinchgoForwardParams{
		Method:  finchgo.F("string"),
		Route:   finchgo.F("string"),
		Data:    finchgo.F("string"),
		Headers: finchgo.F[any](map[string]interface{}{}),
		Params:  finchgo.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
