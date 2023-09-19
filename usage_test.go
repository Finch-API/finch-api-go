// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestUsage(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	page, err := client.HRIS.Directory.ListIndividuals(context.TODO(), finchgo.HRISDirectoryListIndividualsParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", page)
}
