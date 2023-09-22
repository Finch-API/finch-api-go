// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestManualPagination(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	page, err := client.HRIS.Directory.ListIndividuals(context.TODO(), finchgo.HRISDirectoryListIndividualsParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, directory := range page.Individuals {
		t.Logf("%+v\n", directory)
	}
	// Prism mock isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, directory := range page.Individuals {
			t.Logf("%+v\n", directory)
		}
	}
}
