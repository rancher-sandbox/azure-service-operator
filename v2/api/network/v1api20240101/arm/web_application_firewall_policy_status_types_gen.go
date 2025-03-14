// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type WebApplicationFirewallPolicy_STATUS struct {
	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the web application firewall policy.
	Properties *WebApplicationFirewallPolicyPropertiesFormat_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Defines web application firewall policy properties.
type WebApplicationFirewallPolicyPropertiesFormat_STATUS struct {
	// ApplicationGateways: A collection of references to application gateways.
	ApplicationGateways []ApplicationGateway_STATUS_ApplicationGatewayWebApplicationFirewallPolicy_SubResourceEmbedded `json:"applicationGateways,omitempty"`

	// CustomRules: The custom rules inside the policy.
	CustomRules []WebApplicationFirewallCustomRule_STATUS `json:"customRules,omitempty"`

	// HttpListeners: A collection of references to application gateway http listeners.
	HttpListeners []SubResource_STATUS `json:"httpListeners,omitempty"`

	// ManagedRules: Describes the managedRules structure.
	ManagedRules *ManagedRulesDefinition_STATUS `json:"managedRules,omitempty"`

	// PathBasedRules: A collection of references to application gateway path rules.
	PathBasedRules []SubResource_STATUS `json:"pathBasedRules,omitempty"`

	// PolicySettings: The PolicySettings for policy.
	PolicySettings *PolicySettings_STATUS `json:"policySettings,omitempty"`

	// ProvisioningState: The provisioning state of the web application firewall policy resource.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ResourceState: Resource status of the policy.
	ResourceState *WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS `json:"resourceState,omitempty"`
}

// Application gateway resource.
type ApplicationGateway_STATUS_ApplicationGatewayWebApplicationFirewallPolicy_SubResourceEmbedded struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Allow to exclude some variable satisfy the condition for the WAF check.
type ManagedRulesDefinition_STATUS struct {
	// Exclusions: The Exclusions that are applied on the policy.
	Exclusions []OwaspCrsExclusionEntry_STATUS `json:"exclusions,omitempty"`

	// ManagedRuleSets: The managed rule sets that are associated with the policy.
	ManagedRuleSets []ManagedRuleSet_STATUS `json:"managedRuleSets,omitempty"`
}

// Defines contents of a web application firewall global configuration.
type PolicySettings_STATUS struct {
	// CustomBlockResponseBody: If the action type is block, customer can override the response body. The body must be
	// specified in base64 encoding.
	CustomBlockResponseBody *string `json:"customBlockResponseBody,omitempty"`

	// CustomBlockResponseStatusCode: If the action type is block, customer can override the response status code.
	CustomBlockResponseStatusCode *int `json:"customBlockResponseStatusCode,omitempty"`

	// FileUploadEnforcement: Whether allow WAF to enforce file upload limits.
	FileUploadEnforcement *bool `json:"fileUploadEnforcement,omitempty"`

	// FileUploadLimitInMb: Maximum file upload size in Mb for WAF.
	FileUploadLimitInMb *int `json:"fileUploadLimitInMb,omitempty"`

	// JsChallengeCookieExpirationInMins: Web Application Firewall JavaScript Challenge Cookie Expiration time in minutes.
	JsChallengeCookieExpirationInMins *int `json:"jsChallengeCookieExpirationInMins,omitempty"`

	// LogScrubbing: To scrub sensitive log fields
	LogScrubbing *PolicySettings_LogScrubbing_STATUS `json:"logScrubbing,omitempty"`

	// MaxRequestBodySizeInKb: Maximum request body size in Kb for WAF.
	MaxRequestBodySizeInKb *int `json:"maxRequestBodySizeInKb,omitempty"`

	// Mode: The mode of the policy.
	Mode *PolicySettings_Mode_STATUS `json:"mode,omitempty"`

	// RequestBodyCheck: Whether to allow WAF to check request Body.
	RequestBodyCheck *bool `json:"requestBodyCheck,omitempty"`

	// RequestBodyEnforcement: Whether allow WAF to enforce request body limits.
	RequestBodyEnforcement *bool `json:"requestBodyEnforcement,omitempty"`

	// RequestBodyInspectLimitInKB: Max inspection limit in KB for request body inspection for WAF.
	RequestBodyInspectLimitInKB *int `json:"requestBodyInspectLimitInKB,omitempty"`

	// State: The state of the policy.
	State *PolicySettings_State_STATUS `json:"state,omitempty"`
}

