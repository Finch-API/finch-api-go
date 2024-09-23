// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"bytes"
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
	AccessTokens      *AccessTokenService
	HRIS              *HRISService
	Providers         *ProviderService
	Account           *AccountService
	Webhooks          *WebhookService
	RequestForwarding *RequestForwardingService
	Jobs              *JobService
	Sandbox           *SandboxService
	Payroll           *PayrollService
	Connect           *ConnectService
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

	r.AccessTokens = NewAccessTokenService(opts...)
	r.HRIS = NewHRISService(opts...)
	r.Providers = NewProviderService(opts...)
	r.Account = NewAccountService(opts...)
	r.Webhooks = NewWebhookService(opts...)
	r.RequestForwarding = NewRequestForwardingService(opts...)
	r.Jobs = NewJobService(opts...)
	r.Sandbox = NewSandboxService(opts...)
	r.Payroll = NewPayrollService(opts...)
	r.Connect = NewConnectService(opts...)

	return
}

// DEPRECATED: use client.accessTokens().create instead.
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
		buf, err := json.Marshal(body)
		if err != nil {
			return err
		}
		rc.Body = bytes.NewBuffer(buf)
		rc.Request.Header.Set("Content-Type", "application/json")
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

// Returns a copy of the current Finch client with the given access token for
// authentication.
func (r *Client) WithAccessToken(accessToken string) (res Client, err error) {
	opts := append(r.Options[:], option.WithAccessToken(accessToken))
	return Client{Options: opts}, nil
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	opts = append(r.Options, opts...)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params interface{}, res interface{}, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
