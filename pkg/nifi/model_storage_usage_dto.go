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

type StorageUsageDto struct {
	// The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.
	Identifier string `json:"identifier,omitempty"`
	// Amount of free space.
	FreeSpace string `json:"freeSpace,omitempty"`
	// Amount of total space.
	TotalSpace string `json:"totalSpace,omitempty"`
	// Amount of used space.
	UsedSpace string `json:"usedSpace,omitempty"`
	// The number of bytes of free space.
	FreeSpaceBytes int64 `json:"freeSpaceBytes,omitempty"`
	// The number of bytes of total space.
	TotalSpaceBytes int64 `json:"totalSpaceBytes,omitempty"`
	// The number of bytes of used space.
	UsedSpaceBytes int64 `json:"usedSpaceBytes,omitempty"`
	// Utilization of this storage location.
	Utilization string `json:"utilization,omitempty"`
}
