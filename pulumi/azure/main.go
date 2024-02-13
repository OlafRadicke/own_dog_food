package main

import (
	// "azure/PulumiGoCode"
	// "github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	// "github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func AzureCloud(ctx *pulumi.Context) error {
	// resourceGroup, err := resources.NewResourceGroup(ctx, "pulumi", nil)
	// if err != nil {
	// 	return err
	// }

	// err = PulumiGoCode.StaticWebsite(ctx, resourceGroup)
	// if err != nil {
	// 	return err
	// }

	// err = PulumiGoCode.Montitoring(ctx, resourceGroup)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func main() {

	pulumi.Run(AzureCloud)
}
