// Application Insights

package PulumiGoCode

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/appinsights"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Montitoring(ctx *pulumi.Context, resourceGroup *resources.ResourceGroup) error {

	simpleCheck, err := appinsights.NewInsights(ctx, "simplecheck", &appinsights.InsightsArgs{
		Location:          resourceGroup.Location,
		ResourceGroupName: resourceGroup.Name,
		ApplicationType:   pulumi.String("web"),
	})
	if err != nil {
		return err
	}
	ctx.Export("instrumentationKey", simpleCheck.InstrumentationKey)
	ctx.Export("appId", simpleCheck.AppId)

	standardWebTest, err2 := appinsights.NewStandardWebTest(ctx, "olaf-radicke-de", &appinsights.StandardWebTestArgs{
		ResourceGroupName:     resourceGroup.Name,
		Location:              pulumi.String("West Europe"),
		ApplicationInsightsId: simpleCheck.ID(),
		GeoLocations: pulumi.StringArray{
			pulumi.String("example"),
		},
		Request: &appinsights.StandardWebTestRequestArgs{
			Url: pulumi.String("https://olaf-radicke.de"),
		},
	})
	if err2 != nil {
		return err
	}

	_, err = insights.NewAlertRule(ctx, "alertRule", &insights.AlertRuleArgs{
		Actions: pulumi.Array{},
		Condition: insights.ThresholdRuleCondition{
			DataSource: insights.RuleMetricDataSource{
				MetricName:  "Requests",
				OdataType:   "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource",
				ResourceUri: standardWebTest.Name,
				// ResourceUri: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest",
			},
			OdataType:       "Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition",
			Operator:        insights.ConditionOperatorGreaterThan,
			Threshold:       3,
			TimeAggregation: insights.TimeAggregationOperatorTotal,
			WindowSize:      "PT5M",
		},
		Description:       pulumi.String("alert rule for availability test \"website-the-independent-friend-de-website-monitoring\""),
		IsEnabled:         pulumi.Bool(true),
		Location:          pulumi.String("global"),
		Name:              pulumi.String("website-the-independent-friend-de-website-monitoring"),
		ResourceGroupName: resourceGroup.Name,
		RuleName:          pulumi.String("website-the-independent-friend-de-website-monitoring"),
		Tags:              nil,
	})
	if err != nil {
		return err
	}

	return nil
}

// {
//     "id": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/metricalerts/website-the-independent-friend-de-website-monitoring",
//     "name": "website-the-independent-friend-de-website-monitoring",
//     "type": "Microsoft.Insights/metricAlerts",
//     "location": "global",
//     "tags": {
//         "hidden-link:/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring": "Resource",
//         "hidden-link:/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-the-independent-friend-de-website-monitoring": "Resource"
//     },
//     "properties": {
//         "description": "Automatically created alert rule for availability test \"website-the-independent-friend-de-website-monitoring\"",
//         "severity": 1,
//         "enabled": true,
//         "scopes": [
//             "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-the-independent-friend-de-website-monitoring",
//             "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring"
//         ],
//         "evaluationFrequency": "PT1M",
//         "windowSize": "PT5M",
//         "criteria": {
//             "webTestId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-the-independent-friend-de-website-monitoring",
//             "componentId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring",
//             "failedLocationCount": 2,
//             "odata.type": "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria"
//         },
//         "actions": [
//             {
//                 "actionGroupId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/actiongroups/action-group-prod",
//                 "webHookProperties": {}
//             }
//         ]
//     }
// }

// ##########################################################

// {
//     "id": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/metricalerts/website-olaf-radicke-de-website-monitoring",
//     "name": "website-olaf-radicke-de-website-monitoring",
//     "type": "Microsoft.Insights/metricAlerts",
//     "location": "global",
//     "tags": {
//         "hidden-link:/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring": "Resource",
//         "hidden-link:/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-olaf-radicke-de-website-monitoring": "Resource"
//     },
//     "properties": {
//         "description": "Automatically created alert rule for availability test \"website-olaf-radicke-de-website-monitoring\"",
//         "severity": 1,
//         "enabled": true,
//         "scopes": [
//             "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-olaf-radicke-de-website-monitoring",
//             "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring"
//         ],
//         "evaluationFrequency": "PT1M",
//         "windowSize": "PT5M",
//         "criteria": {
//             "webTestId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/webtests/website-olaf-radicke-de-website-monitoring",
//             "componentId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/components/website-monitoring",
//             "failedLocationCount": 2,
//             "odata.type": "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria"
//         },
//         "actions": [
//             {
//                 "actionGroupId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/env-prod-eu-west/providers/microsoft.insights/actiongroups/action-group-prod",
//                 "webHookProperties": {}
//             },
//             {
//                 "actionGroupId": "/subscriptions/532be1e2-6860-471e-841e-2575cd27b43f/resourcegroups/azureapp-auto-alerts-8534e7-briefkasten_olaf_radicke_de/providers/microsoft.insights/actiongroups/azureapp-auto",
//                 "webHookProperties": {}
//             }
//         ]
//     }
// }