// Reference to another subresource.
type SubResource_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

// Defines contents of a web application rule.
type WebApplicationFirewallCustomRule_STATUS struct {
	// Action: Type of Actions.
	Action *WebApplicationFirewallCustomRule_Action_STATUS `json:"action,omitempty"`

	// Etag: A unique read-only string that changes whenever the resource is updated.
	Etag *string `json:"etag,omitempty"`

	// GroupByUserSession: List of user session identifier group by clauses.
	GroupByUserSession []GroupByUserSession_STATUS `json:"groupByUserSession,omitempty"`

	// MatchConditions: List of match conditions.
	MatchConditions []MatchCondition_STATUS `json:"matchConditions,omitempty"`

	// Name: The name of the resource that is unique within a policy. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Priority: Priority of the rule. Rules with a lower value will be evaluated before rules with a higher value.
	Priority *int `json:"priority,omitempty"`

	// RateLimitDuration: Duration over which Rate Limit policy will be applied. Applies only when ruleType is RateLimitRule.
	RateLimitDuration *WebApplicationFirewallCustomRule_RateLimitDuration_STATUS `json:"rateLimitDuration,omitempty"`

	// RateLimitThreshold: Rate Limit threshold to apply in case ruleType is RateLimitRule. Must be greater than or equal to 1
	RateLimitThreshold *int `json:"rateLimitThreshold,omitempty"`

	// RuleType: The rule type.
	RuleType *WebApplicationFirewallCustomRule_RuleType_STATUS `json:"ruleType,omitempty"`

	// State: Describes if the custom rule is in enabled or disabled state. Defaults to Enabled if not specified.
	State *WebApplicationFirewallCustomRule_State_STATUS `json:"state,omitempty"`
}

type WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS string

const (
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Creating  = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Creating")
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Deleting  = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Deleting")
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Disabled  = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Disabled")
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Disabling = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Disabling")
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Enabled   = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Enabled")
	WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Enabling  = WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS("Enabling")
)

// Mapping from string to WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS
var webApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Values = map[string]WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS{
	"creating":  WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Creating,
	"deleting":  WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Deleting,
	"disabled":  WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Disabled,
	"disabling": WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Disabling,
	"enabled":   WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Enabled,
	"enabling":  WebApplicationFirewallPolicyPropertiesFormat_ResourceState_STATUS_Enabling,
}

// Define user session identifier group by clauses.
type GroupByUserSession_STATUS struct {
	// GroupByVariables: List of group by clause variables.
	GroupByVariables []GroupByVariable_STATUS `json:"groupByVariables,omitempty"`
}

