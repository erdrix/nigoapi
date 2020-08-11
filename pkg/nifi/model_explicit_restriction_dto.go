/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.11.4
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type ExplicitRestrictionDto struct {
	// The required permission necessary for this restriction.
	RequiredPermission *RequiredPermissionDto `json:"requiredPermission,omitempty"`
	// The description of why the usage of this component is restricted for this required permission.
	Explanation string `json:"explanation,omitempty"`
}
