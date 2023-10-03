// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"
	"os"

	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Finch API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options   []option.RequestOption
	HRIS      *HRISService
	Providers *ProviderService
	Account   *AccountService
	Webhooks  *WebhookService
	Employer  *EmployerService
}

// NewClient generates a new client with the default option read from the
// environment (FINCH_CLIENT_ID, FINCH_CLIENT_SECRET, FINCH_WEBHOOK_SECRET). The
// option passed in as arguments are applied after these default arguments, and all
// option will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("FINCH_CLIENT_ID"); ok {
		defaults = append(defaults, option.WithClientID(o))
	}
	if o, ok := os.LookupEnv("FINCH_CLIENT_SECRET"); ok {
		defaults = append(defaults, option.WithClientSecret(o))
	}
	if o, ok := os.LookupEnv("FINCH_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.HRIS = NewHRISService(opts...)
	r.Providers = NewProviderService(opts...)
	r.Account = NewAccountService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.Employer = NewEmployerService(opts...)

	return
}

// The Forward API allows you to make direct requests to an employment system. If
// Finch’s unified API doesn’t have a data model that cleanly fits your needs, then
// Forward allows you to push or pull data models directly against an integration’s
// API.
func (r *Client) Forward(ctx context.Context, body FinchgoForwardParams, opts ...option.RequestOption) (res *FinchgoForwardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "forward"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}
