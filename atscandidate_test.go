// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"errors"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestATSCandidateGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.ATS.Candidates.Get(context.TODO(), "string")
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestATSCandidateListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(
		option.WithAccessToken("AccessToken"),
		option.WithBaseURL("http://127.0.0.1:4010"),
	)
	_, err := client.ATS.Candidates.List(context.TODO(), finchgo.ATSCandidateListParams{
		Limit:  finchgo.F(int64(0)),
		Offset: finchgo.F(int64(0)),
	})
	if err != nil {
		var apierr *finchgo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
