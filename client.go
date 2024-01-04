// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the Finch API. You should not instantiate this client directly,
// and instead use the [NewClient] method instead.
type Client struct {
	Options           []option.RequestOption
	HRIS              *HRISService
	Providers         *ProviderService
	Account           *AccountService
	Webhooks          *WebhookService
	RequestForwarding *RequestForwardingService
	Jobs              *JobService
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
	r.RequestForwarding = NewRequestForwardingService(opts...)
	r.Jobs = NewJobService(opts...)

	return
}

// Returns an access token for the Finch API given an authorization code. An
// authorization code can be obtained by visiting the URL returned by
// `GetAuthURL()`.
func (r *Client) GetAccessToken(ctx context.Context, code string, redirectUri string, opts ...option.RequestOption) (res string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append(opts[:], option.WithHeaderDel("authorization"))

	path := "/auth/token"

	var result map[string]string
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, nil, &result, opts...)
	if err != nil {
		return "", err
	}
	if cfg.ClientID == "" {
		return "", errors.New("expected ClientID to be set in order to call GetAccessToken")
	}
	if cfg.ClientSecret == "" {
		return "", errors.New("expected ClientSecret to be set in order to call GetAccessToken")
	}

	body := struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
		RedirectURI  string `json:"redirect_uri"`
	}{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Code:         code,
		RedirectURI:  redirectUri,
	}
	cfg.Apply(func(rc *requestconfig.RequestConfig) (err error) {
		rc.Buffer, err = json.Marshal(body)
		return err
	})

	err = cfg.Execute()
	if err != nil {
		return "", err
	}
	accessToken, ok := result["access_token"]
	if !ok {
		return "", errors.New("access_token not found in response")
	}

	return accessToken, nil
}

// Returns the authorization URL which can be visited in order to obtain an
// authorization code from Finch. The authorization code can then be exchanged for
// an access token for the Finch API by calling getAccessToken().
func (r *Client) GetAuthURL(products string, redirectUri string, sandbox bool, opts ...option.RequestOption) (res string, err error) {
	opts = append(r.Options[:], opts...)
	cfg := requestconfig.RequestConfig{}
	cfg.Apply(opts...)

	if cfg.ClientID == "" {
		return "", errors.New("expected the ClientID to be set in order to call GetAuthUrl")
	}
	u, err := url.Parse("https://connect.tryfinch.com/authorize")
	if err != nil {
		return "", err
	}
	q := u.Query()
	q.Set("client_id", cfg.ClientID)
	q.Set("products", products)
	q.Set("redirect_uri", redirectUri)
	q.Set("sandbox", strconv.FormatBool(sandbox))
	u.RawQuery = q.Encode()
	return u.String(), nil
}
