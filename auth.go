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

// AuthService contains methods and other services that help with interacting with
// the Finch API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewAuthService] method instead.
type AuthService struct {
	Options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r *AuthService) {
	r = &AuthService{}
	r.Options = opts
	return
}

// Exchange the authorization code for an access token
func (r *AuthService) NewToken(ctx context.Context, body AuthNewTokenParams, opts ...option.RequestOption) (res *CreateAccessTokenResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "auth/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AuthNewTokenParams struct {
	ClientID     param.Field[string] `json:"client_id,required"`
	ClientSecret param.Field[string] `json:"client_secret,required"`
	Code         param.Field[string] `json:"code,required"`
	RedirectUri  param.Field[string] `json:"redirect_uri,required"`
}

func (r AuthNewTokenParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
