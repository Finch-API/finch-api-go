// Non-generated wrapper for /employer/company endpoint to handle 202 responses from assisted connections.
// This file will not be overwritten by codegen.

package finchgo

import (
	"context"
	"net/http"
	"reflect"
	"slices"

	"github.com/Finch-API/finch-api-go/internal/apijson"
	"github.com/Finch-API/finch-api-go/internal/requestconfig"
	"github.com/Finch-API/finch-api-go/option"
	"github.com/tidwall/gjson"
)

// GetWithAssistedSupport reads basic company data with support for 202 responses
// from assisted connections. Returns a union type that can be either:
// - Company (200 OK with Company data)
// - CompanyDataSyncInProgress (202 Accepted with sync status)
func (r *HRISCompanyService) GetWithAssistedSupport(ctx context.Context, opts ...option.RequestOption) (res *CompanyResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "employer/company"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// CompanyResponse is the union wrapper for response variants
type CompanyResponse struct {
	Code      float64 `json:"code"`
	FinchCode string  `json:"finch_code"`
	Message   string  `json:"message"`
	Name      string  `json:"name"`
	JSON      companyResponseJSON `json:"-"`
	union     CompanyResponseUnion
}

// companyResponseJSON contains the JSON metadata for the struct [CompanyResponse]
type companyResponseJSON struct {
	Code        apijson.Field
	FinchCode   apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r companyResponseJSON) RawJSON() string {
	return r.raw
}

func (r *CompanyResponse) UnmarshalJSON(data []byte) (err error) {
	*r = CompanyResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CompanyResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [Company],
// [CompanyDataSyncInProgress].
func (r CompanyResponse) AsUnion() CompanyResponseUnion {
	return r.union
}

// Union satisfied by [Company] or [CompanyDataSyncInProgress].
type CompanyResponseUnion interface {
	implementsCompanyResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CompanyResponseUnion)(nil)).Elem(),
		"code",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(Company{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CompanyDataSyncInProgress{}),
		},
	)
}

// Make Company satisfy the union interface
func (r Company) implementsCompanyResponse() {}

// CompanyDataSyncInProgress represents a 202 response when data sync is in progress
type CompanyDataSyncInProgress struct {
	Code      CompanyDataSyncInProgressCode      `json:"code,required"`
	FinchCode CompanyDataSyncInProgressFinchCode `json:"finch_code,required"`
	Message   CompanyDataSyncInProgressMessage   `json:"message,required"`
	Name      CompanyDataSyncInProgressName      `json:"name,required"`
	JSON      companyDataSyncInProgressJSON      `json:"-"`
}

// companyDataSyncInProgressJSON contains the JSON metadata for the struct
// [CompanyDataSyncInProgress]
type companyDataSyncInProgressJSON struct {
	Code        apijson.Field
	FinchCode   apijson.Field
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompanyDataSyncInProgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r companyDataSyncInProgressJSON) RawJSON() string {
	return r.raw
}

func (r CompanyDataSyncInProgress) implementsCompanyResponse() {}

type CompanyDataSyncInProgressCode float64

const (
	CompanyDataSyncInProgressCode202 CompanyDataSyncInProgressCode = 202
)

func (r CompanyDataSyncInProgressCode) IsKnown() bool {
	switch r {
	case CompanyDataSyncInProgressCode202:
		return true
	}
	return false
}

type CompanyDataSyncInProgressFinchCode string

const (
	CompanyDataSyncInProgressFinchCodeDataSyncInProgress CompanyDataSyncInProgressFinchCode = "data_sync_in_progress"
)

func (r CompanyDataSyncInProgressFinchCode) IsKnown() bool {
	switch r {
	case CompanyDataSyncInProgressFinchCodeDataSyncInProgress:
		return true
	}
	return false
}

type CompanyDataSyncInProgressMessage string

const (
	CompanyDataSyncInProgressMessageTheCompanyDataIsBeingFetchedPleaseCheckBackLater CompanyDataSyncInProgressMessage = "The company data is being fetched. Please check back later."
)

func (r CompanyDataSyncInProgressMessage) IsKnown() bool {
	switch r {
	case CompanyDataSyncInProgressMessageTheCompanyDataIsBeingFetchedPleaseCheckBackLater:
		return true
	}
	return false
}

type CompanyDataSyncInProgressName string

const (
	CompanyDataSyncInProgressNameAccepted CompanyDataSyncInProgressName = "accepted"
)

func (r CompanyDataSyncInProgressName) IsKnown() bool {
	switch r {
	case CompanyDataSyncInProgressNameAccepted:
		return true
	}
	return false
}
