// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"context"
	"testing"
	"time"

	finchgo "github.com/Finch-API/finch-api-go"
	"github.com/Finch-API/finch-api-go/option"
)

func TestCancel(t *testing.T) {
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.ATS.Candidates.Get(cancelCtx, "<candidate id>")
	if err == nil && res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

func TestCancelDelay(t *testing.T) {
	client := finchgo.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAccessToken("AccessToken"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * time.Duration(2))
		cancel()
	}()
	res, err := client.ATS.Candidates.Get(cancelCtx, "<candidate id>")
	if err == nil && res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}
