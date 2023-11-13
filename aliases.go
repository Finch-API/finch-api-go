// File generated from our OpenAPI spec by Stainless.

package finchgo

import (
	"github.com/Finch-API/finch-api-go/internal/apierror"
	"github.com/Finch-API/finch-api-go/internal/shared"
)

type Error = apierror.Error

//   - `supported`: This operation is supported by both the provider and Finch <br>
//   - `not_supported_by_finch`: This operation is not supported by Finch but
//     supported by the provider <br>
//   - `not_supported_by_provider`: This operation is not supported by the provider,
//     so Finch cannot support <br>
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
