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

type DeprecationNotice struct {
	// The reason for the deprecation
	Reason string `json:"reason,omitempty"`
	// The alternatives to use
	Alternatives []string `json:"alternatives,omitempty"`
}
