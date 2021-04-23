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

type JaxbLink struct {
	// The href for the link
	Href string `json:"href,omitempty"`
	// The params for the link
	Params map[string]string `json:"params,omitempty"`
}