// Defines a managed rule set.
type ManagedRuleSet_STATUS struct {
	// RuleGroupOverrides: Defines the rule group overrides to apply to the rule set.
	RuleGroupOverrides []ManagedRuleGroupOverride_STATUS `json:"ruleGroupOverrides,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

// Define match conditions.
type MatchCondition_STATUS struct {
	// MatchValues: Match value.
	MatchValues []string `json:"matchValues,omitempty"`

	// MatchVariables: List of match variables.
	MatchVariables []MatchVariable_STATUS `json:"matchVariables,omitempty"`

	// NegationConditon: Whether this is negate condition or not.
	NegationConditon *bool `json:"negationConditon,omitempty"`

	// Operator: The operator to be matched.
	Operator *MatchCondition_Operator_STATUS `json:"operator,omitempty"`

	// Transforms: List of transforms.
	Transforms []Transform_STATUS `json:"transforms,omitempty"`
}

// Allow to exclude some variable satisfy the condition for the WAF check.
type OwaspCrsExclusionEntry_STATUS struct {
	// ExclusionManagedRuleSets: The managed rule sets that are associated with the exclusion.
	ExclusionManagedRuleSets []ExclusionManagedRuleSet_STATUS `json:"exclusionManagedRuleSets,omitempty"`

	// MatchVariable: The variable to be excluded.
	MatchVariable *OwaspCrsExclusionEntry_MatchVariable_STATUS `json:"matchVariable,omitempty"`

	// Selector: When matchVariable is a collection, operator used to specify which elements in the collection this exclusion
	// applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: When matchVariable is a collection, operate on the selector to specify which elements in the
	// collection this exclusion applies to.
	SelectorMatchOperator *OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS `json:"selectorMatchOperator,omitempty"`
}

type PolicySettings_LogScrubbing_STATUS struct {
	// ScrubbingRules: The rules that are applied to the logs for scrubbing.
	ScrubbingRules []WebApplicationFirewallScrubbingRules_STATUS `json:"scrubbingRules,omitempty"`

	// State: State of the log scrubbing config. Default value is Enabled.
	State *PolicySettings_LogScrubbing_State_STATUS `json:"state,omitempty"`
}

type PolicySettings_Mode_STATUS string

const (
	PolicySettings_Mode_STATUS_Detection  = PolicySettings_Mode_STATUS("Detection")
	PolicySettings_Mode_STATUS_Prevention = PolicySettings_Mode_STATUS("Prevention")
)

// Mapping from string to PolicySettings_Mode_STATUS
var policySettings_Mode_STATUS_Values = map[string]PolicySettings_Mode_STATUS{
	"detection":  PolicySettings_Mode_STATUS_Detection,
	"prevention": PolicySettings_Mode_STATUS_Prevention,
}

type PolicySettings_State_STATUS string

const (
	PolicySettings_State_STATUS_Disabled = PolicySettings_State_STATUS("Disabled")
	PolicySettings_State_STATUS_Enabled  = PolicySettings_State_STATUS("Enabled")
)

// Mapping from string to PolicySettings_State_STATUS
var policySettings_State_STATUS_Values = map[string]PolicySettings_State_STATUS{
	"disabled": PolicySettings_State_STATUS_Disabled,
	"enabled":  PolicySettings_State_STATUS_Enabled,
}

type WebApplicationFirewallCustomRule_Action_STATUS string

const (
	WebApplicationFirewallCustomRule_Action_STATUS_Allow       = WebApplicationFirewallCustomRule_Action_STATUS("Allow")
	WebApplicationFirewallCustomRule_Action_STATUS_Block       = WebApplicationFirewallCustomRule_Action_STATUS("Block")
	WebApplicationFirewallCustomRule_Action_STATUS_JSChallenge = WebApplicationFirewallCustomRule_Action_STATUS("JSChallenge")
	WebApplicationFirewallCustomRule_Action_STATUS_Log         = WebApplicationFirewallCustomRule_Action_STATUS("Log")
)

// Mapping from string to WebApplicationFirewallCustomRule_Action_STATUS
var webApplicationFirewallCustomRule_Action_STATUS_Values = map[string]WebApplicationFirewallCustomRule_Action_STATUS{
	"allow":       WebApplicationFirewallCustomRule_Action_STATUS_Allow,
	"block":       WebApplicationFirewallCustomRule_Action_STATUS_Block,
	"jschallenge": WebApplicationFirewallCustomRule_Action_STATUS_JSChallenge,
	"log":         WebApplicationFirewallCustomRule_Action_STATUS_Log,
}

type WebApplicationFirewallCustomRule_RateLimitDuration_STATUS string

const (
	WebApplicationFirewallCustomRule_RateLimitDuration_STATUS_FiveMins = WebApplicationFirewallCustomRule_RateLimitDuration_STATUS("FiveMins")
	WebApplicationFirewallCustomRule_RateLimitDuration_STATUS_OneMin   = WebApplicationFirewallCustomRule_RateLimitDuration_STATUS("OneMin")
)

// Mapping from string to WebApplicationFirewallCustomRule_RateLimitDuration_STATUS
var webApplicationFirewallCustomRule_RateLimitDuration_STATUS_Values = map[string]WebApplicationFirewallCustomRule_RateLimitDuration_STATUS{
	"fivemins": WebApplicationFirewallCustomRule_RateLimitDuration_STATUS_FiveMins,
	"onemin":   WebApplicationFirewallCustomRule_RateLimitDuration_STATUS_OneMin,
}

type WebApplicationFirewallCustomRule_RuleType_STATUS string

const (
	WebApplicationFirewallCustomRule_RuleType_STATUS_Invalid       = WebApplicationFirewallCustomRule_RuleType_STATUS("Invalid")
	WebApplicationFirewallCustomRule_RuleType_STATUS_MatchRule     = WebApplicationFirewallCustomRule_RuleType_STATUS("MatchRule")
	WebApplicationFirewallCustomRule_RuleType_STATUS_RateLimitRule = WebApplicationFirewallCustomRule_RuleType_STATUS("RateLimitRule")
)

// Mapping from string to WebApplicationFirewallCustomRule_RuleType_STATUS
var webApplicationFirewallCustomRule_RuleType_STATUS_Values = map[string]WebApplicationFirewallCustomRule_RuleType_STATUS{
	"invalid":       WebApplicationFirewallCustomRule_RuleType_STATUS_Invalid,
	"matchrule":     WebApplicationFirewallCustomRule_RuleType_STATUS_MatchRule,
	"ratelimitrule": WebApplicationFirewallCustomRule_RuleType_STATUS_RateLimitRule,
}

type WebApplicationFirewallCustomRule_State_STATUS string

const (
	WebApplicationFirewallCustomRule_State_STATUS_Disabled = WebApplicationFirewallCustomRule_State_STATUS("Disabled")
	WebApplicationFirewallCustomRule_State_STATUS_Enabled  = WebApplicationFirewallCustomRule_State_STATUS("Enabled")
)

// Mapping from string to WebApplicationFirewallCustomRule_State_STATUS
var webApplicationFirewallCustomRule_State_STATUS_Values = map[string]WebApplicationFirewallCustomRule_State_STATUS{
	"disabled": WebApplicationFirewallCustomRule_State_STATUS_Disabled,
	"enabled":  WebApplicationFirewallCustomRule_State_STATUS_Enabled,
}

// Defines a managed rule set for Exclusions.
type ExclusionManagedRuleSet_STATUS struct {
	// RuleGroups: Defines the rule groups to apply to the rule set.
	RuleGroups []ExclusionManagedRuleGroup_STATUS `json:"ruleGroups,omitempty"`

	// RuleSetType: Defines the rule set type to use.
	RuleSetType *string `json:"ruleSetType,omitempty"`

	// RuleSetVersion: Defines the version of the rule set to use.
	RuleSetVersion *string `json:"ruleSetVersion,omitempty"`
}

// Define user session group by clause variables.
type GroupByVariable_STATUS struct {
	// VariableName: User Session clause variable.
	VariableName *GroupByVariable_VariableName_STATUS `json:"variableName,omitempty"`
}

// Defines a managed rule group override setting.
type ManagedRuleGroupOverride_STATUS struct {
	// RuleGroupName: The managed rule group to override.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be disabled. If none specified, all rules in the group will be disabled.
	Rules []ManagedRuleOverride_STATUS `json:"rules,omitempty"`
}

type MatchCondition_Operator_STATUS string

const (
	MatchCondition_Operator_STATUS_Any                = MatchCondition_Operator_STATUS("Any")
	MatchCondition_Operator_STATUS_BeginsWith         = MatchCondition_Operator_STATUS("BeginsWith")
	MatchCondition_Operator_STATUS_Contains           = MatchCondition_Operator_STATUS("Contains")
	MatchCondition_Operator_STATUS_EndsWith           = MatchCondition_Operator_STATUS("EndsWith")
	MatchCondition_Operator_STATUS_Equal              = MatchCondition_Operator_STATUS("Equal")
	MatchCondition_Operator_STATUS_GeoMatch           = MatchCondition_Operator_STATUS("GeoMatch")
	MatchCondition_Operator_STATUS_GreaterThan        = MatchCondition_Operator_STATUS("GreaterThan")
	MatchCondition_Operator_STATUS_GreaterThanOrEqual = MatchCondition_Operator_STATUS("GreaterThanOrEqual")
	MatchCondition_Operator_STATUS_IPMatch            = MatchCondition_Operator_STATUS("IPMatch")
	MatchCondition_Operator_STATUS_LessThan           = MatchCondition_Operator_STATUS("LessThan")
	MatchCondition_Operator_STATUS_LessThanOrEqual    = MatchCondition_Operator_STATUS("LessThanOrEqual")
	MatchCondition_Operator_STATUS_Regex              = MatchCondition_Operator_STATUS("Regex")
)

// Mapping from string to MatchCondition_Operator_STATUS
var matchCondition_Operator_STATUS_Values = map[string]MatchCondition_Operator_STATUS{
	"any":                MatchCondition_Operator_STATUS_Any,
	"beginswith":         MatchCondition_Operator_STATUS_BeginsWith,
	"contains":           MatchCondition_Operator_STATUS_Contains,
	"endswith":           MatchCondition_Operator_STATUS_EndsWith,
	"equal":              MatchCondition_Operator_STATUS_Equal,
	"geomatch":           MatchCondition_Operator_STATUS_GeoMatch,
	"greaterthan":        MatchCondition_Operator_STATUS_GreaterThan,
	"greaterthanorequal": MatchCondition_Operator_STATUS_GreaterThanOrEqual,
	"ipmatch":            MatchCondition_Operator_STATUS_IPMatch,
	"lessthan":           MatchCondition_Operator_STATUS_LessThan,
	"lessthanorequal":    MatchCondition_Operator_STATUS_LessThanOrEqual,
	"regex":              MatchCondition_Operator_STATUS_Regex,
}

// Define match variables.
type MatchVariable_STATUS struct {
	// Selector: The selector of match variable.
	Selector *string `json:"selector,omitempty"`

	// VariableName: Match Variable.
	VariableName *MatchVariable_VariableName_STATUS `json:"variableName,omitempty"`
}

type OwaspCrsExclusionEntry_MatchVariable_STATUS string

const (
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgKeys      = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestArgKeys")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgNames     = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestArgNames")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgValues    = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestArgValues")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieKeys   = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestCookieKeys")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieNames  = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestCookieNames")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieValues = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestCookieValues")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderKeys   = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestHeaderKeys")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderNames  = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestHeaderNames")
	OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderValues = OwaspCrsExclusionEntry_MatchVariable_STATUS("RequestHeaderValues")
)

// Mapping from string to OwaspCrsExclusionEntry_MatchVariable_STATUS
var owaspCrsExclusionEntry_MatchVariable_STATUS_Values = map[string]OwaspCrsExclusionEntry_MatchVariable_STATUS{
	"requestargkeys":      OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgKeys,
	"requestargnames":     OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgNames,
	"requestargvalues":    OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestArgValues,
	"requestcookiekeys":   OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieKeys,
	"requestcookienames":  OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieNames,
	"requestcookievalues": OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestCookieValues,
	"requestheaderkeys":   OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderKeys,
	"requestheadernames":  OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderNames,
	"requestheadervalues": OwaspCrsExclusionEntry_MatchVariable_STATUS_RequestHeaderValues,
}

type OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS string

const (
	OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_Contains   = OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS("Contains")
	OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_EndsWith   = OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS("EndsWith")
	OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_Equals     = OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS("Equals")
	OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_EqualsAny  = OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS("EqualsAny")
	OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_StartsWith = OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS("StartsWith")
)

// Mapping from string to OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS
var owaspCrsExclusionEntry_SelectorMatchOperator_STATUS_Values = map[string]OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS{
	"contains":   OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_Contains,
	"endswith":   OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_EndsWith,
	"equals":     OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_Equals,
	"equalsany":  OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_EqualsAny,
	"startswith": OwaspCrsExclusionEntry_SelectorMatchOperator_STATUS_StartsWith,
}

type PolicySettings_LogScrubbing_State_STATUS string

const (
	PolicySettings_LogScrubbing_State_STATUS_Disabled = PolicySettings_LogScrubbing_State_STATUS("Disabled")
	PolicySettings_LogScrubbing_State_STATUS_Enabled  = PolicySettings_LogScrubbing_State_STATUS("Enabled")
)

// Mapping from string to PolicySettings_LogScrubbing_State_STATUS
var policySettings_LogScrubbing_State_STATUS_Values = map[string]PolicySettings_LogScrubbing_State_STATUS{
	"disabled": PolicySettings_LogScrubbing_State_STATUS_Disabled,
	"enabled":  PolicySettings_LogScrubbing_State_STATUS_Enabled,
}

// Transforms applied before matching.
type Transform_STATUS string

const (
	Transform_STATUS_HtmlEntityDecode = Transform_STATUS("HtmlEntityDecode")
	Transform_STATUS_Lowercase        = Transform_STATUS("Lowercase")
	Transform_STATUS_RemoveNulls      = Transform_STATUS("RemoveNulls")
	Transform_STATUS_Trim             = Transform_STATUS("Trim")
	Transform_STATUS_Uppercase        = Transform_STATUS("Uppercase")
	Transform_STATUS_UrlDecode        = Transform_STATUS("UrlDecode")
	Transform_STATUS_UrlEncode        = Transform_STATUS("UrlEncode")
)

// Mapping from string to Transform_STATUS
var transform_STATUS_Values = map[string]Transform_STATUS{
	"htmlentitydecode": Transform_STATUS_HtmlEntityDecode,
	"lowercase":        Transform_STATUS_Lowercase,
	"removenulls":      Transform_STATUS_RemoveNulls,
	"trim":             Transform_STATUS_Trim,
	"uppercase":        Transform_STATUS_Uppercase,
	"urldecode":        Transform_STATUS_UrlDecode,
	"urlencode":        Transform_STATUS_UrlEncode,
}

// Allow certain variables to be scrubbed on WAF logs
type WebApplicationFirewallScrubbingRules_STATUS struct {
	// MatchVariable: The variable to be scrubbed from the logs.
	MatchVariable *WebApplicationFirewallScrubbingRules_MatchVariable_STATUS `json:"matchVariable,omitempty"`

	// Selector: When matchVariable is a collection, operator used to specify which elements in the collection this rule
	// applies to.
	Selector *string `json:"selector,omitempty"`

	// SelectorMatchOperator: When matchVariable is a collection, operate on the selector to specify which elements in the
	// collection this rule applies to.
	SelectorMatchOperator *WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS `json:"selectorMatchOperator,omitempty"`

	// State: Defines the state of log scrubbing rule. Default value is Enabled.
	State *WebApplicationFirewallScrubbingRules_State_STATUS `json:"state,omitempty"`
}

// Defines a managed rule group to use for exclusion.
type ExclusionManagedRuleGroup_STATUS struct {
	// RuleGroupName: The managed rule group for exclusion.
	RuleGroupName *string `json:"ruleGroupName,omitempty"`

	// Rules: List of rules that will be excluded. If none specified, all rules in the group will be excluded.
	Rules []ExclusionManagedRule_STATUS `json:"rules,omitempty"`
}

type GroupByVariable_VariableName_STATUS string

const (
	GroupByVariable_VariableName_STATUS_ClientAddr  = GroupByVariable_VariableName_STATUS("ClientAddr")
	GroupByVariable_VariableName_STATUS_GeoLocation = GroupByVariable_VariableName_STATUS("GeoLocation")
	GroupByVariable_VariableName_STATUS_None        = GroupByVariable_VariableName_STATUS("None")
)

// Mapping from string to GroupByVariable_VariableName_STATUS
var groupByVariable_VariableName_STATUS_Values = map[string]GroupByVariable_VariableName_STATUS{
	"clientaddr":  GroupByVariable_VariableName_STATUS_ClientAddr,
	"geolocation": GroupByVariable_VariableName_STATUS_GeoLocation,
	"none":        GroupByVariable_VariableName_STATUS_None,
}

// Defines a managed rule group override setting.
type ManagedRuleOverride_STATUS struct {
	// Action: Describes the override action to be applied when rule matches.
	Action *ActionType_STATUS `json:"action,omitempty"`

	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`

	// State: The state of the managed rule. Defaults to Disabled if not specified.
	State *ManagedRuleOverride_State_STATUS `json:"state,omitempty"`
}

