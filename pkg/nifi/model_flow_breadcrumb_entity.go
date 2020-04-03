/*
 * NiFi Rest Api
 *
 * The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.
 *
 * API version: 1.11.1-SNAPSHOT
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nifi

type FlowBreadcrumbEntity struct {
	// The id of this ancestor ProcessGroup.
	Id string `json:"id,omitempty"`
	// The permissions for this ancestor ProcessGroup.
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	// The current state of the Process Group, as it relates to the Versioned Flow
	VersionedFlowState string `json:"versionedFlowState,omitempty"`
	// This breadcrumb.
	Breadcrumb *FlowBreadcrumbDto `json:"breadcrumb,omitempty"`
	// The parent breadcrumb for this breadcrumb.
	ParentBreadcrumb *FlowBreadcrumbEntity `json:"parentBreadcrumb,omitempty"`
}