// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

type ComparisonOperator string

// Enum values for ComparisonOperator
const (
	ComparisonOperatorEq ComparisonOperator = "EQ"
	ComparisonOperatorNe ComparisonOperator = "NE"
	ComparisonOperatorLe ComparisonOperator = "LE"
	ComparisonOperatorLt ComparisonOperator = "LT"
	ComparisonOperatorGe ComparisonOperator = "GE"
	ComparisonOperatorGt ComparisonOperator = "GT"
)

func (enum ComparisonOperator) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ComparisonOperator) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type CountryCode string

// Enum values for CountryCode
const (
	CountryCodeAf CountryCode = "AF"
	CountryCodeAx CountryCode = "AX"
	CountryCodeAl CountryCode = "AL"
	CountryCodeDz CountryCode = "DZ"
	CountryCodeAs CountryCode = "AS"
	CountryCodeAd CountryCode = "AD"
	CountryCodeAo CountryCode = "AO"
	CountryCodeAi CountryCode = "AI"
	CountryCodeAq CountryCode = "AQ"
	CountryCodeAg CountryCode = "AG"
	CountryCodeAr CountryCode = "AR"
	CountryCodeAm CountryCode = "AM"
	CountryCodeAw CountryCode = "AW"
	CountryCodeAu CountryCode = "AU"
	CountryCodeAt CountryCode = "AT"
	CountryCodeAz CountryCode = "AZ"
	CountryCodeBs CountryCode = "BS"
	CountryCodeBh CountryCode = "BH"
	CountryCodeBd CountryCode = "BD"
	CountryCodeBb CountryCode = "BB"
	CountryCodeBy CountryCode = "BY"
	CountryCodeBe CountryCode = "BE"
	CountryCodeBz CountryCode = "BZ"
	CountryCodeBj CountryCode = "BJ"
	CountryCodeBm CountryCode = "BM"
	CountryCodeBt CountryCode = "BT"
	CountryCodeBo CountryCode = "BO"
	CountryCodeBq CountryCode = "BQ"
	CountryCodeBa CountryCode = "BA"
	CountryCodeBw CountryCode = "BW"
	CountryCodeBv CountryCode = "BV"
	CountryCodeBr CountryCode = "BR"
	CountryCodeIo CountryCode = "IO"
	CountryCodeBn CountryCode = "BN"
	CountryCodeBg CountryCode = "BG"
	CountryCodeBf CountryCode = "BF"
	CountryCodeBi CountryCode = "BI"
	CountryCodeKh CountryCode = "KH"
	CountryCodeCm CountryCode = "CM"
	CountryCodeCa CountryCode = "CA"
	CountryCodeCv CountryCode = "CV"
	CountryCodeKy CountryCode = "KY"
	CountryCodeCf CountryCode = "CF"
	CountryCodeTd CountryCode = "TD"
	CountryCodeCl CountryCode = "CL"
	CountryCodeCn CountryCode = "CN"
	CountryCodeCx CountryCode = "CX"
	CountryCodeCc CountryCode = "CC"
	CountryCodeCo CountryCode = "CO"
	CountryCodeKm CountryCode = "KM"
	CountryCodeCg CountryCode = "CG"
	CountryCodeCd CountryCode = "CD"
	CountryCodeCk CountryCode = "CK"
	CountryCodeCr CountryCode = "CR"
	CountryCodeCi CountryCode = "CI"
	CountryCodeHr CountryCode = "HR"
	CountryCodeCu CountryCode = "CU"
	CountryCodeCw CountryCode = "CW"
	CountryCodeCy CountryCode = "CY"
	CountryCodeCz CountryCode = "CZ"
	CountryCodeDk CountryCode = "DK"
	CountryCodeDj CountryCode = "DJ"
	CountryCodeDm CountryCode = "DM"
	CountryCodeDo CountryCode = "DO"
	CountryCodeEc CountryCode = "EC"
	CountryCodeEg CountryCode = "EG"
	CountryCodeSv CountryCode = "SV"
	CountryCodeGq CountryCode = "GQ"
	CountryCodeEr CountryCode = "ER"
	CountryCodeEe CountryCode = "EE"
	CountryCodeEt CountryCode = "ET"
	CountryCodeFk CountryCode = "FK"
	CountryCodeFo CountryCode = "FO"
	CountryCodeFj CountryCode = "FJ"
	CountryCodeFi CountryCode = "FI"
	CountryCodeFr CountryCode = "FR"
	CountryCodeGf CountryCode = "GF"
	CountryCodePf CountryCode = "PF"
	CountryCodeTf CountryCode = "TF"
	CountryCodeGa CountryCode = "GA"
	CountryCodeGm CountryCode = "GM"
	CountryCodeGe CountryCode = "GE"
	CountryCodeDe CountryCode = "DE"
	CountryCodeGh CountryCode = "GH"
	CountryCodeGi CountryCode = "GI"
	CountryCodeGr CountryCode = "GR"
	CountryCodeGl CountryCode = "GL"
	CountryCodeGd CountryCode = "GD"
	CountryCodeGp CountryCode = "GP"
	CountryCodeGu CountryCode = "GU"
	CountryCodeGt CountryCode = "GT"
	CountryCodeGg CountryCode = "GG"
	CountryCodeGn CountryCode = "GN"
	CountryCodeGw CountryCode = "GW"
	CountryCodeGy CountryCode = "GY"
	CountryCodeHt CountryCode = "HT"
	CountryCodeHm CountryCode = "HM"
	CountryCodeVa CountryCode = "VA"
	CountryCodeHn CountryCode = "HN"
	CountryCodeHk CountryCode = "HK"
	CountryCodeHu CountryCode = "HU"
	CountryCodeIs CountryCode = "IS"
	CountryCodeIn CountryCode = "IN"
	CountryCodeId CountryCode = "ID"
	CountryCodeIr CountryCode = "IR"
	CountryCodeIq CountryCode = "IQ"
	CountryCodeIe CountryCode = "IE"
	CountryCodeIm CountryCode = "IM"
	CountryCodeIl CountryCode = "IL"
	CountryCodeIt CountryCode = "IT"
	CountryCodeJm CountryCode = "JM"
	CountryCodeJp CountryCode = "JP"
	CountryCodeJe CountryCode = "JE"
	CountryCodeJo CountryCode = "JO"
	CountryCodeKz CountryCode = "KZ"
	CountryCodeKe CountryCode = "KE"
	CountryCodeKi CountryCode = "KI"
	CountryCodeKp CountryCode = "KP"
	CountryCodeKr CountryCode = "KR"
	CountryCodeKw CountryCode = "KW"
	CountryCodeKg CountryCode = "KG"
	CountryCodeLa CountryCode = "LA"
	CountryCodeLv CountryCode = "LV"
	CountryCodeLb CountryCode = "LB"
	CountryCodeLs CountryCode = "LS"
	CountryCodeLr CountryCode = "LR"
	CountryCodeLy CountryCode = "LY"
	CountryCodeLi CountryCode = "LI"
	CountryCodeLt CountryCode = "LT"
	CountryCodeLu CountryCode = "LU"
	CountryCodeMo CountryCode = "MO"
	CountryCodeMk CountryCode = "MK"
	CountryCodeMg CountryCode = "MG"
	CountryCodeMw CountryCode = "MW"
	CountryCodeMy CountryCode = "MY"
	CountryCodeMv CountryCode = "MV"
	CountryCodeMl CountryCode = "ML"
	CountryCodeMt CountryCode = "MT"
	CountryCodeMh CountryCode = "MH"
	CountryCodeMq CountryCode = "MQ"
	CountryCodeMr CountryCode = "MR"
	CountryCodeMu CountryCode = "MU"
	CountryCodeYt CountryCode = "YT"
	CountryCodeMx CountryCode = "MX"
	CountryCodeFm CountryCode = "FM"
	CountryCodeMd CountryCode = "MD"
	CountryCodeMc CountryCode = "MC"
	CountryCodeMn CountryCode = "MN"
	CountryCodeMe CountryCode = "ME"
	CountryCodeMs CountryCode = "MS"
	CountryCodeMa CountryCode = "MA"
	CountryCodeMz CountryCode = "MZ"
	CountryCodeMm CountryCode = "MM"
	CountryCodeNa CountryCode = "NA"
	CountryCodeNr CountryCode = "NR"
	CountryCodeNp CountryCode = "NP"
	CountryCodeNl CountryCode = "NL"
	CountryCodeNc CountryCode = "NC"
	CountryCodeNz CountryCode = "NZ"
	CountryCodeNi CountryCode = "NI"
	CountryCodeNe CountryCode = "NE"
	CountryCodeNg CountryCode = "NG"
	CountryCodeNu CountryCode = "NU"
	CountryCodeNf CountryCode = "NF"
	CountryCodeMp CountryCode = "MP"
	CountryCodeNo CountryCode = "NO"
	CountryCodeOm CountryCode = "OM"
	CountryCodePk CountryCode = "PK"
	CountryCodePw CountryCode = "PW"
	CountryCodePs CountryCode = "PS"
	CountryCodePa CountryCode = "PA"
	CountryCodePg CountryCode = "PG"
	CountryCodePy CountryCode = "PY"
	CountryCodePe CountryCode = "PE"
	CountryCodePh CountryCode = "PH"
	CountryCodePn CountryCode = "PN"
	CountryCodePl CountryCode = "PL"
	CountryCodePt CountryCode = "PT"
	CountryCodePr CountryCode = "PR"
	CountryCodeQa CountryCode = "QA"
	CountryCodeRe CountryCode = "RE"
	CountryCodeRo CountryCode = "RO"
	CountryCodeRu CountryCode = "RU"
	CountryCodeRw CountryCode = "RW"
	CountryCodeBl CountryCode = "BL"
	CountryCodeSh CountryCode = "SH"
	CountryCodeKn CountryCode = "KN"
	CountryCodeLc CountryCode = "LC"
	CountryCodeMf CountryCode = "MF"
	CountryCodePm CountryCode = "PM"
	CountryCodeVc CountryCode = "VC"
	CountryCodeWs CountryCode = "WS"
	CountryCodeSm CountryCode = "SM"
	CountryCodeSt CountryCode = "ST"
	CountryCodeSa CountryCode = "SA"
	CountryCodeSn CountryCode = "SN"
	CountryCodeRs CountryCode = "RS"
	CountryCodeSc CountryCode = "SC"
	CountryCodeSl CountryCode = "SL"
	CountryCodeSg CountryCode = "SG"
	CountryCodeSx CountryCode = "SX"
	CountryCodeSk CountryCode = "SK"
	CountryCodeSi CountryCode = "SI"
	CountryCodeSb CountryCode = "SB"
	CountryCodeSo CountryCode = "SO"
	CountryCodeZa CountryCode = "ZA"
	CountryCodeGs CountryCode = "GS"
	CountryCodeSs CountryCode = "SS"
	CountryCodeEs CountryCode = "ES"
	CountryCodeLk CountryCode = "LK"
	CountryCodeSd CountryCode = "SD"
	CountryCodeSr CountryCode = "SR"
	CountryCodeSj CountryCode = "SJ"
	CountryCodeSz CountryCode = "SZ"
	CountryCodeSe CountryCode = "SE"
	CountryCodeCh CountryCode = "CH"
	CountryCodeSy CountryCode = "SY"
	CountryCodeTw CountryCode = "TW"
	CountryCodeTj CountryCode = "TJ"
	CountryCodeTz CountryCode = "TZ"
	CountryCodeTh CountryCode = "TH"
	CountryCodeTl CountryCode = "TL"
	CountryCodeTg CountryCode = "TG"
	CountryCodeTk CountryCode = "TK"
	CountryCodeTo CountryCode = "TO"
	CountryCodeTt CountryCode = "TT"
	CountryCodeTn CountryCode = "TN"
	CountryCodeTr CountryCode = "TR"
	CountryCodeTm CountryCode = "TM"
	CountryCodeTc CountryCode = "TC"
	CountryCodeTv CountryCode = "TV"
	CountryCodeUg CountryCode = "UG"
	CountryCodeUa CountryCode = "UA"
	CountryCodeAe CountryCode = "AE"
	CountryCodeGb CountryCode = "GB"
	CountryCodeUs CountryCode = "US"
	CountryCodeUm CountryCode = "UM"
	CountryCodeUy CountryCode = "UY"
	CountryCodeUz CountryCode = "UZ"
	CountryCodeVu CountryCode = "VU"
	CountryCodeVe CountryCode = "VE"
	CountryCodeVn CountryCode = "VN"
	CountryCodeVg CountryCode = "VG"
	CountryCodeVi CountryCode = "VI"
	CountryCodeWf CountryCode = "WF"
	CountryCodeEh CountryCode = "EH"
	CountryCodeYe CountryCode = "YE"
	CountryCodeZm CountryCode = "ZM"
	CountryCodeZw CountryCode = "ZW"
)