type MatchVariable_VariableName_STATUS string

const (
	MatchVariable_VariableName_STATUS_PostArgs       = MatchVariable_VariableName_STATUS("PostArgs")
	MatchVariable_VariableName_STATUS_QueryString    = MatchVariable_VariableName_STATUS("QueryString")
	MatchVariable_VariableName_STATUS_RemoteAddr     = MatchVariable_VariableName_STATUS("RemoteAddr")
	MatchVariable_VariableName_STATUS_RequestBody    = MatchVariable_VariableName_STATUS("RequestBody")
	MatchVariable_VariableName_STATUS_RequestCookies = MatchVariable_VariableName_STATUS("RequestCookies")
	MatchVariable_VariableName_STATUS_RequestHeaders = MatchVariable_VariableName_STATUS("RequestHeaders")
	MatchVariable_VariableName_STATUS_RequestMethod  = MatchVariable_VariableName_STATUS("RequestMethod")
	MatchVariable_VariableName_STATUS_RequestUri     = MatchVariable_VariableName_STATUS("RequestUri")
)

// Mapping from string to MatchVariable_VariableName_STATUS
var matchVariable_VariableName_STATUS_Values = map[string]MatchVariable_VariableName_STATUS{
	"postargs":       MatchVariable_VariableName_STATUS_PostArgs,
	"querystring":    MatchVariable_VariableName_STATUS_QueryString,
	"remoteaddr":     MatchVariable_VariableName_STATUS_RemoteAddr,
	"requestbody":    MatchVariable_VariableName_STATUS_RequestBody,
	"requestcookies": MatchVariable_VariableName_STATUS_RequestCookies,
	"requestheaders": MatchVariable_VariableName_STATUS_RequestHeaders,
	"requestmethod":  MatchVariable_VariableName_STATUS_RequestMethod,
	"requesturi":     MatchVariable_VariableName_STATUS_RequestUri,
}

