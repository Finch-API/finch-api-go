// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apierror"
	"github.com/Finch-API/finch-api-go/internal/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type IntrospectResponseConnectionStatus = shared.IntrospectResponseConnectionStatus

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusPending = shared.IntrospectResponseConnectionStatusPending

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusProcessing = shared.IntrospectResponseConnectionStatusProcessing

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusConnected = shared.IntrospectResponseConnectionStatusConnected

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusErrorNoAccountSetup = shared.IntrospectResponseConnectionStatusErrorNoAccountSetup

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusErrorPermissions = shared.IntrospectResponseConnectionStatusErrorPermissions

// This is an alias to an internal value.
const IntrospectResponseConnectionStatusReauth = shared.IntrospectResponseConnectionStatusReauth

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
