// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type CmkType string

// Enum values for CmkType
const (
	CmkTypeCmCmk CmkType = "CUSTOMER_MANAGED_KMS_KEY"
	CmkTypeAoCmk CmkType = "AWS_OWNED_KMS_KEY"
)

// Values returns all known values for CmkType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (CmkType) Values() []CmkType {
	return []CmkType{
		"CUSTOMER_MANAGED_KMS_KEY",
		"AWS_OWNED_KMS_KEY",
	}
}

type DatastoreStatus string

// Enum values for DatastoreStatus
const (
	DatastoreStatusCreating DatastoreStatus = "CREATING"
	DatastoreStatusActive   DatastoreStatus = "ACTIVE"
	DatastoreStatusDeleting DatastoreStatus = "DELETING"
	DatastoreStatusDeleted  DatastoreStatus = "DELETED"
)

// Values returns all known values for DatastoreStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DatastoreStatus) Values() []DatastoreStatus {
	return []DatastoreStatus{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

type FHIRVersion string

// Enum values for FHIRVersion
const (
	FHIRVersionR4 FHIRVersion = "R4"
)

// Values returns all known values for FHIRVersion. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FHIRVersion) Values() []FHIRVersion {
	return []FHIRVersion{
		"R4",
	}
}

type JobStatus string

// Enum values for JobStatus
const (
	JobStatusSubmitted           JobStatus = "SUBMITTED"
	JobStatusInProgress          JobStatus = "IN_PROGRESS"
	JobStatusCompletedWithErrors JobStatus = "COMPLETED_WITH_ERRORS"
	JobStatusCompleted           JobStatus = "COMPLETED"
	JobStatusFailed              JobStatus = "FAILED"
)

// Values returns all known values for JobStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (JobStatus) Values() []JobStatus {
	return []JobStatus{
		"SUBMITTED",
		"IN_PROGRESS",
		"COMPLETED_WITH_ERRORS",
		"COMPLETED",
		"FAILED",
	}
}

type PreloadDataType string

// Enum values for PreloadDataType
const (
	PreloadDataTypeSynthea PreloadDataType = "SYNTHEA"
)

// Values returns all known values for PreloadDataType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (PreloadDataType) Values() []PreloadDataType {
	return []PreloadDataType{
		"SYNTHEA",
	}
}
