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

type AccessPolicy struct {
	// The id of the policy. Set by server at creation time.
	Identifier string `json:"identifier,omitempty"`
	// The resource for this access policy.
	Resource string `json:"resource"`
	// The action associated with this access policy.
	Action string `json:"action"`
	// Indicates if this access policy is configurable, based on which Authorizer has been configured to manage it.
	Configurable bool `json:"configurable,omitempty"`
	// The revision of this entity used for optimistic-locking during updates.
	Revision *RevisionInfo `json:"revision,omitempty"`
	// The set of user IDs associated with this access policy.
	Users []Tenant `json:"users,omitempty"`
	// The set of user group IDs associated with this access policy.
	UserGroups []Tenant `json:"userGroups,omitempty"`
}
