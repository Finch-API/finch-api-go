// File generated from our OpenAPI spec by Stainless.

package finchgo_test

import (
	"net/http"
	"testing"
	"time"

	finchgo "github.com/Finch-API/finch-api-go"
)

func TestVerifySignature(t *testing.T) {
	secret := "5WbX5kEWLlfzsGNjH64I8lOOqUB6e8FH"

	payload := `{"company_id":"720be419-0293-4d32-a707-32179b0827ab"}`

	header := http.Header{}
	header.Add("finch-event-id", "msg_2Lh9KRb0pzN4LePd3XiA4v12Axj")
	header.Add("finch-timestamp", "1676312382")
	header.Add("finch-signature", "v1,m7y0TV2C+hlHxU42wCieApTSTaA8/047OAplBqxIV/s=")

	client := finchgo.NewClient()
	err := client.Webhooks.VerifySignature([]byte(payload), header, secret, time.Unix(1676312382, 0))
	if err != nil {
		t.Fatalf("did not expect error %s", err.Error())
	}
}
