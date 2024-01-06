// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"net/http"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// AccessTokenService contains methods and other services that help with
// interacting with the Finch API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccessTokenService] method instead.
type AccessTokenService struct {
	Options []option.RequestOption
}

// NewAccessTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessTokenService(opts ...option.RequestOption) (r *AccessTokenService) {
	r = &AccessTokenService{}
	r.Options = opts
	return
}

// Exchange the authorization code for an access token
func (r *AccessTokenService) New(ctx context.Context, body AccessTokenNewParams, opts ...option.RequestOption) (res *CreateAccessTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CreateAccessTokenResponse struct {
	AccessToken string                        `json:"access_token,required"`
	JSON        createAccessTokenResponseJSON `json:"-"`
}

// createAccessTokenResponseJSON contains the JSON metadata for the struct
// [CreateAccessTokenResponse]
type createAccessTokenResponseJSON struct {
	AccessToken apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CreateAccessTokenResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessTokenNewParams struct {
	ClientID     param.Field[string] `json:"client_id,required"`
	ClientSecret param.Field[string] `json:"client_secret,required"`
	Code         param.Field[string] `json:"code,required"`
	RedirectUri  param.Field[string] `json:"redirect_uri,required"`
}

func (r AccessTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
