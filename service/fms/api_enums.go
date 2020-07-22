// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fms

type AccountRoleStatus string

// Enum values for AccountRoleStatus
const (
	AccountRoleStatusReady           AccountRoleStatus = "READY"
	AccountRoleStatusCreating        AccountRoleStatus = "CREATING"
	AccountRoleStatusPendingDeletion AccountRoleStatus = "PENDING_DELETION"
	AccountRoleStatusDeleting        AccountRoleStatus = "DELETING"
	AccountRoleStatusDeleted         AccountRoleStatus = "DELETED"
)

func (enum AccountRoleStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AccountRoleStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CustomerPolicyScopeIdType string

// Enum values for CustomerPolicyScopeIdType
const (
	CustomerPolicyScopeIdTypeAccount CustomerPolicyScopeIdType = "ACCOUNT"
	CustomerPolicyScopeIdTypeOrgUnit CustomerPolicyScopeIdType = "ORG_UNIT"
)

func (enum CustomerPolicyScopeIdType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CustomerPolicyScopeIdType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type DependentServiceName string

// Enum values for DependentServiceName
const (
	DependentServiceNameAwsconfig         DependentServiceName = "AWSCONFIG"
	DependentServiceNameAwswaf            DependentServiceName = "AWSWAF"
	DependentServiceNameAwsshieldAdvanced DependentServiceName = "AWSSHIELD_ADVANCED"
	DependentServiceNameAwsvpc            DependentServiceName = "AWSVPC"
)

func (enum DependentServiceName) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DependentServiceName) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PolicyComplianceStatusType string

// Enum values for PolicyComplianceStatusType
const (
	PolicyComplianceStatusTypeCompliant    PolicyComplianceStatusType = "COMPLIANT"
	PolicyComplianceStatusTypeNonCompliant PolicyComplianceStatusType = "NON_COMPLIANT"
)

func (enum PolicyComplianceStatusType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PolicyComplianceStatusType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RemediationActionType string

// Enum values for RemediationActionType
const (
	RemediationActionTypeRemove RemediationActionType = "REMOVE"
	RemediationActionTypeModify RemediationActionType = "MODIFY"
)

func (enum RemediationActionType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RemediationActionType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SecurityServiceType string

// Enum values for SecurityServiceType
const (
	SecurityServiceTypeWaf                        SecurityServiceType = "WAF"
	SecurityServiceTypeWafv2                      SecurityServiceType = "WAFV2"
	SecurityServiceTypeShieldAdvanced             SecurityServiceType = "SHIELD_ADVANCED"
	SecurityServiceTypeSecurityGroupsCommon       SecurityServiceType = "SECURITY_GROUPS_COMMON"
	SecurityServiceTypeSecurityGroupsContentAudit SecurityServiceType = "SECURITY_GROUPS_CONTENT_AUDIT"
	SecurityServiceTypeSecurityGroupsUsageAudit   SecurityServiceType = "SECURITY_GROUPS_USAGE_AUDIT"
)

func (enum SecurityServiceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SecurityServiceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ViolationReason string

// Enum values for ViolationReason
const (
	ViolationReasonWebAclMissingRuleGroup                  ViolationReason = "WEB_ACL_MISSING_RULE_GROUP"
	ViolationReasonResourceMissingWebAcl                   ViolationReason = "RESOURCE_MISSING_WEB_ACL"
	ViolationReasonResourceIncorrectWebAcl                 ViolationReason = "RESOURCE_INCORRECT_WEB_ACL"
	ViolationReasonResourceMissingShieldProtection         ViolationReason = "RESOURCE_MISSING_SHIELD_PROTECTION"
	ViolationReasonResourceMissingWebAclOrShieldProtection ViolationReason = "RESOURCE_MISSING_WEB_ACL_OR_SHIELD_PROTECTION"
	ViolationReasonResourceMissingSecurityGroup            ViolationReason = "RESOURCE_MISSING_SECURITY_GROUP"
	ViolationReasonResourceViolatesAuditSecurityGroup      ViolationReason = "RESOURCE_VIOLATES_AUDIT_SECURITY_GROUP"
	ViolationReasonSecurityGroupUnused                     ViolationReason = "SECURITY_GROUP_UNUSED"
	ViolationReasonSecurityGroupRedundant                  ViolationReason = "SECURITY_GROUP_REDUNDANT"
)

func (enum ViolationReason) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ViolationReason) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