func (enum CountryCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum CountryCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type FallbackBehavior string

// Enum values for FallbackBehavior
const (
	FallbackBehaviorMatch   FallbackBehavior = "MATCH"
	FallbackBehaviorNoMatch FallbackBehavior = "NO_MATCH"
)

func (enum FallbackBehavior) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum FallbackBehavior) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ForwardedIPPosition string

// Enum values for ForwardedIPPosition
const (
	ForwardedIPPositionFirst ForwardedIPPosition = "FIRST"
	ForwardedIPPositionLast  ForwardedIPPosition = "LAST"
	ForwardedIPPositionAny   ForwardedIPPosition = "ANY"
)

func (enum ForwardedIPPosition) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ForwardedIPPosition) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type IPAddressVersion string

// Enum values for IPAddressVersion
const (
	IPAddressVersionIpv4 IPAddressVersion = "IPV4"
	IPAddressVersionIpv6 IPAddressVersion = "IPV6"
)

func (enum IPAddressVersion) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum IPAddressVersion) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ParameterExceptionField string

// Enum values for ParameterExceptionField
const (
	ParameterExceptionFieldWebAcl                         ParameterExceptionField = "WEB_ACL"
	ParameterExceptionFieldRuleGroup                      ParameterExceptionField = "RULE_GROUP"
	ParameterExceptionFieldRegexPatternSet                ParameterExceptionField = "REGEX_PATTERN_SET"
	ParameterExceptionFieldIpSet                          ParameterExceptionField = "IP_SET"
	ParameterExceptionFieldManagedRuleSet                 ParameterExceptionField = "MANAGED_RULE_SET"
	ParameterExceptionFieldRule                           ParameterExceptionField = "RULE"
	ParameterExceptionFieldExcludedRule                   ParameterExceptionField = "EXCLUDED_RULE"
	ParameterExceptionFieldStatement                      ParameterExceptionField = "STATEMENT"
	ParameterExceptionFieldByteMatchStatement             ParameterExceptionField = "BYTE_MATCH_STATEMENT"
	ParameterExceptionFieldSqliMatchStatement             ParameterExceptionField = "SQLI_MATCH_STATEMENT"
	ParameterExceptionFieldXssMatchStatement              ParameterExceptionField = "XSS_MATCH_STATEMENT"
	ParameterExceptionFieldSizeConstraintStatement        ParameterExceptionField = "SIZE_CONSTRAINT_STATEMENT"
	ParameterExceptionFieldGeoMatchStatement              ParameterExceptionField = "GEO_MATCH_STATEMENT"
	ParameterExceptionFieldRateBasedStatement             ParameterExceptionField = "RATE_BASED_STATEMENT"
	ParameterExceptionFieldRuleGroupReferenceStatement    ParameterExceptionField = "RULE_GROUP_REFERENCE_STATEMENT"
	ParameterExceptionFieldRegexPatternReferenceStatement ParameterExceptionField = "REGEX_PATTERN_REFERENCE_STATEMENT"
	ParameterExceptionFieldIpSetReferenceStatement        ParameterExceptionField = "IP_SET_REFERENCE_STATEMENT"
	ParameterExceptionFieldManagedRuleSetStatement        ParameterExceptionField = "MANAGED_RULE_SET_STATEMENT"
	ParameterExceptionFieldAndStatement                   ParameterExceptionField = "AND_STATEMENT"
	ParameterExceptionFieldOrStatement                    ParameterExceptionField = "OR_STATEMENT"
	ParameterExceptionFieldNotStatement                   ParameterExceptionField = "NOT_STATEMENT"
	ParameterExceptionFieldIpAddress                      ParameterExceptionField = "IP_ADDRESS"
	ParameterExceptionFieldIpAddressVersion               ParameterExceptionField = "IP_ADDRESS_VERSION"
	ParameterExceptionFieldFieldToMatch                   ParameterExceptionField = "FIELD_TO_MATCH"
	ParameterExceptionFieldTextTransformation             ParameterExceptionField = "TEXT_TRANSFORMATION"
	ParameterExceptionFieldSingleQueryArgument            ParameterExceptionField = "SINGLE_QUERY_ARGUMENT"
	ParameterExceptionFieldSingleHeader                   ParameterExceptionField = "SINGLE_HEADER"
	ParameterExceptionFieldDefaultAction                  ParameterExceptionField = "DEFAULT_ACTION"
	ParameterExceptionFieldRuleAction                     ParameterExceptionField = "RULE_ACTION"
	ParameterExceptionFieldEntityLimit                    ParameterExceptionField = "ENTITY_LIMIT"
	ParameterExceptionFieldOverrideAction                 ParameterExceptionField = "OVERRIDE_ACTION"
	ParameterExceptionFieldScopeValue                     ParameterExceptionField = "SCOPE_VALUE"
	ParameterExceptionFieldResourceArn                    ParameterExceptionField = "RESOURCE_ARN"
	ParameterExceptionFieldResourceType                   ParameterExceptionField = "RESOURCE_TYPE"
	ParameterExceptionFieldTags                           ParameterExceptionField = "TAGS"
	ParameterExceptionFieldTagKeys                        ParameterExceptionField = "TAG_KEYS"
	ParameterExceptionFieldMetricName                     ParameterExceptionField = "METRIC_NAME"
	ParameterExceptionFieldFirewallManagerStatement       ParameterExceptionField = "FIREWALL_MANAGER_STATEMENT"
	ParameterExceptionFieldFallbackBehavior               ParameterExceptionField = "FALLBACK_BEHAVIOR"
	ParameterExceptionFieldPosition                       ParameterExceptionField = "POSITION"
	ParameterExceptionFieldForwardedIpConfig              ParameterExceptionField = "FORWARDED_IP_CONFIG"
	ParameterExceptionFieldIpSetForwardedIpConfig         ParameterExceptionField = "IP_SET_FORWARDED_IP_CONFIG"
	ParameterExceptionFieldHeaderName                     ParameterExceptionField = "HEADER_NAME"
)

