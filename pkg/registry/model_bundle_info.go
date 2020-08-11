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

type BundleInfo struct {
	// The id of the bucket where the bundle is located
	BucketId string `json:"bucketId,omitempty"`
	// The name of the bucket where the bundle is located
	BucketName string `json:"bucketName,omitempty"`
	// The id of the bundle
	BundleId string `json:"bundleId,omitempty"`
	// The type of bundle (i.e. a NiFi NAR vs MiNiFi CPP)
	BundleType string `json:"bundleType,omitempty"`
	// The group id of the bundle
	GroupId string `json:"groupId,omitempty"`
	// The artifact id of the bundle
	ArtifactId string `json:"artifactId,omitempty"`
	// The version of the bundle
	Version string `json:"version,omitempty"`
	// The version of the system API the bundle was built against
	SystemApiVersion string `json:"systemApiVersion,omitempty"`
}
