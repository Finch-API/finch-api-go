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
func (UnionString) ImplementsRequestForwardingForwardResponseRequestDataUnion()      {}
func (UnionString) ImplementsSandboxDirectoryNewParamsBodyCustomFieldsValueUnion()   {}
func (UnionString) ImplementsSandboxEmploymentUpdateResponseCustomFieldsValueUnion() {}
func (UnionString) ImplementsSandboxEmploymentUpdateParamsCustomFieldsValueUnion()   {}

type UnionBool bool

func (UnionBool) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {}
func (UnionBool) ImplementsSandboxDirectoryNewParamsBodyCustomFieldsValueUnion()            {}
func (UnionBool) ImplementsSandboxEmploymentUpdateResponseCustomFieldsValueUnion()          {}
func (UnionBool) ImplementsSandboxEmploymentUpdateParamsCustomFieldsValueUnion()            {}

type UnionFloat float64

func (UnionFloat) ImplementsEmploymentDataEmploymentDataResponseBodyCustomFieldsValueUnion() {}
func (UnionFloat) ImplementsSandboxDirectoryNewParamsBodyCustomFieldsValueUnion()            {}
func (UnionFloat) ImplementsSandboxEmploymentUpdateResponseCustomFieldsValueUnion()          {}
func (UnionFloat) ImplementsSandboxEmploymentUpdateParamsCustomFieldsValueUnion()            {}
