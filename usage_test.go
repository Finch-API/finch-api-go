// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"fmt"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/internal/testutil"
	"github.com/Finch-API/finch-api-go/option"
)

func TestUsage(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	directories, err := client.HRIS.Directory.ListIndividuals(context.TODO(), finchgo.HRISDirectoryListIndividualsParams{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", directories)
}
