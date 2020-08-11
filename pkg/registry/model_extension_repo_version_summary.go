/*
 * Apache NiFi Registry REST API
 *
 * The REST API provides an interface to a registry with operations for saving, versioning, reading NiFi flows and components.
 *
 * API version: 0.7.0
 * Contact: dev@nifi.apache.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package registry

type ExtensionRepoVersionSummary struct {
	// An WebLink to this entity.
	Link *JaxbLink `json:"link,omitempty"`
	// The bucket name
	BucketName string `json:"bucketName,omitempty"`
	// The group id
	GroupId string `json:"groupId,omitempty"`
	// The artifact id
	ArtifactId string `json:"artifactId,omitempty"`
	// The version
	Version string `json:"version,omitempty"`
	// The identity of the user that created this version
	Author string `json:"author,omitempty"`
	// The timestamp of when this version was created
	Timestamp int64 `json:"timestamp,omitempty"`
}
