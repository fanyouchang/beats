// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

type ContextKeyTypeEnum string

// Enum values for ContextKeyTypeEnum
const (
	ContextKeyTypeEnumString      ContextKeyTypeEnum = "string"
	ContextKeyTypeEnumStringList  ContextKeyTypeEnum = "stringList"
	ContextKeyTypeEnumNumeric     ContextKeyTypeEnum = "numeric"
	ContextKeyTypeEnumNumericList ContextKeyTypeEnum = "numericList"
	ContextKeyTypeEnumBoolean     ContextKeyTypeEnum = "boolean"
	ContextKeyTypeEnumBooleanList ContextKeyTypeEnum = "booleanList"
	ContextKeyTypeEnumIp          ContextKeyTypeEnum = "ip"
	ContextKeyTypeEnumIpList      ContextKeyTypeEnum = "ipList"
	ContextKeyTypeEnumBinary      ContextKeyTypeEnum = "binary"
	ContextKeyTypeEnumBinaryList  ContextKeyTypeEnum = "binaryList"
	ContextKeyTypeEnumDate        ContextKeyTypeEnum = "date"
	ContextKeyTypeEnumDateList    ContextKeyTypeEnum = "dateList"
)

func (enum ContextKeyTypeEnum) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ContextKeyTypeEnum) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DeletionTaskStatusType string

// Enum values for DeletionTaskStatusType
const (
	DeletionTaskStatusTypeSucceeded  DeletionTaskStatusType = "SUCCEEDED"
	DeletionTaskStatusTypeInProgress DeletionTaskStatusType = "IN_PROGRESS"
	DeletionTaskStatusTypeFailed     DeletionTaskStatusType = "FAILED"
	DeletionTaskStatusTypeNotStarted DeletionTaskStatusType = "NOT_STARTED"
)

