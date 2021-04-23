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

type AllowableValue struct {
	// The value of the allowable value
	Value string `json:"value,omitempty"`
	// The display name of the allowable value
	DisplayName string `json:"displayName,omitempty"`
	// The description of the allowable value
	Description string `json:"description,omitempty"`
}
