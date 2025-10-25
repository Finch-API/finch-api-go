// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo_test

import (
	"context"
	"os"
	"testing"

	"github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestManualPagination(t *testing.T) {
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
	page, err := client.HRIS.Directory.List(context.TODO(), finchgo.HRISDirectoryListParams{
		EntityIDs: finchgo.F([]string{"550e8400-e29b-41d4-a716-446655440000"}),
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, directory := range page.Individuals {
		t.Logf("%+v\n", directory.ID)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, directory := range page.Individuals {
			t.Logf("%+v\n", directory.ID)
		}
	}
}