type WebApplicationFirewallScrubbingRules_MatchVariable_STATUS string

const (
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestArgNames     = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestArgNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestCookieNames  = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestCookieNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestHeaderNames  = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestHeaderNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestIPAddress    = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestIPAddress")
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestJSONArgNames = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestJSONArgNames")
	WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestPostArgNames = WebApplicationFirewallScrubbingRules_MatchVariable_STATUS("RequestPostArgNames")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_MatchVariable_STATUS
var webApplicationFirewallScrubbingRules_MatchVariable_STATUS_Values = map[string]WebApplicationFirewallScrubbingRules_MatchVariable_STATUS{
	"requestargnames":     WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestArgNames,
	"requestcookienames":  WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestCookieNames,
	"requestheadernames":  WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestHeaderNames,
	"requestipaddress":    WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestIPAddress,
	"requestjsonargnames": WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestJSONArgNames,
	"requestpostargnames": WebApplicationFirewallScrubbingRules_MatchVariable_STATUS_RequestPostArgNames,
}

type WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS string

const (
	WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS_Equals    = WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS("Equals")
	WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS_EqualsAny = WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS("EqualsAny")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS
var webApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS_Values = map[string]WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS{
	"equals":    WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS_Equals,
	"equalsany": WebApplicationFirewallScrubbingRules_SelectorMatchOperator_STATUS_EqualsAny,
}

type WebApplicationFirewallScrubbingRules_State_STATUS string

const (
	WebApplicationFirewallScrubbingRules_State_STATUS_Disabled = WebApplicationFirewallScrubbingRules_State_STATUS("Disabled")
	WebApplicationFirewallScrubbingRules_State_STATUS_Enabled  = WebApplicationFirewallScrubbingRules_State_STATUS("Enabled")
)

// Mapping from string to WebApplicationFirewallScrubbingRules_State_STATUS
var webApplicationFirewallScrubbingRules_State_STATUS_Values = map[string]WebApplicationFirewallScrubbingRules_State_STATUS{
	"disabled": WebApplicationFirewallScrubbingRules_State_STATUS_Disabled,
	"enabled":  WebApplicationFirewallScrubbingRules_State_STATUS_Enabled,
}

// Defines the action to take on rule match.
type ActionType_STATUS string

const (
	ActionType_STATUS_Allow          = ActionType_STATUS("Allow")
	ActionType_STATUS_AnomalyScoring = ActionType_STATUS("AnomalyScoring")
	ActionType_STATUS_Block          = ActionType_STATUS("Block")
	ActionType_STATUS_JSChallenge    = ActionType_STATUS("JSChallenge")
	ActionType_STATUS_Log            = ActionType_STATUS("Log")
)

// Mapping from string to ActionType_STATUS
var actionType_STATUS_Values = map[string]ActionType_STATUS{
	"allow":          ActionType_STATUS_Allow,
	"anomalyscoring": ActionType_STATUS_AnomalyScoring,
	"block":          ActionType_STATUS_Block,
	"jschallenge":    ActionType_STATUS_JSChallenge,
	"log":            ActionType_STATUS_Log,
}

// Defines a managed rule to use for exclusion.
type ExclusionManagedRule_STATUS struct {
	// RuleId: Identifier for the managed rule.
	RuleId *string `json:"ruleId,omitempty"`
}

type ManagedRuleOverride_State_STATUS string

const (
	ManagedRuleOverride_State_STATUS_Disabled = ManagedRuleOverride_State_STATUS("Disabled")
	ManagedRuleOverride_State_STATUS_Enabled  = ManagedRuleOverride_State_STATUS("Enabled")
)

// Mapping from string to ManagedRuleOverride_State_STATUS
var managedRuleOverride_State_STATUS_Values = map[string]ManagedRuleOverride_State_STATUS{
	"disabled": ManagedRuleOverride_State_STATUS_Disabled,
	"enabled":  ManagedRuleOverride_State_STATUS_Enabled,
}
