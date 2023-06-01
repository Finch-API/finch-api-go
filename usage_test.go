package finchgo_test

import (
	"context"
	"fmt"
	"testing"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func TestUsage(t *testing.T) {
	client := finchgo.NewClient(option.WithAccessToken("AccessToken"), option.WithBaseURL("http://127.0.0.1:4010"))
	directories, err := client.HRIS.Directory.ListIndividuals(context.TODO(), finchgo.HRISDirectoryListIndividualsParams{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", directories)
}
