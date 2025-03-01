---
title: insights.azure.com/v1api20240101preview
---
<h2 id="insights.azure.com/v1api20240101preview">insights.azure.com/v1api20240101preview</h2>
<div>
<p>Package v1api20240101preview contains API Schema definitions for the insights v1api20240101preview API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="insights.azure.com/v1api20240101preview.APIVersion">APIVersion
(<code>string</code> alias)</h3>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;2024-01-01-preview&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Actions">Actions
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
<p>Actions to invoke when the alert fires.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actionGroupsReferences</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
[]genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ActionGroupsReferences: Action Group resource Ids to invoke when the alert fires.</p>
</td>
</tr>
<tr>
<td>
<code>actionProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ActionProperties: The properties of an action properties.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: The properties of an alert payload.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Actions_STATUS">Actions_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
<p>Actions to invoke when the alert fires.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actionGroups</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>ActionGroups: Action Group resource Ids to invoke when the alert fires.</p>
</td>
</tr>
<tr>
<td>
<code>actionProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>ActionProperties: The properties of an action properties.</p>
</td>
</tr>
<tr>
<td>
<code>customProperties</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>CustomProperties: The properties of an alert payload.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition">Condition
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria">ScheduledQueryRuleCriteria</a>)
</p>
<div>
<p>A condition of the scheduled query rule.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>alertSensitivity</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlertSensitivity: The extent of deviation required to trigger an alert. Allowed values are &lsquo;Low&rsquo;, &lsquo;Medium&rsquo; and &lsquo;High&rsquo;.
This will affect how tight the threshold is to the metric series pattern. Relevant and required only for dynamic
threshold rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>criterionType</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_CriterionType">
Condition_CriterionType
</a>
</em>
</td>
<td>
<p>CriterionType: Specifies the type of threshold criteria</p>
</td>
</tr>
<tr>
<td>
<code>dimensions</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Dimension">
[]Dimension
</a>
</em>
</td>
<td>
<p>Dimensions: List of Dimensions conditions</p>
</td>
</tr>
<tr>
<td>
<code>failingPeriods</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_FailingPeriods">
Condition_FailingPeriods
</a>
</em>
</td>
<td>
<p>FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
alert. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>ignoreDataBefore</code><br/>
<em>
string
</em>
</td>
<td>
<p>IgnoreDataBefore: Use this option to set the date from which to start learning the metric historical data and calculate
the dynamic thresholds (in ISO8601 format). Relevant only for dynamic threshold rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>metricMeasureColumn</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricMeasureColumn: The column containing the metric measure number. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>metricName</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricName: The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.</p>
</td>
</tr>
<tr>
<td>
<code>operator</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_Operator">
Condition_Operator
</a>
</em>
</td>
<td>
<p>Operator: The criteria operator. Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Log query alert</p>
</td>
</tr>
<tr>
<td>
<code>resourceIdColumnReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ResourceIdColumnReference: The column containing the resource id. The content of the column must be a uri formatted as
resource id. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>threshold</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Threshold: the criteria threshold value that activates the alert. Relevant and required only for static threshold rules
of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>timeAggregation</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_TimeAggregation">
Condition_TimeAggregation
</a>
</em>
</td>
<td>
<p>TimeAggregation: Aggregation type. Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_CriterionType">Condition_CriterionType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition">Condition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;DynamicThresholdCriterion&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StaticThresholdCriterion&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_CriterionType_STATUS">Condition_CriterionType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;DynamicThresholdCriterion&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;StaticThresholdCriterion&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_FailingPeriods">Condition_FailingPeriods
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition">Condition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minFailingPeriodsToAlert</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
numberOfEvaluationPeriods. Default value is 1</p>
</td>
</tr>
<tr>
<td>
<code>numberOfEvaluationPeriods</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_FailingPeriods_STATUS">Condition_FailingPeriods_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>minFailingPeriodsToAlert</code><br/>
<em>
int
</em>
</td>
<td>
<p>MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
numberOfEvaluationPeriods. Default value is 1</p>
</td>
</tr>
<tr>
<td>
<code>numberOfEvaluationPeriods</code><br/>
<em>
int
</em>
</td>
<td>
<p>NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_Operator">Condition_Operator
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition">Condition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Equals&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterOrLessThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterThanOrEqual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LessThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LessThanOrEqual&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_Operator_STATUS">Condition_Operator_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Equals&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterOrLessThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;GreaterThanOrEqual&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LessThan&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LessThanOrEqual&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria_STATUS">ScheduledQueryRuleCriteria_STATUS</a>)
</p>
<div>
<p>A condition of the scheduled query rule.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>alertSensitivity</code><br/>
<em>
string
</em>
</td>
<td>
<p>AlertSensitivity: The extent of deviation required to trigger an alert. Allowed values are &lsquo;Low&rsquo;, &lsquo;Medium&rsquo; and &lsquo;High&rsquo;.
This will affect how tight the threshold is to the metric series pattern. Relevant and required only for dynamic
threshold rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>criterionType</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_CriterionType_STATUS">
Condition_CriterionType_STATUS
</a>
</em>
</td>
<td>
<p>CriterionType: Specifies the type of threshold criteria</p>
</td>
</tr>
<tr>
<td>
<code>dimensions</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Dimension_STATUS">
[]Dimension_STATUS
</a>
</em>
</td>
<td>
<p>Dimensions: List of Dimensions conditions</p>
</td>
</tr>
<tr>
<td>
<code>failingPeriods</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_FailingPeriods_STATUS">
Condition_FailingPeriods_STATUS
</a>
</em>
</td>
<td>
<p>FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
alert. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>ignoreDataBefore</code><br/>
<em>
string
</em>
</td>
<td>
<p>IgnoreDataBefore: Use this option to set the date from which to start learning the metric historical data and calculate
the dynamic thresholds (in ISO8601 format). Relevant only for dynamic threshold rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>metricMeasureColumn</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricMeasureColumn: The column containing the metric measure number. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>metricName</code><br/>
<em>
string
</em>
</td>
<td>
<p>MetricName: The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.</p>
</td>
</tr>
<tr>
<td>
<code>operator</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_Operator_STATUS">
Condition_Operator_STATUS
</a>
</em>
</td>
<td>
<p>Operator: The criteria operator. Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>query</code><br/>
<em>
string
</em>
</td>
<td>
<p>Query: Log query alert</p>
</td>
</tr>
<tr>
<td>
<code>resourceIdColumn</code><br/>
<em>
string
</em>
</td>
<td>
<p>ResourceIdColumn: The column containing the resource id. The content of the column must be a uri formatted as resource
id. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>threshold</code><br/>
<em>
float64
</em>
</td>
<td>
<p>Threshold: the criteria threshold value that activates the alert. Relevant and required only for static threshold rules
of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>timeAggregation</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_TimeAggregation_STATUS">
Condition_TimeAggregation_STATUS
</a>
</em>
</td>
<td>
<p>TimeAggregation: Aggregation type. Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_TimeAggregation">Condition_TimeAggregation
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition">Condition</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Average&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Count&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Maximum&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Minimum&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Total&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Condition_TimeAggregation_STATUS">Condition_TimeAggregation_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Average&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Count&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Maximum&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Minimum&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Total&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Dimension">Dimension
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition">Condition</a>)
</p>
<div>
<p>Dimension splitting and filtering definition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the dimension</p>
</td>
</tr>
<tr>
<td>
<code>operator</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Dimension_Operator">
Dimension_Operator
</a>
</em>
</td>
<td>
<p>Operator: Operator for dimension values</p>
</td>
</tr>
<tr>
<td>
<code>values</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Values: List of dimension values</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Dimension_Operator">Dimension_Operator
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Dimension">Dimension</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Exclude&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Include&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Dimension_Operator_STATUS">Dimension_Operator_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Dimension_STATUS">Dimension_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Exclude&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Include&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Dimension_STATUS">Dimension_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">Condition_STATUS</a>)
</p>
<div>
<p>Dimension splitting and filtering definition</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: Name of the dimension</p>
</td>
</tr>
<tr>
<td>
<code>operator</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Dimension_Operator_STATUS">
Dimension_Operator_STATUS
</a>
</em>
</td>
<td>
<p>Operator: Operator for dimension values</p>
</td>
</tr>
<tr>
<td>
<code>values</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Values: List of dimension values</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Identity">Identity
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
<p>Identity for the resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Identity_Type">
Identity_Type
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.UserAssignedIdentityDetails">
[]UserAssignedIdentityDetails
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the resource. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Identity_STATUS">Identity_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
<p>Identity for the resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal ID of resource identity.</p>
</td>
</tr>
<tr>
<td>
<code>tenantId</code><br/>
<em>
string
</em>
</td>
<td>
<p>TenantId: The tenant ID of resource.</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Identity_Type_STATUS">
Identity_Type_STATUS
</a>
</em>
</td>
<td>
<p>Type: Type of managed service identity.</p>
</td>
</tr>
<tr>
<td>
<code>userAssignedIdentities</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.UserIdentityProperties_STATUS">
map[string]./api/insights/v1api20240101preview.UserIdentityProperties_STATUS
</a>
</em>
</td>
<td>
<p>UserAssignedIdentities: The list of user identities associated with the resource. The user identity dictionary key
references will be ARM resource ids in the form:
&lsquo;/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.ManagedIdentity/&#x200b;userAssignedIdentities/&#x200b;{identityName}&rsquo;.</&#x200b;p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Identity_Type">Identity_Type
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Identity">Identity</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.Identity_Type_STATUS">Identity_Type_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Identity_STATUS">Identity_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;None&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;SystemAssigned&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;UserAssigned&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.RuleResolveConfiguration">RuleResolveConfiguration
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
<p>TBD. Relevant only for rules of the kind LogAlert.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoResolved</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoResolved: The flag that indicates whether or not to auto resolve a fired alert.</p>
</td>
</tr>
<tr>
<td>
<code>timeToResolve</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeToResolve: The duration a rule must evaluate as healthy before the fired alert is automatically resolved represented
in ISO 8601 duration format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.RuleResolveConfiguration_STATUS">RuleResolveConfiguration_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
<p>TBD. Relevant only for rules of the kind LogAlert.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>autoResolved</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoResolved: The flag that indicates whether or not to auto resolve a fired alert.</p>
</td>
</tr>
<tr>
<td>
<code>timeToResolve</code><br/>
<em>
string
</em>
</td>
<td>
<p>TimeToResolve: The duration a rule must evaluate as healthy before the fired alert is automatically resolved represented
in ISO 8601 duration format.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRule">ScheduledQueryRule
</h3>
<div>
<p>Generator information:
- Generated from: /monitor/resource-manager/Microsoft.Insights/preview/2024-01-01-preview/scheduledQueryRule_API.json
- ARM URI: /&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;Microsoft.Insights/&#x200b;scheduledQueryRules/&#x200b;{ruleName}</&#x200b;p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">
ScheduledQueryRule_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>actions</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Actions">
Actions
</a>
</em>
</td>
<td>
<p>Actions: Actions to invoke when the alert fires.</p>
</td>
</tr>
<tr>
<td>
<code>autoMitigate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoMitigate: The flag that indicates whether the alert should be automatically resolved or not. The default is true.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>checkWorkspaceAlertsStorageConfigured</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CheckWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule should be stored in
the customer&rsquo;s storage. The default is false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>criteria</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria">
ScheduledQueryRuleCriteria
</a>
</em>
</td>
<td>
<p>Criteria: The rule criteria that defines the conditions of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: The description of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: The display name of the alert rule</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: The flag which indicates whether this scheduled query rule is enabled. Value should be true or false</p>
</td>
</tr>
<tr>
<td>
<code>evaluationFrequency</code><br/>
<em>
string
</em>
</td>
<td>
<p>EvaluationFrequency: How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant
and required only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Kind_Spec">
ScheduledQueryRule_Kind_Spec
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of scheduled query rule. The default is LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>muteActionsDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>MuteActionsDuration: Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleOperatorSpec">
ScheduledQueryRuleOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
</td>
</tr>
<tr>
<td>
<code>overrideQueryTimeRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>OverrideQueryTimeRange: If specified then overrides the query time range (default is
WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>resolveConfiguration</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.RuleResolveConfiguration">
RuleResolveConfiguration
</a>
</em>
</td>
<td>
<p>ResolveConfiguration: Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>scopesReferences</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
[]genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ScopesReferences: The list of resource id&rsquo;s that this scheduled query rule is scoped to.</p>
</td>
</tr>
<tr>
<td>
<code>severity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleProperties_Severity">
ScheduledQueryRuleProperties_Severity
</a>
</em>
</td>
<td>
<p>Severity: Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only
for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>skipQueryValidation</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipQueryValidation: The flag which indicates whether the provided query should be validated or not. The default is
false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>targetResourceTypes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>TargetResourceTypes: List of resource type of the target resource(s) on which the alert is created/updated. For example
if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert
will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of
the kind LogAlert</p>
</td>
</tr>
<tr>
<td>
<code>windowSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>WindowSize: The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">
ScheduledQueryRule_STATUS
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria">ScheduledQueryRuleCriteria
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
<p>The rule criteria that defines the conditions of the scheduled query rule.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allOf</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition">
[]Condition
</a>
</em>
</td>
<td>
<p>AllOf: A list of conditions to evaluate against the specified scopes</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria_STATUS">ScheduledQueryRuleCriteria_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
<p>The rule criteria that defines the conditions of the scheduled query rule.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>allOf</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Condition_STATUS">
[]Condition_STATUS
</a>
</em>
</td>
<td>
<p>AllOf: A list of conditions to evaluate against the specified scopes</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRuleOperatorSpec">ScheduledQueryRuleOperatorSpec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
<p>Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>configMapExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).</p>
</td>
</tr>
<tr>
<td>
<code>secretExpressions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#DestinationExpression">
[]genruntime/core.DestinationExpression
</a>
</em>
</td>
<td>
<p>SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRuleProperties_Severity">ScheduledQueryRuleProperties_Severity
(<code>int</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>0</p></td>
<td></td>
</tr><tr><td><p>1</p></td>
<td></td>
</tr><tr><td><p>2</p></td>
<td></td>
</tr><tr><td><p>3</p></td>
<td></td>
</tr><tr><td><p>4</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRuleProperties_Severity_STATUS">ScheduledQueryRuleProperties_Severity_STATUS
(<code>int</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>0</p></td>
<td></td>
</tr><tr><td><p>1</p></td>
<td></td>
</tr><tr><td><p>2</p></td>
<td></td>
</tr><tr><td><p>3</p></td>
<td></td>
</tr><tr><td><p>4</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRule_Kind_STATUS">ScheduledQueryRule_Kind_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EventLogAlert&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAlert&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogToMetric&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRule_Kind_Spec">ScheduledQueryRule_Kind_Spec
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;EventLogAlert&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAlert&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogToMetric&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule">ScheduledQueryRule</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actions</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Actions_STATUS">
Actions_STATUS
</a>
</em>
</td>
<td>
<p>Actions: Actions to invoke when the alert fires.</p>
</td>
</tr>
<tr>
<td>
<code>autoMitigate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoMitigate: The flag that indicates whether the alert should be automatically resolved or not. The default is true.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>checkWorkspaceAlertsStorageConfigured</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CheckWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule should be stored in
the customer&rsquo;s storage. The default is false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>createdWithApiVersion</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedWithApiVersion: The api-version used when creating this alert rule</p>
</td>
</tr>
<tr>
<td>
<code>criteria</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria_STATUS">
ScheduledQueryRuleCriteria_STATUS
</a>
</em>
</td>
<td>
<p>Criteria: The rule criteria that defines the conditions of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: The description of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: The display name of the alert rule</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: The flag which indicates whether this scheduled query rule is enabled. Value should be true or false</p>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
<p>Etag: The etag field is <em>not</em> required. If it is provided in the response body, it must also be provided as a header per
the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource.
HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and
If-Range (section 14.27) header fields.</p>
</td>
</tr>
<tr>
<td>
<code>evaluationFrequency</code><br/>
<em>
string
</em>
</td>
<td>
<p>EvaluationFrequency: How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant
and required only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
<p>Id: Fully qualified resource ID for the resource. Ex -
/&#x200b;subscriptions/&#x200b;{subscriptionId}/&#x200b;resourceGroups/&#x200b;{resourceGroupName}/&#x200b;providers/&#x200b;{resourceProviderNamespace}/&#x200b;{resourceType}/&#x200b;{resourceName}</&#x200b;p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Identity_STATUS">
Identity_STATUS
</a>
</em>
</td>
<td>
<p>Identity: The identity of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>isLegacyLogAnalyticsRule</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsLegacyLogAnalyticsRule: True if alert rule is legacy Log Analytic rule</p>
</td>
</tr>
<tr>
<td>
<code>isWorkspaceAlertsStorageConfigured</code><br/>
<em>
bool
</em>
</td>
<td>
<p>IsWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule has been configured to be
stored in the customer&rsquo;s storage. The default is false.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Kind_STATUS">
ScheduledQueryRule_Kind_STATUS
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of scheduled query rule. The default is LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>muteActionsDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>MuteActionsDuration: Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
<p>Name: The name of the resource</p>
</td>
</tr>
<tr>
<td>
<code>overrideQueryTimeRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>OverrideQueryTimeRange: If specified then overrides the query time range (default is
WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>resolveConfiguration</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.RuleResolveConfiguration_STATUS">
RuleResolveConfiguration_STATUS
</a>
</em>
</td>
<td>
<p>ResolveConfiguration: Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>scopes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>Scopes: The list of resource id&rsquo;s that this scheduled query rule is scoped to.</p>
</td>
</tr>
<tr>
<td>
<code>severity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleProperties_Severity_STATUS">
ScheduledQueryRuleProperties_Severity_STATUS
</a>
</em>
</td>
<td>
<p>Severity: Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only
for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>skipQueryValidation</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipQueryValidation: The flag which indicates whether the provided query should be validated or not. The default is
false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>systemData</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.SystemData_STATUS">
SystemData_STATUS
</a>
</em>
</td>
<td>
<p>SystemData: SystemData of ScheduledQueryRule.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>targetResourceTypes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>TargetResourceTypes: List of resource type of the target resource(s) on which the alert is created/updated. For example
if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert
will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of
the kind LogAlert</p>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
<p>Type: The type of the resource. E.g. &ldquo;Microsoft.Compute/virtualMachines&rdquo; or &ldquo;Microsoft.Storage/storageAccounts&rdquo;</p>
</td>
</tr>
<tr>
<td>
<code>windowSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>WindowSize: The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.ScheduledQueryRule_Spec">ScheduledQueryRule_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule">ScheduledQueryRule</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>actions</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Actions">
Actions
</a>
</em>
</td>
<td>
<p>Actions: Actions to invoke when the alert fires.</p>
</td>
</tr>
<tr>
<td>
<code>autoMitigate</code><br/>
<em>
bool
</em>
</td>
<td>
<p>AutoMitigate: The flag that indicates whether the alert should be automatically resolved or not. The default is true.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>checkWorkspaceAlertsStorageConfigured</code><br/>
<em>
bool
</em>
</td>
<td>
<p>CheckWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule should be stored in
the customer&rsquo;s storage. The default is false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>criteria</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleCriteria">
ScheduledQueryRuleCriteria
</a>
</em>
</td>
<td>
<p>Criteria: The rule criteria that defines the conditions of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>description</code><br/>
<em>
string
</em>
</td>
<td>
<p>Description: The description of the scheduled query rule.</p>
</td>
</tr>
<tr>
<td>
<code>displayName</code><br/>
<em>
string
</em>
</td>
<td>
<p>DisplayName: The display name of the alert rule</p>
</td>
</tr>
<tr>
<td>
<code>enabled</code><br/>
<em>
bool
</em>
</td>
<td>
<p>Enabled: The flag which indicates whether this scheduled query rule is enabled. Value should be true or false</p>
</td>
</tr>
<tr>
<td>
<code>evaluationFrequency</code><br/>
<em>
string
</em>
</td>
<td>
<p>EvaluationFrequency: How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant
and required only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>identity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.Identity">
Identity
</a>
</em>
</td>
<td>
<p>Identity: The identity of the resource.</p>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_Kind_Spec">
ScheduledQueryRule_Kind_Spec
</a>
</em>
</td>
<td>
<p>Kind: Indicates the type of scheduled query rule. The default is LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
<p>Location: The geo-location where the resource lives</p>
</td>
</tr>
<tr>
<td>
<code>muteActionsDuration</code><br/>
<em>
string
</em>
</td>
<td>
<p>MuteActionsDuration: Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>operatorSpec</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleOperatorSpec">
ScheduledQueryRuleOperatorSpec
</a>
</em>
</td>
<td>
<p>OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
passed directly to Azure</p>
</td>
</tr>
<tr>
<td>
<code>overrideQueryTimeRange</code><br/>
<em>
string
</em>
</td>
<td>
<p>OverrideQueryTimeRange: If specified then overrides the query time range (default is
WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>resolveConfiguration</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.RuleResolveConfiguration">
RuleResolveConfiguration
</a>
</em>
</td>
<td>
<p>ResolveConfiguration: Defines the configuration for resolving fired alerts. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>scopesReferences</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
[]genruntime.ResourceReference
</a>
</em>
</td>
<td>
<p>ScopesReferences: The list of resource id&rsquo;s that this scheduled query rule is scoped to.</p>
</td>
</tr>
<tr>
<td>
<code>severity</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRuleProperties_Severity">
ScheduledQueryRuleProperties_Severity
</a>
</em>
</td>
<td>
<p>Severity: Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only
for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>skipQueryValidation</code><br/>
<em>
bool
</em>
</td>
<td>
<p>SkipQueryValidation: The flag which indicates whether the provided query should be validated or not. The default is
false. Relevant only for rules of the kind LogAlert.</p>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
<p>Tags: Resource tags.</p>
</td>
</tr>
<tr>
<td>
<code>targetResourceTypes</code><br/>
<em>
[]string
</em>
</td>
<td>
<p>TargetResourceTypes: List of resource type of the target resource(s) on which the alert is created/updated. For example
if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert
will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of
the kind LogAlert</p>
</td>
</tr>
<tr>
<td>
<code>windowSize</code><br/>
<em>
string
</em>
</td>
<td>
<p>WindowSize: The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
Relevant and required only for rules of the kind LogAlert.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.SystemData_CreatedByType_STATUS">SystemData_CreatedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.SystemData_STATUS">SystemData_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.SystemData_LastModifiedByType_STATUS">SystemData_LastModifiedByType_STATUS
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.SystemData_STATUS">SystemData_STATUS</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Application&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Key&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ManagedIdentity&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;User&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.SystemData_STATUS">SystemData_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.ScheduledQueryRule_STATUS">ScheduledQueryRule_STATUS</a>)
</p>
<div>
<p>Metadata pertaining to creation and last modification of the resource.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>createdAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedAt: The timestamp of resource creation (UTC).</p>
</td>
</tr>
<tr>
<td>
<code>createdBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>CreatedBy: The identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>createdByType</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.SystemData_CreatedByType_STATUS">
SystemData_CreatedByType_STATUS
</a>
</em>
</td>
<td>
<p>CreatedByType: The type of identity that created the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedAt</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedAt: The timestamp of resource last modification (UTC)</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedBy</code><br/>
<em>
string
</em>
</td>
<td>
<p>LastModifiedBy: The identity that last modified the resource.</p>
</td>
</tr>
<tr>
<td>
<code>lastModifiedByType</code><br/>
<em>
<a href="#insights.azure.com/v1api20240101preview.SystemData_LastModifiedByType_STATUS">
SystemData_LastModifiedByType_STATUS
</a>
</em>
</td>
<td>
<p>LastModifiedByType: The type of identity that last modified the resource.</p>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.UserAssignedIdentityDetails">UserAssignedIdentityDetails
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Identity">Identity</a>)
</p>
<div>
<p>Information about the user assigned identity for the resource</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>reference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1api20240101preview.UserIdentityProperties_STATUS">UserIdentityProperties_STATUS
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1api20240101preview.Identity_STATUS">Identity_STATUS</a>)
</p>
<div>
<p>User assigned identity properties.</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>clientId</code><br/>
<em>
string
</em>
</td>
<td>
<p>ClientId: The client id of user assigned identity.</p>
</td>
</tr>
<tr>
<td>
<code>principalId</code><br/>
<em>
string
</em>
</td>
<td>
<p>PrincipalId: The principal id of user assigned identity.</p>
</td>
</tr>
</tbody>
</table>
<hr/>
