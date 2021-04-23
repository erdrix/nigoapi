/*
 * Apache NiFi Registry REST API
 *
 * The REST API provides an interface to a registry with operations for saving, versioning, reading NiFi flows and components.
 *
 * API version: 0.8.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package registry

type VersionedFlowSnapshotMetadata struct {
	// An WebLink to this entity.
	Link *JaxbLink `json:"link,omitempty"`
	// The identifier of the bucket this snapshot belongs to.
	BucketIdentifier string `json:"bucketIdentifier"`
	// The identifier of the flow this snapshot belongs to.
	FlowIdentifier string `json:"flowIdentifier"`
	// The version of this snapshot of the flow.
	Version int32 `json:"version"`
	// The timestamp when the flow was saved, as milliseconds since epoch.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The user that created this snapshot of the flow.
	Author string `json:"author,omitempty"`
	// The comments provided by the user when creating the snapshot.
	Comments string `json:"comments,omitempty"`
}
