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

type ConnectionEntity struct {
	// The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.
	Revision *RevisionDto `json:"revision,omitempty"`
	// The id of the component.
	Id string `json:"id,omitempty"`
	// The URI for futures requests to the component.
	Uri string `json:"uri,omitempty"`
	// The position of this component in the UI if applicable.
	Position *PositionDto `json:"position,omitempty"`
	// The permissions for this component.
	Permissions *PermissionsDto `json:"permissions,omitempty"`
	// The bulletins for this component.
	Bulletins []BulletinEntity `json:"bulletins,omitempty"`
	// Acknowledges that this node is disconnected to allow for mutable requests to proceed.
	DisconnectedNodeAcknowledged bool `json:"disconnectedNodeAcknowledged,omitempty"`
	Component *ConnectionDto `json:"component,omitempty"`
	// The status of the connection.
	Status *ConnectionStatusDto `json:"status,omitempty"`
	// The bend points on the connection.
	Bends []PositionDto `json:"bends,omitempty"`
	// The index of the bend point where to place the connection label.
	LabelIndex int32 `json:"labelIndex,omitempty"`
	// The z index of the connection.
	GetzIndex int64 `json:"getzIndex,omitempty"`
	// The identifier of the source of this connection.
	SourceId string `json:"sourceId,omitempty"`
	// The identifier of the group of the source of this connection.
	SourceGroupId string `json:"sourceGroupId,omitempty"`
	// The type of component the source connectable is.
	SourceType string `json:"sourceType"`
	// The identifier of the destination of this connection.
	DestinationId string `json:"destinationId,omitempty"`
	// The identifier of the group of the destination of this connection.
	DestinationGroupId string `json:"destinationGroupId,omitempty"`
	// The type of component the destination connectable is.
	DestinationType string `json:"destinationType"`
}