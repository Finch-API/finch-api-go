// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"
)

type UnionTime time.Time

func (UnionTime) ImplementsIntrospectionConnectionStatusLastSuccessfulSyncUnion() {}
func (UnionTime) ImplementsIntrospectionAuthenticationMethodsConnectionStatusLastSuccessfulSyncUnion() {
}

type UnionString string

func (UnionString) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {}
func (UnionString) ImplementsIntrospectionConnectionStatusLastSuccessfulSyncUnion()           {}
func (UnionString) ImplementsIntrospectionAuthenticationMethodsConnectionStatusLastSuccessfulSyncUnion() {
}
func (UnionString) ImplementsRequestForwardingForwardResponseRequestDataUnion() {}

type UnionBool bool

func (UnionBool) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {}
