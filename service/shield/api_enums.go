// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

type AttackLayer string

// Enum values for AttackLayer
const (
	AttackLayerNetwork     AttackLayer = "NETWORK"
	AttackLayerApplication AttackLayer = "APPLICATION"
)

func (enum AttackLayer) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AttackLayer) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AttackPropertyIdentifier string

// Enum values for AttackPropertyIdentifier
const (
	AttackPropertyIdentifierDestinationUrl             AttackPropertyIdentifier = "DESTINATION_URL"
	AttackPropertyIdentifierReferrer                   AttackPropertyIdentifier = "REFERRER"
	AttackPropertyIdentifierSourceAsn                  AttackPropertyIdentifier = "SOURCE_ASN"
	AttackPropertyIdentifierSourceCountry              AttackPropertyIdentifier = "SOURCE_COUNTRY"
	AttackPropertyIdentifierSourceIpAddress            AttackPropertyIdentifier = "SOURCE_IP_ADDRESS"
	AttackPropertyIdentifierSourceUserAgent            AttackPropertyIdentifier = "SOURCE_USER_AGENT"
	AttackPropertyIdentifierWordpressPingbackReflector AttackPropertyIdentifier = "WORDPRESS_PINGBACK_REFLECTOR"
	AttackPropertyIdentifierWordpressPingbackSource    AttackPropertyIdentifier = "WORDPRESS_PINGBACK_SOURCE"
)

func (enum AttackPropertyIdentifier) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AttackPropertyIdentifier) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type AutoRenew string

// Enum values for AutoRenew
const (
	AutoRenewEnabled  AutoRenew = "ENABLED"
	AutoRenewDisabled AutoRenew = "DISABLED"
)

func (enum AutoRenew) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum AutoRenew) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ProactiveEngagementStatus string

// Enum values for ProactiveEngagementStatus
const (
	ProactiveEngagementStatusEnabled  ProactiveEngagementStatus = "ENABLED"
	ProactiveEngagementStatusDisabled ProactiveEngagementStatus = "DISABLED"
	ProactiveEngagementStatusPending  ProactiveEngagementStatus = "PENDING"
)

func (enum ProactiveEngagementStatus) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ProactiveEngagementStatus) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SubResourceType string

// Enum values for SubResourceType
const (
	SubResourceTypeIp  SubResourceType = "IP"
	SubResourceTypeUrl SubResourceType = "URL"
)

func (enum SubResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SubResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type SubscriptionState string

// Enum values for SubscriptionState
const (
	SubscriptionStateActive   SubscriptionState = "ACTIVE"
	SubscriptionStateInactive SubscriptionState = "INACTIVE"
)

func (enum SubscriptionState) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum SubscriptionState) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Unit string

// Enum values for Unit
const (
	UnitBits     Unit = "BITS"
	UnitBytes    Unit = "BYTES"
	UnitPackets  Unit = "PACKETS"
	UnitRequests Unit = "REQUESTS"
)

func (enum Unit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Unit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
