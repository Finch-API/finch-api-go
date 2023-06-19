// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestAutoPagination(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	iter := client.ATS.Jobs.ListAutoPaging(context.TODO(), finchgo.ATSJobListParams{})
	// Prism mock isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		job := iter.Current()
		t.Logf("%+v\n", job)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
