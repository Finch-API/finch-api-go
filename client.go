// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"os"

	"github.com/Finch-API/finch-api-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Finch API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options           []option.RequestOption
	AccessTokens      *AccessTokenService
	HRIS              *HRISService
	Providers         *ProviderService
	Account           *AccountService
	Webhooks          *WebhookService
	RequestForwarding *RequestForwardingService
	Jobs              *JobService
	Sandbox           *SandboxService
}

// NewClient generates a new client with the default option read from the
// environment (FINCH_CLIENT_ID, FINCH_CLIENT_SECRET, FINCH_SANDBOX_CLIENT_ID,
// FINCH_SANDBOX_CLIENT_SECRET, FINCH_WEBHOOK_SECRET). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("FINCH_CLIENT_ID"); ok {
		defaults = append(defaults, option.WithClientID(o))
	}
	if o, ok := os.LookupEnv("FINCH_CLIENT_SECRET"); ok {
		defaults = append(defaults, option.WithClientSecret(o))
	}
	if o, ok := os.LookupEnv("FINCH_SANDBOX_CLIENT_ID"); ok {
		defaults = append(defaults, option.WithSandboxClientID(o))
	}
	if o, ok := os.LookupEnv("FINCH_SANDBOX_CLIENT_SECRET"); ok {
		defaults = append(defaults, option.WithSandboxClientSecret(o))
	}
	if o, ok := os.LookupEnv("FINCH_WEBHOOK_SECRET"); ok {
		defaults = append(defaults, option.WithWebhookSecret(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.AccessTokens = NewAccessTokenService(opts...)
	r.HRIS = NewHRISService(opts...)
	r.Providers = NewProviderService(opts...)
	r.Account = NewAccountService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.RequestForwarding = NewRequestForwardingService(opts...)
	r.Jobs = NewJobService(opts...)
	r.Sandbox = NewSandboxService(opts...)

	return
}
