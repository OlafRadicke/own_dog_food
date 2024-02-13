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
