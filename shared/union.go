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

func (UnionString) ImplementsEmploymentDataEmploymentDataCustomFieldsValueUnion()   {}
func (UnionString) ImplementsIntrospectionConnectionStatusLastSuccessfulSyncUnion() {}
func (UnionString) ImplementsIntrospectionAuthenticationMethodsConnectionStatusLastSuccessfulSyncUnion() {
}
func (UnionString) ImplementsRequestForwardingForwardResponseRequestDataUnion() {}

type UnionBool bool

func (UnionBool) ImplementsEmploymentDataEmploymentDataCustomFieldsValueUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsEmploymentDataEmploymentDataCustomFieldsValueUnion() {}
