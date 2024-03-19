// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apierror"
	"github.com/Finch-API/finch-api-go/internal/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type ConnectionStatusType = shared.ConnectionStatusType

// This is an alias to an internal value.
const ConnectionStatusTypePending = shared.ConnectionStatusTypePending

// This is an alias to an internal value.
const ConnectionStatusTypeProcessing = shared.ConnectionStatusTypeProcessing

// This is an alias to an internal value.
const ConnectionStatusTypeConnected = shared.ConnectionStatusTypeConnected

// This is an alias to an internal value.
const ConnectionStatusTypeErrorNoAccountSetup = shared.ConnectionStatusTypeErrorNoAccountSetup

// This is an alias to an internal value.
const ConnectionStatusTypeErrorPermissions = shared.ConnectionStatusTypeErrorPermissions

// This is an alias to an internal value.
const ConnectionStatusTypeReauth = shared.ConnectionStatusTypeReauth

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
//
// This is an alias to an internal type.
type OperationSupport = shared.OperationSupport

// This is an alias to an internal value.
const OperationSupportSupported = shared.OperationSupportSupported

// This is an alias to an internal value.
const OperationSupportNotSupportedByFinch = shared.OperationSupportNotSupportedByFinch

// This is an alias to an internal value.
const OperationSupportNotSupportedByProvider = shared.OperationSupportNotSupportedByProvider

// This is an alias to an internal value.
const OperationSupportClientAccessOnly = shared.OperationSupportClientAccessOnly

// This is an alias to an internal type.
type OperationSupportMatrix = shared.OperationSupportMatrix

// This is an alias to an internal type.
type Paging = shared.Paging
