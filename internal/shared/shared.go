// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/Finch-API/finch-api-go/internal/apijson"
)

type ConnectionStatusType string

const (
	ConnectionStatusTypePending             ConnectionStatusType = "pending"
	ConnectionStatusTypeProcessing          ConnectionStatusType = "processing"
	ConnectionStatusTypeConnected           ConnectionStatusType = "connected"
	ConnectionStatusTypeErrorNoAccountSetup ConnectionStatusType = "error_no_account_setup"
	ConnectionStatusTypeErrorPermissions    ConnectionStatusType = "error_permissions"
	ConnectionStatusTypeReauth              ConnectionStatusType = "reauth"
)

// - `supported`: This operation is supported by both the provider and Finch
//
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider
//
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support
//
//   - `client_access_only`: This behavior is supported by the provider, but only
//     available to the client and not to Finch
type OperationSupport string

const (
	OperationSupportSupported              OperationSupport = "supported"
	OperationSupportNotSupportedByFinch    OperationSupport = "not_supported_by_finch"
	OperationSupportNotSupportedByProvider OperationSupport = "not_supported_by_provider"
	OperationSupportClientAccessOnly       OperationSupport = "client_access_only"
)

type OperationSupportMatrix struct {
	// - `supported`: This operation is supported by both the provider and Finch
	//
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider
	//
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support
	//
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Create OperationSupport `json:"create"`
	// - `supported`: This operation is supported by both the provider and Finch
	//
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider
	//
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support
	//
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Delete OperationSupport `json:"delete"`
	// - `supported`: This operation is supported by both the provider and Finch
	//
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider
	//
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support
	//
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Read OperationSupport `json:"read"`
	// - `supported`: This operation is supported by both the provider and Finch
	//
	//   - `not_supported_by_finch`: This operation is not supported by Finch but
	//     supported by the provider
	//
	//   - `not_supported_by_provider`: This operation is not supported by the provider,
	//     so Finch cannot support
	//
	//   - `client_access_only`: This behavior is supported by the provider, but only
	//     available to the client and not to Finch
	Update OperationSupport           `json:"update"`
	JSON   operationSupportMatrixJSON `json:"-"`
}

// operationSupportMatrixJSON contains the JSON metadata for the struct
// [OperationSupportMatrix]
type operationSupportMatrixJSON struct {
	Create      apijson.Field
	Delete      apijson.Field
	Read        apijson.Field
	Update      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationSupportMatrix) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationSupportMatrixJSON) RawJSON() string {
	return r.raw
}

type Paging struct {
	// The total number of elements for the entire query (not just the given page)
	Count int64 `json:"count"`
	// The current start index of the returned list of elements
	Offset int64      `json:"offset"`
	JSON   pagingJSON `json:"-"`
}

// pagingJSON contains the JSON metadata for the struct [Paging]
type pagingJSON struct {
	Count       apijson.Field
	Offset      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Paging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagingJSON) RawJSON() string {
	return r.raw
}
