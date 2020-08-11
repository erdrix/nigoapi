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

type SnippetDto struct {
	// The id of the snippet.
	Id string `json:"id,omitempty"`
	// The URI of the snippet.
	Uri string `json:"uri,omitempty"`
	// The group id for the components in the snippet.
	ParentGroupId string `json:"parentGroupId,omitempty"`
	// The ids of the process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	ProcessGroups map[string]RevisionDto `json:"processGroups,omitempty"`
	// The ids of the remote process groups in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	RemoteProcessGroups map[string]RevisionDto `json:"remoteProcessGroups,omitempty"`
	// The ids of the processors in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Processors map[string]RevisionDto `json:"processors,omitempty"`
	// The ids of the input ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	InputPorts map[string]RevisionDto `json:"inputPorts,omitempty"`
	// The ids of the output ports in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	OutputPorts map[string]RevisionDto `json:"outputPorts,omitempty"`
	// The ids of the connections in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Connections map[string]RevisionDto `json:"connections,omitempty"`
	// The ids of the labels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Labels map[string]RevisionDto `json:"labels,omitempty"`
	// The ids of the funnels in this snippet. These ids will be populated within each response. They can be specified when creating a snippet. However, once a snippet has been created its contents cannot be modified (these ids are ignored during update requests).
	Funnels map[string]RevisionDto `json:"funnels,omitempty"`
}
