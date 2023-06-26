// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/apiquery"
	"github.com/Finch-API/finch-api-go/internal/param"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
)

// ATSOfferService contains methods and other services that help with interacting
// with the Finch API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewATSOfferService] method instead.
type ATSOfferService struct {
	Options []option.RequestOption
}

// NewATSOfferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewATSOfferService(opts ...option.RequestOption) (r *ATSOfferService) {
	r = &ATSOfferService{}
	r.Options = opts
	return
}

// Get a single offer from an organization.
func (r *ATSOfferService) Get(ctx context.Context, offerID string, opts ...option.RequestOption) (res *Offer, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ats/offers/%s", offerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all offers put out by an organization.
func (r *ATSOfferService) List(ctx context.Context, query ATSOfferListParams, opts ...option.RequestOption) (res *OffersPage, err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ats/offers"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all offers put out by an organization.
func (r *ATSOfferService) ListAutoPaging(ctx context.Context, query ATSOfferListParams, opts ...option.RequestOption) *OffersPageAutoPager {
	return NewOffersPageAutoPager(r.List(ctx, query, opts...))
}

type OffersPage struct {
	Paging Paging  `json:"paging,required"`
	Offers []Offer `json:"offers,required"`
	JSON   offersPageJSON
	cfg    *requestconfig.RequestConfig
	res    *http.Response
}

// offersPageJSON contains the JSON metadata for the struct [OffersPage]
type offersPageJSON struct {
	Paging      apijson.Field
	Offers      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OffersPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *OffersPage) GetNextPage() (res *OffersPage, err error) {
	// This page represents a response that isn't actually paginated at the API level
	// so there will never be a next page.
	cfg := (*requestconfig.RequestConfig)(nil)
	if cfg == nil {
		return nil, nil
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OffersPage) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type OffersPageAutoPager struct {
	page *OffersPage
	cur  Offer
	idx  int
	run  int
	err  error
}

func NewOffersPageAutoPager(page *OffersPage, err error) *OffersPageAutoPager {
	return &OffersPageAutoPager{
		page: page,
		err:  err,
	}
}

func (r *OffersPageAutoPager) Next() bool {
	if r.page == nil || len(r.page.Offers) == 0 {
		return false
	}
	if r.idx >= len(r.page.Offers) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Offers[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OffersPageAutoPager) Current() Offer {
	return r.cur
}

func (r *OffersPageAutoPager) Err() error {
	return r.err
}

func (r *OffersPageAutoPager) Index() int {
	return r.run
}

type Offer struct {
	ID            string      `json:"id,required" format:"uuid"`
	ApplicationID string      `json:"application_id,required" format:"uuid"`
	CandidateID   string      `json:"candidate_id,required" format:"uuid"`
	JobID         string      `json:"job_id,required" format:"uuid"`
	CreatedAt     time.Time   `json:"created_at,required" format:"date-time"`
	UpdatedAt     time.Time   `json:"updated_at,required" format:"date-time"`
	Status        OfferStatus `json:"status,required"`
	JSON          offerJSON
}

// offerJSON contains the JSON metadata for the struct [Offer]
type offerJSON struct {
	ID            apijson.Field
	ApplicationID apijson.Field
	CandidateID   apijson.Field
	JobID         apijson.Field
	CreatedAt     apijson.Field
	UpdatedAt     apijson.Field
	Status        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Offer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type OfferStatus string

const (
	OfferStatusSigned       OfferStatus = "signed"
	OfferStatusRejected     OfferStatus = "rejected"
	OfferStatusDraft        OfferStatus = "draft"
	OfferStatusApprovalSent OfferStatus = "approval-sent"
	OfferStatusApproved     OfferStatus = "approved"
	OfferStatusSent         OfferStatus = "sent"
	OfferStatusSentManually OfferStatus = "sent-manually"
	OfferStatusOpened       OfferStatus = "opened"
	OfferStatusArchived     OfferStatus = "archived"
)

type ATSOfferListParams struct {
	// Number of offers to return (defaults to all)
	Limit param.Field[int64] `query:"limit"`
	// Index to start from (defaults to 0)
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [ATSOfferListParams]'s query parameters as `url.Values`.
func (r ATSOfferListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
