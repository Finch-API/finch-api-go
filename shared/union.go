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

func (UnionString) ImplementsEmploymentDataObjectCustomFieldsValueUnion()           {}
func (UnionString) ImplementsIntrospectionConnectionStatusLastSuccessfulSyncUnion() {}
func (UnionString) ImplementsIntrospectionAuthenticationMethodsConnectionStatusLastSuccessfulSyncUnion() {
}

type UnionBool bool

func (UnionBool) ImplementsEmploymentDataObjectCustomFieldsValueUnion() {}

type UnionFloat float64

func (UnionFloat) ImplementsEmploymentDataObjectCustomFieldsValueUnion() {}
