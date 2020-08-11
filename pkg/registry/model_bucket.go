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

type Bucket struct {
	// An WebLink to this entity.
	Link *JaxbLink `json:"link,omitempty"`
	// An ID to uniquely identify this object.
	Identifier string `json:"identifier,omitempty"`
	// The name of the bucket.
	Name string `json:"name"`
	// The timestamp of when the bucket was first created. This is set by the server at creation time.
	CreatedTimestamp int64 `json:"createdTimestamp,omitempty"`
	// A description of the bucket.
	Description string `json:"description,omitempty"`
	// Indicates if this bucket allows the same version of an extension bundle to be redeployed and thus overwrite the existing artifact. By default this is false.
	AllowBundleRedeploy bool `json:"allowBundleRedeploy,omitempty"`
	// Indicates if this bucket allows read access to unauthenticated anonymous users
	AllowPublicRead bool `json:"allowPublicRead,omitempty"`
	// The access that the current user has to this bucket.
	Permissions *Permissions `json:"permissions,omitempty"`
	// The revision of this entity used for optimistic-locking during updates.
	Revision *RevisionInfo `json:"revision,omitempty"`
}
