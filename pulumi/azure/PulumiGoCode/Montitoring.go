// Application Insights

package PulumiGoCode

import (
	// "github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/appinsights"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Montitoring(ctx *pulumi.Context, resourceGroup *resources.ResourceGroup) error {

	simpleCheck, err := appinsights.NewInsights(ctx, "simplecheck", &appinsights.InsightsArgs{
		Location:          resourceGroup.Location,
		ResourceGroupName: resourceGroup.Name,
		ApplicationType:   pulumi.String("web"),
		Tags: pulumi.StringMap{
			"deploy": pulumi.String("pulumi"),
		},
	})
	if err != nil {
		return err
	}
	ctx.Export("instrumentationKey", simpleCheck.InstrumentationKey)
	ctx.Export("appId", simpleCheck.AppId)

	_, err2 := appinsights.NewStandardWebTest(ctx, "webtest-01", &appinsights.StandardWebTestArgs{
		ResourceGroupName:     resourceGroup.Name,
		ApplicationInsightsId: simpleCheck.ID(),
		GeoLocations: pulumi.StringArray{
			pulumi.String("emea-nl-ams-azr"),
		},
		Request: &appinsights.StandardWebTestRequestArgs{
			Url: pulumi.String("https://olaf-radicke.de"),
		},
		Frequency: pulumi.Int(900),
		Location:  resourceGroup.Location,
		Name:      pulumi.String("olaf-radicke-de"),
		Tags: pulumi.StringMap{
			"deploy": pulumi.String("pulumi"),
		},
		Timeout: pulumi.Int(120),
	})
	if err2 != nil {
		return err
	}

	// _, err = insights.NewAlertRule(ctx, "alertRule", &insights.AlertRuleArgs{
	// 	// Action: insights.RuleEmailActionArgs{
	// 	// 	CustomEmails: {
	// 	// 		pulumi.String("briefkasten@olaf-radicke.de")
	// 	// 	},
	// 	// },
	// 	Condition: insights.ThresholdRuleCondition{
	// 		DataSource: insights.RuleMetricDataSource{
	// 			MetricName:  &pulumi.String("Requests"),
	// 			OdataType:   pulumi.String("Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"),
	// 			ResourceUri: standardWebTest.Name,
	// 			// ResourceUri: "/subscriptions/b67f7fec-69fc-4974-9099-a26bd6ffeda3/resourceGroups/Rac46PostSwapRG/providers/Microsoft.Web/sites/leoalerttest",
	// 		},
	// 		OdataType:       pulumi.String("Microsoft.Azure.Management.Insights.Models.ThresholdRuleCondition"),
	// 		Operator:        insights.ConditionOperatorGreaterThan,
	// 		Threshold:       3,
	// 		TimeAggregation: insights.TimeAggregationOperatorTotal,
	// 		WindowSize:      pulumi.String("PT5M"),
	// 	},
	// 	Description:       pulumi.String("alert rule for availability test \"olaf-radicke-de\""),
	// 	IsEnabled:         pulumi.Bool(true),
	// 	Location:          pulumi.String("global"),
	// 	Name:              pulumi.String("olaf-radicke-de"),
	// 	ResourceGroupName: resourceGroup.Name,
	// 	RuleName:          pulumi.String("olaf-radicke-de"),
	// 	Tags: pulumi.StringMap{
	// 		"deploy": pulumi.String("pulumi"),
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }

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