func (enum ParameterExceptionField) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ParameterExceptionField) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PositionalConstraint string

// Enum values for PositionalConstraint
const (
	PositionalConstraintExactly      PositionalConstraint = "EXACTLY"
	PositionalConstraintStartsWith   PositionalConstraint = "STARTS_WITH"
	PositionalConstraintEndsWith     PositionalConstraint = "ENDS_WITH"
	PositionalConstraintContains     PositionalConstraint = "CONTAINS"
	PositionalConstraintContainsWord PositionalConstraint = "CONTAINS_WORD"
)

func (enum PositionalConstraint) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PositionalConstraint) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RateBasedStatementAggregateKeyType string

// Enum values for RateBasedStatementAggregateKeyType
const (
	RateBasedStatementAggregateKeyTypeIp          RateBasedStatementAggregateKeyType = "IP"
	RateBasedStatementAggregateKeyTypeForwardedIp RateBasedStatementAggregateKeyType = "FORWARDED_IP"
)

func (enum RateBasedStatementAggregateKeyType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RateBasedStatementAggregateKeyType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeApplicationLoadBalancer ResourceType = "APPLICATION_LOAD_BALANCER"
	ResourceTypeApiGateway              ResourceType = "API_GATEWAY"
)

func (enum ResourceType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ResourceType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type Scope string

// Enum values for Scope
const (
	ScopeCloudfront Scope = "CLOUDFRONT"
	ScopeRegional   Scope = "REGIONAL"
)

func (enum Scope) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum Scope) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TextTransformationType string

// Enum values for TextTransformationType
const (
	TextTransformationTypeNone               TextTransformationType = "NONE"
	TextTransformationTypeCompressWhiteSpace TextTransformationType = "COMPRESS_WHITE_SPACE"
	TextTransformationTypeHtmlEntityDecode   TextTransformationType = "HTML_ENTITY_DECODE"
	TextTransformationTypeLowercase          TextTransformationType = "LOWERCASE"
	TextTransformationTypeCmdLine            TextTransformationType = "CMD_LINE"
	TextTransformationTypeUrlDecode          TextTransformationType = "URL_DECODE"
)

func (enum TextTransformationType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TextTransformationType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