func (enum DeletionTaskStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DeletionTaskStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EntityType string

// Enum values for EntityType
const (
	EntityTypeUser               EntityType = "User"
	EntityTypeRole               EntityType = "Role"
	EntityTypeGroup              EntityType = "Group"
	EntityTypeLocalManagedPolicy EntityType = "LocalManagedPolicy"
	EntityTypeAwsmanagedPolicy   EntityType = "AWSManagedPolicy"
)

func (enum EntityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EntityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PermissionsBoundaryAttachmentType string

// Enum values for PermissionsBoundaryAttachmentType
const (
	PermissionsBoundaryAttachmentTypePermissionsBoundaryPolicy PermissionsBoundaryAttachmentType = "PermissionsBoundaryPolicy"
)

func (enum PermissionsBoundaryAttachmentType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PermissionsBoundaryAttachmentType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyEvaluationDecisionType string

// Enum values for PolicyEvaluationDecisionType
const (
	PolicyEvaluationDecisionTypeAllowed      PolicyEvaluationDecisionType = "allowed"
	PolicyEvaluationDecisionTypeExplicitDeny PolicyEvaluationDecisionType = "explicitDeny"
	PolicyEvaluationDecisionTypeImplicitDeny PolicyEvaluationDecisionType = "implicitDeny"
)

func (enum PolicyEvaluationDecisionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyEvaluationDecisionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicySourceType string

// Enum values for PolicySourceType
const (
	PolicySourceTypeUser        PolicySourceType = "user"
	PolicySourceTypeGroup       PolicySourceType = "group"
	PolicySourceTypeRole        PolicySourceType = "role"
	PolicySourceTypeAwsManaged  PolicySourceType = "aws-managed"
	PolicySourceTypeUserManaged PolicySourceType = "user-managed"
	PolicySourceTypeResource    PolicySourceType = "resource"
	PolicySourceTypeNone        PolicySourceType = "none"
)

func (enum PolicySourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicySourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

// The policy usage type that indicates whether the policy is used as a permissions
// policy or as the permissions boundary for an entity.
//
// For more information about permissions boundaries, see Permissions Boundaries
// for IAM Identities (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_boundaries.html)
// in the IAM User Guide.
type PolicyUsageType string

// Enum values for PolicyUsageType
const (
	PolicyUsageTypePermissionsPolicy   PolicyUsageType = "PermissionsPolicy"
	PolicyUsageTypePermissionsBoundary PolicyUsageType = "PermissionsBoundary"
)

func (enum PolicyUsageType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyUsageType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReportFormatType string

// Enum values for ReportFormatType
const (
	ReportFormatTypeTextCsv ReportFormatType = "text/csv"
)

func (enum ReportFormatType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReportFormatType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ReportStateType string

// Enum values for ReportStateType
const (
	ReportStateTypeStarted    ReportStateType = "STARTED"
	ReportStateTypeInprogress ReportStateType = "INPROGRESS"
	ReportStateTypeComplete   ReportStateType = "COMPLETE"
)

func (enum ReportStateType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ReportStateType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AssignmentStatusType string

// Enum values for AssignmentStatusType
const (
	AssignmentStatusTypeAssigned   AssignmentStatusType = "Assigned"
	AssignmentStatusTypeUnassigned AssignmentStatusType = "Unassigned"
	AssignmentStatusTypeAny        AssignmentStatusType = "Any"
)

func (enum AssignmentStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AssignmentStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EncodingType string

// Enum values for EncodingType
const (
	EncodingTypeSsh EncodingType = "SSH"
	EncodingTypePem EncodingType = "PEM"
)

func (enum EncodingType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EncodingType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GlobalEndpointTokenVersion string

// Enum values for GlobalEndpointTokenVersion
const (
	GlobalEndpointTokenVersionV1token GlobalEndpointTokenVersion = "v1Token"
	GlobalEndpointTokenVersionV2token GlobalEndpointTokenVersion = "v2Token"
)

func (enum GlobalEndpointTokenVersion) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GlobalEndpointTokenVersion) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type JobStatusType string

// Enum values for JobStatusType
const (
	JobStatusTypeInProgress JobStatusType = "IN_PROGRESS"
	JobStatusTypeCompleted  JobStatusType = "COMPLETED"
	JobStatusTypeFailed     JobStatusType = "FAILED"
)

func (enum JobStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum JobStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyOwnerEntityType string

// Enum values for PolicyOwnerEntityType
const (
	PolicyOwnerEntityTypeUser  PolicyOwnerEntityType = "USER"
	PolicyOwnerEntityTypeRole  PolicyOwnerEntityType = "ROLE"
	PolicyOwnerEntityTypeGroup PolicyOwnerEntityType = "GROUP"
)

func (enum PolicyOwnerEntityType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyOwnerEntityType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyScopeType string

// Enum values for PolicyScopeType
const (
	PolicyScopeTypeAll   PolicyScopeType = "All"
	PolicyScopeTypeAws   PolicyScopeType = "AWS"
	PolicyScopeTypeLocal PolicyScopeType = "Local"
)

func (enum PolicyScopeType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyScopeType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyType string

// Enum values for PolicyType
const (
	PolicyTypeInline  PolicyType = "INLINE"
	PolicyTypeManaged PolicyType = "MANAGED"
)

func (enum PolicyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type StatusType string

// Enum values for StatusType
const (
	StatusTypeActive   StatusType = "Active"
	StatusTypeInactive StatusType = "Inactive"
)

func (enum StatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum StatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SummaryKeyType string

// Enum values for SummaryKeyType
const (
	SummaryKeyTypeUsers                             SummaryKeyType = "Users"
	SummaryKeyTypeUsersQuota                        SummaryKeyType = "UsersQuota"
	SummaryKeyTypeGroups                            SummaryKeyType = "Groups"
	SummaryKeyTypeGroupsQuota                       SummaryKeyType = "GroupsQuota"
	SummaryKeyTypeServerCertificates                SummaryKeyType = "ServerCertificates"
	SummaryKeyTypeServerCertificatesQuota           SummaryKeyType = "ServerCertificatesQuota"
	SummaryKeyTypeUserPolicySizeQuota               SummaryKeyType = "UserPolicySizeQuota"
	SummaryKeyTypeGroupPolicySizeQuota              SummaryKeyType = "GroupPolicySizeQuota"
	SummaryKeyTypeGroupsPerUserQuota                SummaryKeyType = "GroupsPerUserQuota"
	SummaryKeyTypeSigningCertificatesPerUserQuota   SummaryKeyType = "SigningCertificatesPerUserQuota"
	SummaryKeyTypeAccessKeysPerUserQuota            SummaryKeyType = "AccessKeysPerUserQuota"
	SummaryKeyTypeMfadevices                        SummaryKeyType = "MFADevices"
	SummaryKeyTypeMfadevicesInUse                   SummaryKeyType = "MFADevicesInUse"
	SummaryKeyTypeAccountMfaenabled                 SummaryKeyType = "AccountMFAEnabled"
	SummaryKeyTypeAccountAccessKeysPresent          SummaryKeyType = "AccountAccessKeysPresent"
	SummaryKeyTypeAccountSigningCertificatesPresent SummaryKeyType = "AccountSigningCertificatesPresent"
	SummaryKeyTypeAttachedPoliciesPerGroupQuota     SummaryKeyType = "AttachedPoliciesPerGroupQuota"
	SummaryKeyTypeAttachedPoliciesPerRoleQuota      SummaryKeyType = "AttachedPoliciesPerRoleQuota"
	SummaryKeyTypeAttachedPoliciesPerUserQuota      SummaryKeyType = "AttachedPoliciesPerUserQuota"
	SummaryKeyTypePolicies                          SummaryKeyType = "Policies"
	SummaryKeyTypePoliciesQuota                     SummaryKeyType = "PoliciesQuota"
	SummaryKeyTypePolicySizeQuota                   SummaryKeyType = "PolicySizeQuota"
	SummaryKeyTypePolicyVersionsInUse               SummaryKeyType = "PolicyVersionsInUse"
	SummaryKeyTypePolicyVersionsInUseQuota          SummaryKeyType = "PolicyVersionsInUseQuota"
	SummaryKeyTypeVersionsPerPolicyQuota            SummaryKeyType = "VersionsPerPolicyQuota"
	SummaryKeyTypeGlobalEndpointTokenVersion        SummaryKeyType = "GlobalEndpointTokenVersion"
)

func (enum SummaryKeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SummaryKeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}